package main

import (
	pg "Go/databases/pokemon_postgres"
	pgconfig "Go/databases/pokemon_postgres/config"
	log "Go/logging"
	"Go/time_completion"
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"os"
	"strconv"
)

//func SetupTaskRoutes(router fiber.Router) {
//	group := router.Group("/tasks")
//
//	group.Get("/", controllers.GetTasks)
//	group.Post("/", controllers.CreateTask)
//	group.Patch("/:id", controllers.UpdateTask)
//	group.Delete("/:id", controllers.DeleteTask)

/*
	todo
	docs: https://docs.gofiber.io/api/app/
	read pongo2, cassandra
	ssl
	raise threads?
	front end
*/

func main() {
	pgconfig.LoadConfig()
	db, err := sql.Open("postgres", viper.GetString(pgconfig.CONTEXT))
	//todo separate config for logger in logger directory
	viper.Set("loglevel", "debug")
	lg := log.NewLogger()

	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		os.Exit(-1)
	}

	//minio := minio.Client{
	//	endpointURL: "localhost:9000",
	//	secure:      false,
	//}
	endpoint := "localhost:9000"
	accessKeyID := "poke"
	secretAccessKey := "pokemon123"
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}

	app := fiber.New()
	app.Use(recover.New()) //Allows recovery from errors

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Established a successful connection!")
	})

	app.Get("/pokedex/:dexnum", func(c *fiber.Ctx) error {

		defer time_completion.LogTimer(&lg, fmt.Sprintf("Get %s", c.Path()))()

		atoi, err := strconv.Atoi(c.Params("dexnum"))
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Dexnum %s is not an Int\n", c.Params("dexnum")))
		}
		if atoi < 1 || atoi > 151 {
			lg.Error().Msgf("pokedex number out of range %s", c.Params("dexnum"))
		}
		if atoi < 1 || atoi > 151 {
			return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("Dexnum %s is out of range\n", c.Params("dexnum")))
		}

		pokemon, err := pg.GetAPokemon(c.Params("dexnum"), db)
		if err != nil {
			println(err)
			return err
		} else {
			return c.SendString(fmt.Sprintf("%+v", pokemon))
		}

	})

	app.Get("/pokedex", func(c *fiber.Ctx) error {

		allPokemon, err := pg.GetAllPokemon(db)
		if err != nil {
			println(err)
			return c.SendString(err.Error())
		} else {
			rstr := ""
			//todo return json
			for k := range allPokemon {
				rstr = rstr + fmt.Sprintf("%+-10v\n\n", allPokemon[k])
			}
			return c.SendString(rstr)
		}
	})

	app.Get("/image/:dexnum", func(c *fiber.Ctx) error {
		atoi, err := strconv.Atoi(c.Params("dexnum"))
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Dexnum %s is not an Int\n", c.Params("dexnum")))
		}
		println(atoi)
		if atoi < 1 || atoi > 151 {
			return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("Dexnum %s is out of range\n", c.Params("dexnum")))
		}
		pokemon, err := pg.GetAPokemon(c.Params("dexnum"), db)
		fmt.Printf("%+v\n", pokemon)
		filePath := fmt.Sprintf("%03d%s.png", pokemon.Dexnum, pokemon.Name)
		bucket := "pokemon"

		println(filePath)
		//todo cache the file
		//todo check if file already exists
		err = minioClient.FGetObject(context.Background(), bucket, "pokemon\\"+filePath, "./"+filePath, minio.GetObjectOptions{})
		if err != nil{
			println(err.Error())
		}
		println("sending file")
		return c.SendFile(filePath, true)
	})

	err = app.Listen(":3000")
	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}

}
