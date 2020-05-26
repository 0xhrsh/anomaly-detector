package main

import (
	"errors"
)

// AnomalyDetector provides operations to detect anomalies.
type AnomalyDetector interface {
	FindAnomaly(string) (int, error)
}

// appInfo is a concrete implementation of AnomalyDetector
type anomalyDetector struct {
	apps map[string]app
}

type app struct {
	Date     string `json:"date"`
	ID       string `json:"app"`
	Dau      int    `json:"dau"`
	Requests int    `json:"requests"`
	Response int    `json:"responses"`
}

// FindAnomaly finds anomaly for a given app
func (svc anomalyDetector) FindAnomaly(ID string) (int, error) {

	if ID == "" {
		return 404, nil
	}

	ret := getAppData(ID)

	return int(ret[0]), nil

}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")
