package domain

import "encoding/json"

type MType = string
type MName = string

var (
	Gauge   MType = "gauge"
	Counter MType = "counter"
)

var (
	Alloc       MName = "Alloc"
	FreeMemory  MName = "FreeMemory"
	PollCount   MName = "PollCount"
	TotalMemory MName = "TotalMemory"
)

type Metric struct {
	Type        MType  `json:"Type"`
	Name        MName  `json:"Name"`
	Description string `json:"Description"`
	Value       int64  `json:"Value"`
}

func (m Metric) Bytes() ([]byte, error) {
	return json.Marshal(m)
}

func NewMetric(data []byte) (Metric, error) {
	var m Metric
	err := json.Unmarshal(data, &m)
	return m, err
}

type List struct {
	Alloc       Metric `json:"Alloc"`
	FreeMemory  Metric `json:"FreeMemory"`
	PollCount   Metric `json:"PollCount"`
	TotalMemory Metric `json:"TotalMemory"`
}
