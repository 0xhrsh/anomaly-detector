package main

import (
	"time"
)



type Config struct {
	UserID    string `required:"true"`
	AuthToken string `required:"true"`
	Endpoint  string `required:"true"`
}

// App contains all fields of app
type App struct {
	Date        time.Time `json:"date"`
	ID          string    `json:"app"`
	Dau         int       `json:"dau"`
	Requests    int       `json:"requests"`
	Responses   int       `json:"responses"`
	Impressions int       `json:"impressions"`
}

type appNumbers struct {
	app             App
	meanDau         float64
	stdDau          float64
	meanRequests    float64
	stdRequests     float64
	meanResponses   float64
	stdResponses    float64
	meanImpressions float64
	stdImpressions  float64
}




