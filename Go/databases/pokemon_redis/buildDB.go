package redis

import (
	"Go/csv"
	"Go/docker"
	m "Go/models"
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"os"
	"path"
	"strconv"
)


func BuildDB(){
	ctx := context.Background()

	build := docker.ContainerBuild{
		ImgName:       "redis/redis-stack",
		ContainerName: "pokemon-redis",
		Port:          "6379",
		EnvVars: []string{},
	}

	docker.BuildContainer(ctx, build)

	wd, _ := os.Getwd()
	fPath := path.Join(wd, "csv", "Pokemon.csv")

	pokedex, err := csv.LoadPokemonCSV(fPath)
	if err != nil {
		println(err.Error())
	}

	dbm, err := NewConnection(ctx,)
	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}

	var b bytes.Buffer
	enc := gob.NewEncoder(&b)

	println(dbm.Len(ctx))
	for _, p := range pokedex{
		err = enc.Encode(p)
		if err != nil {
			println(err.Error())
		}
		dbm.Set(ctx, strconv.Itoa(p.Dexnum), b.Bytes())
		//err = dbm.DB.Set(ctx, , b.Bytes(), 5 * time.Second).Err()
		//if err != nil {
		//	println(err.Error())
		//}
	}
	println(dbm.Len(ctx))
	cmdb, _ := dbm.Get(ctx,"1")

	//cmdb, err := dbm.DB.Get(ctx, "1").Bytes()
	//if err != nil{
	//	println(err.Error())
	//}
	br := bytes.NewReader(cmdb)

	var res m.Pokemon
	dec := gob.NewDecoder(br)
	err = dec.Decode(&res)
	if err != nil{
		println(err.Error())
	}
	fmt.Printf("%v\n", res)

	err = enc.Encode(res)
	if err != nil {
		println(err.Error())
	}

	dbm.Set(ctx,"2", b.Bytes())
	cmdb, _ = dbm.Get(ctx,"2")
	err = dec.Decode(&res)
	if err != nil {
		println(err.Error())
	}
	println(dbm.Len(ctx))
	dbm.Delete(ctx,"150")
	println(dbm.Len(ctx))
	dbm.Delete(ctx,"150")

}




