package main

import (
	"log"

	"github.com/its-kos/web-crawler/api"
)

func main() {
	server := api.NewApiServer("localhost:8080")
	log.Fatal(server.Run())
}
