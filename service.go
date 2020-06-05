package main

import (
	"errors"
	"math"
	"time"
)

// AnomalyDetector provides operations to detect anomalies.
type AnomalyDetector interface {
	FindAnomaly(string, string) (int, error)
}

// appInfo is a concrete implementation of AnomalyDetector
type anomalyDetector struct {
	nostalgia Nostalgia
	num       appNumbers
	config    Config
}

// FindAnomaly finds anomaly for a given app
func (svc anomalyDetector) FindAnomaly(ID string, Date string) (int, error) {

	if ID == "" {
		return 0, ErrEmpty
	}

	if Date == "" {
		Date = time.Now().Format("2006-01-02")
	}

	var err error

	svc.num.app.ID = ID

	svc.num.app.Date, err = time.Parse("2006-01-02", Date)
	if err != nil {
		return 0, err
	}

	err = svc.num.getAppNumbers(svc.nostalgia)
	if err != nil {
		return 0, err
	}

	code := 0

	if compareMetric(float64(svc.num.app.Dau), svc.num.meanDau, svc.num.stdDau) {
		code += 1000
	}
	if compareMetric(float64(svc.num.app.Impressions), svc.num.meanImpressions, svc.num.stdImpressions) {
		code += 100
	}
	if compareMetric(float64(svc.num.app.Requests), svc.num.meanRequests, svc.num.stdRequests) {
		code += 10
	}
	if compareMetric(float64(svc.num.app.Responses), svc.num.meanResponses, svc.num.stdResponses) {
		code++
	}

	return code, nil

}

func compareMetric(num float64, mean float64, stdDev float64) bool {
	if num > mean+math.Min(2*stdDev, 0.2*mean) {
		return true
	}
	if num < mean-math.Min(2*stdDev, 0.15*mean) {
		return true
	}
	return false
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

func NewAnomalyDetector(config Config) AnomalyDetector {
	svc := &anomalyDetector{
		config:    config,
		nostalgia: NewNostalgiaService(config),
	}

	return svc
}
