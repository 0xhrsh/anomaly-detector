package main

import (
	"errors"
	"fmt"
	"math"
	"time"
)

// AnomalyDetector provides operations to detect anomalies.
type AnomalyDetector interface {
	FindAnomaly(string, string, string) ([]AppResponse, error)
}

// appInfo isAnomaly a concrete implementation of AnomalyDetector
type anomalyDetector struct {
	nostalgia Nostalgia
	hermes    Hermes
	num       appNumbers
	config    Config
}

// FindAnomaly finds anomaly for a given app
func (svc anomalyDetector) FindAnomaly(ID string, Start string, End string) ([]AppResponse, error) {
	var (
		resp  []AppResponse
		err   error
		start time.Time
		end   time.Time
	)

	if ID == "" {
		return resp, ErrEmpty
	}

	svc.num.app.ID = ID

	start, err = time.Parse("2006-01-02", Start)
	if err != nil {
		return resp, err
	}

	end, err = time.Parse("2006-01-02", End)
	if err != nil {
		return resp, err
	}

	for d := start; d.Before(end); d = d.AddDate(0, 0, 1) {
		var (
			dateResponse AppResponse
			isAnomaly    IsAnomaly
		)
		dateResponse.AnomalyTime = d.Format("2006-01-02")

		svc.num.app.Date = d
		err = svc.num.getAppNumbers(svc.nostalgia)
		if err != nil {
			dateResponse.Err = fmt.Sprint(err)
			resp = append(resp, dateResponse)
			continue
		}

		dateResponse.AnomalyDau, isAnomaly.Dau = compareMetric(float64(svc.num.app.Dau), svc.num.meanDau, svc.num.stdDau)
		dateResponse.AnomalyImpressions, isAnomaly.Impressions = compareMetric(float64(svc.num.app.Impressions), svc.num.meanImpressions, svc.num.stdImpressions)
		dateResponse.AnomalyRequests, isAnomaly.Requests = compareMetric(float64(svc.num.app.Requests), svc.num.meanRequests, svc.num.stdRequests)
		dateResponse.AnomalyResponses, isAnomaly.Responses = compareMetric(float64(svc.num.app.Responses), svc.num.meanResponses, svc.num.stdResponses)

		if isAnomaly.Dau || isAnomaly.Impressions || isAnomaly.Requests || isAnomaly.Responses {
			codeChanges, err := svc.hermes.CodeChanges(d, isAnomaly)
			if err != nil {
				dateResponse.Err = fmt.Sprint(err)
				resp = append(resp, dateResponse)
				continue
			}

			activityLog, err := svc.hermes.SystemChanges(d, isAnomaly)
			if err != nil {
				dateResponse.Err = fmt.Sprint(err)
				resp = append(resp, dateResponse)
				continue
			}
			dateResponse.CodeChanges = codeChanges
			dateResponse.ActivityLog = activityLog
		}
		resp = append(resp, dateResponse)
	}

	return resp, nil

}

func compareMetric(num float64, mean float64, stdDev float64) (int, bool) {
	if num > mean+math.Min(2*stdDev, 0.2*mean) {
		return 1, true
	}
	if num < mean-math.Min(2*stdDev, 0.15*mean) {
		return -1, true
	}
	return 0, false
}

// ErrEmpty isAnomaly returned when an input string isAnomaly empty.
var ErrEmpty = errors.New("empty string")

func newAnomalyDetector(config Config) AnomalyDetector {
	svc := &anomalyDetector{
		config:    config,
		nostalgia: newNostalgiaService(config),
		hermes:    newHermesService(config),
	}

	return svc
}
