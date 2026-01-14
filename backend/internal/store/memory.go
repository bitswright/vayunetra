package store

import (
	"sync"

	"github.com/bitswright/vayunetra/backend/internal/model"
)

type MemoryStore struct {
	mutex    sync.Mutex
	readings []model.SensorReading
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		readings: make([]model.SensorReading, 0),
	}
}

func (ms *MemoryStore) Save(sr model.SensorReading) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	ms.readings = append(ms.readings, sr)
}

func (ms *MemoryStore) FetchLatest() (model.SensorReading, error) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	if len(ms.readings) == 0 {
		return model.SensorReading{}, NoDataAvailableErr
	}
	return ms.readings[len(ms.readings)-1], nil
}
