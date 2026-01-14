package model

type SensorReading struct {
	SensorId  string `json:"sensor_id"`
	Timestamp int64  `json:"timestamp"`

	TempC       float64 `json:"temperature_c"`
	HumidityPct float64 `json:"humidity_pct"`
	PressureHpa float64 `json:"pressure_hpa"`

	PM25 float64 `json:"pm25_ugm3,omitempty"`
	PM10 float64 `json:"pm10_ugm3,omitempty"`
}
