package models

import "time"

type DatadogQueryResponse struct {
	Logs []struct {
		ID      string `json:"id"`
		Content struct {
			Timestamp  time.Time `json:"timestamp"`
			Tags       []string  `json:"tags"`
			Attributes struct {
				CustomAttribute int `json:"customAttribute"`
				Duration        int `json:"duration"`
			} `json:"attributes"`
			Host    string `json:"host"`
			Service string `json:"service"`
			Message string `json:"message"`
		} `json:"content"`
	} `json:"logs"`
	NextLogID string `json:"nextLogId"`
	Status    string `json:"status"`
}

