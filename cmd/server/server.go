package main

import (
	"log"

	"github.com/ken343/maritime-royale/pkg/server"
	_ "github.com/ken343/maritime-royale/pkg/server"
)

func main() {
	log.Println("Maritime Royale Server Listening on Port :8080...")
	server.Server(":8080")
}
