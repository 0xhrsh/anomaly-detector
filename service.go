package main

import (
	"database/sql"
	"errors"
)

// AnomalyDetector provides operations to detect anomalies.
type AnomalyDetector interface {
	FindAnomaly(string, string) (int, error)
}

// appInfo is a concrete implementation of AnomalyDetector
type anomalyDetector struct {
	db   *sql.DB
	apps map[string]App
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
		return 404, nil
	}

	ret, app := getAppData(ID, Date, svc.db)

	return app.Dau - int(ret.meanDau), nil

}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")
