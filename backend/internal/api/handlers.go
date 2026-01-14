package api

import (
	"encoding/json"
	"net/http"

	"github.com/bitswright/vayunetra/backend/internal/model"
	"github.com/bitswright/vayunetra/backend/internal/store"
)

type Handler struct {
	Store store.Store
}

func NewHandler(store store.Store) *Handler {
	return &Handler{
		Store: store,
	}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func (h *Handler) PostReading(w http.ResponseWriter, r *http.Request) {
	var sr model.SensorReading

	if err := json.NewDecoder(r.Body).Decode(&sr); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
	}

	h.Store.Save(sr)
	w.WriteHeader(http.StatusAccepted)
}

func (h *Handler) GetLatestReading(w http.ResponseWriter, r *http.Request) {
	sensorReading, err := h.Store.FetchLatest()
	if err != nil {
		// TODO: is StatusNoContent correct status?
		http.Error(w, err.Error(), http.StatusNoContent)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&sensorReading)
}
