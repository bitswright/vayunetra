package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Start(handler http.Handler, port int) {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
