package models


type TimeSeriesResponse []struct {
	Target     string       `json:"target"`
	Datapoints [][]float64  `json:"datapoints"`
}
