package main

import (
	"errors"
	"time"
)

// AnomalyDetector provides operations to detect anomalies.
type AnomalyDetector interface {
	FindAnomaly(string, string) (int, error)
}

// appInfo is a concrete implementation of AnomalyDetector
type anomalyDetector struct {
	app App
}

// App contains all fields of app
type App struct {
	Date     string `json:"date"`
	ID       string `json:"app"`
	Dau      int    `json:"dau"`
	Requests int    `json:"requests"`
	Response int    `json:"responses"`
}

// FindAnomaly finds anomaly for a given app
func (svc anomalyDetector) FindAnomaly(ID string, Date string) (int, error) {

	if ID == "" {
		return 0, ErrEmpty
	}

	if Date == "" {
		Date = time.Now().Format("2006-01-02")
	}
	var numbers appNumbers

	numbers.app.ID = ID
	numbers.app.Date = Date

	err := numbers.getAppNumbers()
	if err != nil {
		return 0, err
	}
	err = numbers.app.getAppData(3)
	if err != nil {
		return 0, err
	}
	code := 0
	if float64(numbers.app.Dau) > numbers.meanDau+2*numbers.stdDau || float64(numbers.app.Dau) < numbers.meanDau-2*numbers.stdDau {
		code++
	}

	if float64(numbers.app.Requests) > numbers.meanRequests+2*numbers.stdRequests || float64(numbers.app.Requests) < numbers.meanRequests-2*numbers.stdRequests {
		code += 10
	}

	if float64(numbers.app.Response) > numbers.meanResponse+2*numbers.stdResponse || float64(numbers.app.Response) < numbers.meanResponse-2*numbers.stdResponse {
		code += 100
	}

	return code, nil

}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")
