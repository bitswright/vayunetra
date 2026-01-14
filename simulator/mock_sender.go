package simulator

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/bitswright/vayunetra/backend/internal/model"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	sensorId := "vayunetra-sensor-001"
	backendURL := "http://localhost:8080/readings"
	for {
		sr := mockSensorReading(sensorId)
		jsonData, err := json.Marshal(sr)
		if err != nil {
			log.Println("Error marshalling reading:", err)
			continue
		}
		resp, err := http.Post(backendURL, "application/json", bytes.NewBuffer(jsonData))
		defer resp.Body.Close()
		if err != nil {
			log.Println("Error sending reading", err)
		} else {
			log.Printf(
				"Sent reading: Temp=%.2f°C, Humidity=%.2f%%, Pressure=%.2f hPa\n",
				sr.TemperatureC,
				sr.HumidityPct,
				sr.PressureHpa,
			)
		}

		if resp.StatusCode != http.StatusBadRequest {
			log.Printf(
				"Error while sending request: Sent reading: Temp=%.2f°C, Humidity=%.2f%%, Pressure=%.2f hPa\n",
				sr.TemperatureC,
				sr.HumidityPct,
				sr.PressureHpa,
			)
		}

		time.Sleep(2 * time.Second)
	}
}

func mockSensorReading(sensorId string) model.SensorReading {
	return model.SensorReading{
		SensorId:     sensorId,
		Timestamp:    time.Now().Unix(),
		TemperatureC: 20 + rand.Float64()*10,
		HumidityPct:  40 + rand.Float64()*20,
		PressureHpa:  1000 + rand.Float64()*20,
	}
}
