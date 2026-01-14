package server

import (
	"net/http"

	"github.com/bitswright/vayunetra/backend/internal/api"
	"github.com/bitswright/vayunetra/backend/internal/store"
)

func NewRouter(s store.Store) http.Handler {
	h := api.NewHandler(s)
	mux := http.NewServeMux()

	mux.HandleFunc("/health", h.Health)
	mux.HandleFunc("/readings", h.PostReading)
	mux.HandleFunc("/latest", h.GetLatestReading)

	return mux
}
