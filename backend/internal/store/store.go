package store

import (
	"errors"

	"github.com/bitswright/vayunetra/backend/internal/model"
)

var NoDataAvailableErr = errors.New("No Readings Available")

type Store interface {
	Save(model.SensorReading)
	FetchLatest() (model.SensorReading, error)
}
