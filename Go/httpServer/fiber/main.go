package main

import (
	cfg "Go/configs"
	pg "Go/databases/pokemon_postgres"
	h "Go/httpServer/fiber/helpers"
	log "Go/logging"
	m "Go/models"
	"Go/time_completion"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
	"net/http"
	"os"
	"path"
	"sort"
	"strconv"
	"time"
)

//todo json DB
//todo metrics, ssl certs, benchmarks
//todo Learn db unmarshal

type Fiber struct {
	lg         zerolog.Logger
	bunDB      *bun.DB
	ctx        context.Context
	listenAddr string
}

//func SetupTaskRoutes(router fiber.Router) {
//	group := router.Group("/tasks")
//
//	group.Select("/", controllers.GetTasks)
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

func (f *Fiber) appRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		//TODO return list of url links if html req
		appRoutes := app.GetRoutes()

		sort.Slice(appRoutes, func(i, j int) bool {
			return appRoutes[i].Path < appRoutes[j].Path
		})

		routeList := ""
		for _, v := range appRoutes {
			if v.Path == "/" || v.Method == "HEAD" {
				continue
			} else {
				routeList += fmt.Sprintf("%-25s %6s %s \n", v.Path, v.Method, v.Params)
			}
		}

		return c.SendString(routeList)
	})

	pokedex := app.Group("/pokedex")
	pokedex.Get(":dexnum?/:count?", f.GetPokedex)

	images := app.Group("/image")
	images.Get(":dexnum", f.GetImage)

	icon := app.Group("/icon")
	icon.Get(":typeIcon", f.GetTypeIcon)

	app.Get("/fmetrics", monitor.New(monitor.Config{
		Title:   "Fiber Metrics",
		Refresh: time.Second * 10,
	}))

	//app.Get("/metrics", promhttp.Handler())

}

func (f *Fiber) GetPokedex(c *fiber.Ctx) (err error) {
	//c.Accepts("application/json")
	defer time_completion.LogTimer(&f.lg, fmt.Sprintf("GetPokedex %s", c.Path()))()
	f.lg.Printf("%s %s:%s-->%s ", c.Context().Method(), c.IP(), c.Port(), c.Context().URI())

	if c.Method() != http.MethodGet {
		return fiber.NewError(fiber.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
	}

	var dexnum, count int
	var pokemons []m.Pokemon

	if c.Params("dexnum") != "" {
		dexnum, err = strconv.Atoi(c.Params("dexnum"))
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("dexnum %s is not an int\n", c.Params("dexnum")))
		}
		if dexnum < 1 || dexnum > 151 {
			f.lg.Error().Msgf("pokedex number out of range %s", c.Params("dexnum"))
			return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("dexnum %s is out of range\n", c.Params("dexnum")))
		}
	}

	if c.Params("count") != "" {
		count, err = strconv.Atoi(c.Params("count"))
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("count %s is not an int\n", c.Params("count")))
		}
		if count < dexnum || count > 151 {
			f.lg.Error().Msgf("pokedex number out of range %s", c.Params("count"))
			return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("count %s is out of range\n", c.Params("count")))
		}
	}

	switch count {
	case 0:
		err = f.bunDB.NewSelect().
			Model(&pokemons).
			Where("dexnum = ?", dexnum).
			Scan(c.Context())
	default:
		err = f.bunDB.NewSelect().
			Model(&pokemons).
			Where("dexnum BETWEEN ? and ?", dexnum, count).
			Scan(c.Context())
	}

	if err != nil {
		f.lg.Error().Msgf("query failed %w", err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	//lg.Debug().Msgf("Pokemon %+v\n", pokemons)

	if err := h.WriteJson(c, http.StatusOK, pokemons, f.lg); err != nil {
		return fiber.NewError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	return nil
}

func (f *Fiber) GetImage(c *fiber.Ctx) error {
	defer time_completion.LogTimer(&f.lg, fmt.Sprintf("GetImage %s", c.Path()))()
	f.lg.Printf("%s %s:%s-->%s ", c.Context().Method(), c.IP(), c.Port(), c.Context().URI())

	if c.Method() != http.MethodGet {
		return fiber.NewError(fiber.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
	}

	atoi, err := strconv.Atoi(c.Params("dexnum"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Id %s is not an Int\n", c.Params("dexnum")))
	}

	if atoi < 1 || atoi > 251 {
		return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("Id %s is out of range\n", c.Params("dexnum")))
	}

	filePath := path.Join("images", "pokemon", fmt.Sprintf("%d.png", atoi))
	return c.SendFile(filePath, true)
}

func (f *Fiber) GetTypeIcon(c *fiber.Ctx) error {
	defer time_completion.LogTimer(&f.lg, fmt.Sprintf("GetTypeIcon %s", c.Path()))()
	f.lg.Printf("%s %s:%s-->%s ", c.Context().Method(), c.IP(), c.Port(), c.Context().URI())

	if c.Method() != http.MethodGet {
		return fiber.NewError(fiber.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
	}
	typeIcon := c.Params("typeIcon")

	filePath := path.Join("images", "typeIcons", fmt.Sprintf("%s.png", typeIcon))
	return c.SendFile(filePath, true)
}

//func (f *Fiber) GetMetrics(c *fiber.Ctx) error{
//	//h.WriteJson(c, http.StatusOK, )
//	http.Handle("/metrics", promhttp.Handler())
//	return c.SendString("")
//}

func main() {
	c := cfg.NewPgConfig("pokemon_postgres")
	c.LoadDBConfig()

	f := Fiber{
		ctx:        context.Background(),
		lg:         log.NewZeroLogger(),
		listenAddr: ":3001",
	}

	//graf := cfg.NewGrafanaConfig("grafana")
	//graf.LoadGrafanaConfig()

	f.bunDB = pg.BuildDB(f.ctx)

	//metrics.BuildGrafana(f.ctx)
	//prometheus.BuildPrometheus(f.ctx)
	//
	//metrics.Metrics()

	f.startFiber()

}

func (f *Fiber) startFiber() {
	app := fiber.New()
	app.Use(recover.New()) //Allows recovery from errors

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	})) //https://docs.gofiber.io/api/middleware/cors/
	f.appRoutes(app)

	err := app.Listen(f.listenAddr)
	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}

	//app.Use("/metrics", promhttp.Handler())
}
