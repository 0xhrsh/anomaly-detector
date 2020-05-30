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
	num appNumbers
}

// FindAnomaly finds anomaly for a given app
func (svc anomalyDetector) FindAnomaly(ID string, Date string) (int, error) {

	if ID == "" {
		return 0, ErrEmpty
	}

	if Date == "" {
		Date = time.Now().Format("2006-01-02T15:04:05.000Z")
	}

	var err error

	svc.num.app.ID = ID

	svc.num.app.Date, err = time.Parse("2006-01-02T15:04:05.000Z", Date)
	if err != nil {
		return 0, err
	}

	err = svc.num.getAppNumbers()
	if err != nil {
		return 0, err
	}

	code := 0
	if float64(svc.num.app.Dau) > svc.num.meanDau+2*svc.num.stdDau || float64(svc.num.app.Dau) < svc.num.meanDau-2*svc.num.stdDau {
		code++
	}

	if float64(svc.num.app.Requests) > svc.num.meanRequests+2*svc.num.stdRequests || float64(svc.num.app.Requests) < svc.num.meanRequests-2*svc.num.stdRequests {
		code += 10
	}

	if float64(svc.num.app.Response) > svc.num.meanResponse+2*svc.num.stdResponse || float64(svc.num.app.Response) < svc.num.meanResponse-2*svc.num.stdResponse {
		code += 100
	}

	return code, nil

}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")
