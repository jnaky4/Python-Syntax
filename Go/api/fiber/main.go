package main

import (
	pg "Go/databases/pokemon_postgres"
	pgconfig "Go/databases/pokemon_postgres/config"
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
	//"gorm.io/gorm/schema"
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

func main(){
	pgconfig.LoadConfig()
	db, err := sql.Open("postgres", viper.GetString(pgconfig.CONTEXT))
	if err != nil {
		panic(err)
	}
	//dbg, err := gorm.Open(postgres.Open(viper.GetString(pgconfig.CONTEXT)), &gorm.Config{
	//	NamingStrategy: schema.NamingStrategy{
	//		SingularTable: true,
	//	},
	//})
	//err = dbg.AutoMigrate(&pg.Pokemon{})
	//if err != nil {
	//	panic(err)
	//}


	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	app.Use(recover.New()) //Allows recovery from errors

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Established a successful connection!")
	})

	app.Get("/pokemon/:dexnum", func(c *fiber.Ctx) error {

		//var tables []string
		//if err := dbg.Table("information_schema.tables").Where("table_schema = ?", "public").Pluck("table_name", &tables).Error; err != nil {
		//	panic(err)
		//}
		//fmt.Printf("%+v\n", tables)


		atoi, err := strconv.Atoi(c.Params("dexnum"))
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Dexnum %s is not an Int\n", c.Params("dexnum")))
		}
		if atoi < 1 || atoi > 151{
			return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("Dexnum %s is out of range\n", c.Params("dexnum")))
		}

		//pokemon, err := pg.GormAPokemon(atoi, dbg)
		pokemon, err := pg.GetAPokemon(c.Params("dexnum"), db )
		if err != nil {
			println(err)
			return err
		} else {
			return c.SendString(fmt.Sprintf("%+v", pokemon))
		}

	})

	app.Get("/pokemon", func(c *fiber.Ctx) error {
		//pokemon, err := pg.GormAllPokemon(dbg)
		//if err != nil {
		//	println(err)
		//	return c.SendString(err.Error())
		//}
		//
		//return c.SendString(fmt.Sprintf("%+-10v\n\n", pokemon))

		allPokemon, err := pg.GetAllPokemon(db)
		if err != nil {
			println(err)
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
