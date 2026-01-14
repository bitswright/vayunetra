package main

import (
	"log"

	"github.com/bitswright/vayunetra/backend/internal/server"
	"github.com/bitswright/vayunetra/backend/internal/store"
)

var port = 8080

func main() {
	memoryStore := store.NewMemoryStore()
	router := server.NewRouter(memoryStore)

	log.Printf("VayuNetra Backend running on port %d", port)
	server.Start(router, port)
}
