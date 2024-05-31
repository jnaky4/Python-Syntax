package main

import (
	cfg "Go/configs"
	"Go/docker"
	m "Go/models"
	"Go/time_completion"
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"reflect"
	"time"
)

func main(){
	//redis.BuildDB()
	d := docker.Manager{}
	err := d.Connect(context.Background())
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	defer d.Client.Close()
	version, err := d.GetLatestCachedImgVersion(context.Background(), "postgres")
	if err != nil {
		return
	}
	fmt.Printf("Latest Version: %+v\n", version)
}


type pgConfig struct {
	q pgQuery
	DB *sql.DB
	cfg m.ConfigDBDriver
}

type pgQuery struct {
	op command
	data interface{}
	qstr string
}

type command string

const (
	ALTER command = "ALTER"
	CREATE command = "CREATE"
	DELETE command = "DELETE"
	INSERT command = "INSERT"
	SELECT command = "SELECT"
	UPDATE command = "UPDATE"
)



//type pg interface{
//	Select(string)
//	Insert(string)
//
//}

//todo Select("*").From("Pokedex").Where("x = x")

//func main(){
//	pg := pgConfig{
//		cfg: m.ConfigDBDriver{
//			Port: 5432,
//			Host: "localhost",
//			Username: "postgres",
//			Password: "pokemon",
//		},
//
//		q: pgQuery{
//			op: SELECT,
//			data: m.Pokemon{
//				Id: 1,
//				Name:   "Bulbasaur",
//				Type1: "Grass",
//				Type2: "Poison",
//				Stage: "Basic",
//				EvolveLevel: 16,
//				GenderRatio: "1:07",
//				Height: 0.7,
//				Weight: 6.9,
//				Description: "A strange seed was planted on its back at birth. The plant sprouts and grows with this Pokemon.",
//				Category: "Seed",
//				LvlSpeed: 1,
//				BaseExp: 64,
//				CatchRate: 45,
//			},
//		},
//	}
//
//	err := pg.NewPostgres()
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	//pg.args = []string{"test", "1"}
//	//query := fmt.Sprintf("%s INTO %s(%+v)\n\tVALUES(%s)", pg.op, pg.table, strings.Join(pg.args, ", "), strings.Join(pg.args, ", "))
//	pg.createQuery()
//	err = pg.Execute()
//	if err != nil {
//		println(err.Error())
//	} else {
//		fmt.Printf("%v\n", pg.q.data)
//	}
//
//	//print(pg.q.qstr)
//}

func (p *pgConfig) createQuery() {
	defer time_completion.FunctionTimer(p.createQuery)()
	switch p.q.op {
	case INSERT:
		p.Insert()
	case DELETE:
		p.Delete()
	case SELECT:
		p.Select()
	}

}

func (p *pgConfig) Insert(){
	defer time_completion.FunctionTimer(p.Insert)()
	if reflect.ValueOf(p.q.data).Kind() == reflect.Struct {
		tn := reflect.TypeOf(p.q.data).Name()
		query := fmt.Sprintf("%s INTO %s VALUES(", p.q.op, tn)
		v := reflect.ValueOf(p.q.data)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			case reflect.Float32:
				if i == 0 {
					query = fmt.Sprintf("%s %f", query, v.Field(i).Float())
				} else {
					query = fmt.Sprintf("%s, %f", query, v.Field(i).Float())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		p.q.qstr = query
	}
	fmt.Println("unsupported type")
}

func (p *pgConfig) Delete(){
	defer time_completion.FunctionTimer(p.Delete)()
	tn := reflect.TypeOf(p.q.data).Name()
	da := reflect.ValueOf(p.q.data)
	//id := da.Field(0).Int()
	//name := da.Type().Field(0).Name

	p.q.qstr = fmt.Sprintf("%s FROM %s WHERE %v=%d", p.q.op, tn, da.Type().Field(0).Name, da.Field(0).Int())
	//println(p.q.qstr)
}

func (p *pgConfig) Select(){
	defer time_completion.FunctionTimer(p.Select)()
	tn := reflect.TypeOf(p.q.data).Name()
	da := reflect.ValueOf(p.q.data)
	p.q.qstr = fmt.Sprintf("%s * FROM %s WHERE %v=%d", p.q.op, tn, da.Type().Field(0).Name, da.Field(0).Int())
}

func (p *pgConfig) Execute() error{
	defer time_completion.FunctionTimer(p.Execute)()
	poke := m.Pokemon{}
	//todo fix to take data struct and 1 argument
	err := p.DB.QueryRow(p.q.qstr).Scan(&poke.Dexnum, &poke.Name, &poke.Type1, &poke.Type2, &poke.Stage, &poke.EvolveLevel,
		&poke.GenderRatio, &poke.Height, &poke.Weight, &poke.Description, &poke.Category, &poke.LvlSpeed,
		&poke.BaseExp, &poke.CatchRate)
	p.q.data = poke
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return errors.New("record not found")
	}

	return err
}

func (p *pgConfig) NewPostgres() error {
	cfg.LoadConfig("pokemon_postgres")

	var err error
	db, err := sql.Open("postgres", viper.GetString(cfg.DB_SOURCE))

	//db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		if db != nil {
			db.Close()
		}

		//logger.Error("ERROR from opening postgres db " + err.Error())
		return err
	}

	for i := 0; i < 3; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		time.Sleep(time.Duration(i) * time.Second)
	}

	if err != nil {
		//logger.Error("ERROR from Ping: " + err.Error())
		db.Close()
		return err
	}

	p.DB = db
	//logger.Debug("Postgres db opened successfully")

	//if p.InitialSchemaFile != "" {
	//	schemaFile, err := os.Open(p.InitialSchemaFile)
	//	if err != nil {
	//		logger.Error("error opening initial schema file: " + err.Error())
	//		db.Close()
	//		return nil, err
	//	}
	//	schemaBytes, err := io.ReadAll(schemaFile)
	//	if err != nil {
	//		logger.Error("error reading initial schema file: " + err.Error())
	//		db.Close()
	//		return nil, err
	//	}
	//	_, err = db.Exec(string(schemaBytes))
	//	if err != nil {
	//		logger.Error("error executing initial schema file: " + err.Error())
	//		db.Close()
	//		return nil, err
	//	}
	//	logger.Debug("Postgres schema initialized successfully")
	//}

	return nil
}