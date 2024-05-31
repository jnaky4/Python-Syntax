package main

import (
	"fmt"
	"github.com/cockroachdb/pebble"
)

func main(){

	db, err := pebble.Open("pebble", &pebble.Options{})
	if err != nil {
		println(err.Error())
	}

	key := []byte("hello")
	if err := db.Set(key, []byte("world"), pebble.Sync); err != nil {
		println(err.Error())
	}
	value, closer, err := db.Get(key)
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("%s %s\n", key, value)
	if err := closer.Close(); err != nil {
		println(err.Error())
	}
	if err := db.Close(); err != nil {
		println(err.Error())
	}
}