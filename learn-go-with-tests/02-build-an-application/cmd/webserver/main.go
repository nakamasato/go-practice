package main

import (
	"log"
	"net/http"
	"tmp/learn-go-with-tests/02-build-an-application"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server, err := poker.NewPlayerServer(store)
	if err != nil {
		log.Fatal(err)
	}
	// server := NewPlayerServer(NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5000", server))
}
