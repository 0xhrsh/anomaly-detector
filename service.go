package main

import (
	"errors"
	"math"
	"time"
)

// AnomalyDetector provides operations to detect anomalies.
type AnomalyDetector interface {
	FindAnomaly(string, string) (findAnomalyResponse, error)
}

// appInfo is a concrete implementation of AnomalyDetector
type anomalyDetector struct {
	nostalgia Nostalgia
	num       appNumbers
	config    Config
}

// FindAnomaly finds anomaly for a given app
func (svc anomalyDetector) FindAnomaly(ID string, Date string) (findAnomalyResponse, error) {

	var resp findAnomalyResponse
	resp.AnomalyTime = time.Now()

	if ID == "" {
		return resp, ErrEmpty
	}

	if Date == "" {
		Date = time.Now().Format("2006-01-02")
	}

	var err error

	svc.num.app.ID = ID

	svc.num.app.Date, err = time.Parse("2006-01-02", Date)
	if err != nil {
		return resp, err
	}

	err = svc.num.getAppNumbers(svc.nostalgia)
	if err != nil {
		return resp, err
	}

	resp.AnomalyDau = compareMetric(float64(svc.num.app.Dau), svc.num.meanDau, svc.num.stdDau)

	resp.AnomalyImpressions = compareMetric(float64(svc.num.app.Impressions), svc.num.meanImpressions, svc.num.stdImpressions)

	resp.AnomalyRequests = compareMetric(float64(svc.num.app.Requests), svc.num.meanRequests, svc.num.stdRequests)

	resp.AnomalyResponses = compareMetric(float64(svc.num.app.Responses), svc.num.meanResponses, svc.num.stdResponses)

	return resp, nil

}

func compareMetric(num float64, mean float64, stdDev float64) int {
	if num > mean+math.Min(2*stdDev, 0.2*mean) {
		return 1
	}
	if num < mean-math.Min(2*stdDev, 0.15*mean) {
		return -1
	}
	return 0
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

func newAnomalyDetector(config Config) AnomalyDetector {
	svc := &anomalyDetector{
		config:    config,
		nostalgia: newNostalgiaService(config),
	}

	return svc
}
