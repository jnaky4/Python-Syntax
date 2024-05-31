package main

import (
	"log"
	"net/http"
)

const notesFilename = "notes.json"

type Notes map[string][]string

func main() {
	// Serve static files from the current directory
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	log.Println("Listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}


