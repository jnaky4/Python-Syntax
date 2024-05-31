package pokemon_postgres

import (
	cfg "Go/configs"
	"Go/csv"
	"Go/docker"
	"Go/models"
	"context"
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func ConnectDB() (Models, error) {
	var dbModel Models

	db, err := sql.Open("postgres", viper.GetString(cfg.DB_SOURCE))
	if err != nil {
		return dbModel, err
	}

	for i := 0; i < 3; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		time.Sleep(time.Duration(i) * time.Second)
	}

	if err != nil && err.Error() != "EOF" {
		return dbModel, fmt.Errorf("failed to connect to db error -> %s", err.Error())
	}

	dbModel = NewModels(db)
	return dbModel, nil

}

func BuildDB(ctx context.Context) *bun.DB {
	build := PostgresContainerBuild()
	err := docker.BuildContainer(ctx, build)
	if err != nil {
		log.Fatalf("build failed -> %s ", err.Error())
	}

	pokedex := LoadPokemonFromCSV()

	dbm, err := ConnectDB()
	if err != nil {
		println(err.Error())
		os.Exit(-2)
	}

	db := bun.NewDB(dbm.Pokemon.DB, pgdialect.New())

	_, err = db.NewCreateTable().
		Model((*models.Pokemon)(nil)).
		Exec(ctx)
	if err != nil {
		println(err.Error())
	}

	for _, pokemon := range pokedex {
		_, err = db.NewInsert().
			Model(&pokemon).
			Exec(ctx)
		if err != nil {
			println(err.Error())
		}
	}
	var poke []models.Pokemon

	err = db.NewSelect().
		Model(&poke).
		Where("dexnum = ?", 1).
		Scan(ctx)
	if err != nil {
		println(err.Error())
	}

	if len(poke) < 1 {
		println("Error failed to get pokemon from DB")
	}

	return db
}

func LoadPokemonFromCSV() []models.Pokemon {
	wd, _ := os.Getwd()

	fPath := path.Join(wd, "csv", "Pokemon.csv")
	pokedex, err := csv.LoadPokemonCSV(fPath)
	if err != nil {
		println(err.Error())
	}

	fPath = path.Join(wd, "csv", "Base_Stats.csv")
	stats, err := csv.LoadBaseStatsCSV(fPath)
	if err != nil {
		println(err.Error())
	}

	pokedex = models.MergePokemonStructs(pokedex, stats)

	for i := range pokedex {
		pokedex[i].CalculateTypeEffectiveness()
	}

	return pokedex
}

func PostgresContainerBuild() docker.ContainerBuild {

	return docker.ContainerBuild{
		ImgName:       viper.GetString(cfg.IMAGE_NAME),
		Version:       viper.GetString(cfg.IMAGE_VERSION),
		ContainerName: viper.GetString(cfg.CONTAINER_NAME),
		Port:          viper.GetString(cfg.CONTAINER_PORT),
		EnvVars: []string{
			fmt.Sprintf("POSTGRES_PASSWORD=%s", viper.GetString(cfg.DB_PASSWORD)),
			fmt.Sprintf("POSTGRES_USER=%s", viper.GetString(cfg.DB_USER)),
			fmt.Sprintf("POSTGRES_DB=%s", viper.GetString(cfg.DB_NAME)),
		},
	}
}

/* SQL commands
   psql -d Pokemon -U postgres
   \d pokemon
   CREATE DATABASE Pokemon
   CREATE ROLE dbManager WITH LOGIN DBPASSWORD 'pokemon';
   GRANT SELECT, INSERT, UPDATE, DELETE ON warmState TO furnaceManager
   GRANT USAGE, SELECT ON SEQUENCE warmState_id_seq TO furnaceManager
*/
