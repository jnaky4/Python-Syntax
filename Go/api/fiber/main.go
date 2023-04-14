package main

import (
	pg "Go/databases/pokemon_postgres"
	pgconfig "Go/databases/pokemon_postgres/config"
	log "Go/logging"
	"Go/time_completion"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
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
	logging, metrics
	read pongo2, cassandra
	ssl
	raise threads?
	front end
	return json

 */
func main(){
	//todo separate config for logger in logger directory
	viper.Set("loglevel", "debug")
	db, err := pgconfig.LoadConfig()
	lg := log.NewLogger()
	if err != nil {
		lg.Fatal().
			Err(err).
			Msg("Failed to connect to Postgres db")
	}
	defer db.Close()


	err = db.Ping()
	if err != nil {
		lg.Fatal().
			Err(err).
			Msg("Failed to ping DB")
	}

	app := fiber.New()
	app.Use(recover.New()) //Allows recovery from errors

	app.Get("/", func(c *fiber.Ctx) error {
		defer time_completion.LogTimer(&lg, fmt.Sprintf("Get %s", c.Path()))()
		return c.SendString("Established a successful connection!")
	})

	app.Get("/pokedex/:dexnum", func(c *fiber.Ctx) error {
		defer time_completion.LogTimer(&lg, fmt.Sprintf("Get %s", c.Path()))()

		atoi, err := strconv.Atoi(c.Params("dexnum"))
		if err != nil {
			lg.Error().Msgf("pokedex number invalid %s", c.Params("dexnum"))
			return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Dexnum %s is not an Int\n", c.Params("dexnum")))
		}
		if atoi < 1 || atoi > 151{
			lg.Error().Msgf("pokedex number out of range %s", c.Params("dexnum"))
			return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("Dexnum %s is out of range\n", c.Params("dexnum")))
		}

		pokemon, err := pg.GetAPokemon(c.Params("dexnum"), db )
		if err != nil {
			lg.Error().Err(err).Msg("Query Failed")
			return err
		} else {
			return c.SendString(fmt.Sprintf("%+v", pokemon))
		}

	})

	app.Get("/pokedex", func(c *fiber.Ctx) error {
		defer time_completion.LogTimer(&lg, fmt.Sprintf("Get %s", c.Path()))()

		allPokemon, err := pg.GetAllPokemon(db)
		if err != nil {
			lg.Error().Msgf("Query Failed: %s", err)
			return c.SendString(err.Error())
		} else{
			rstr := ""
			//todo return json
			for k := range allPokemon {
				rstr = rstr + fmt.Sprintf("%+-10v\n\n", allPokemon[k])
			}
			return c.SendString(rstr)
		}
	})

	err = app.Listen(":3000")
	if err != nil {
		return
	}
}
