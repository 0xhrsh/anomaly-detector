package main

import (
	"errors"
	"math"
	"sort"
)

func findStdDev(arr []float64) (float64, float64) {
	sum := 0.0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	mean := float64(sum) / float64(len(arr))

	sum2 := 0.0

	for i := 0; i < len(arr); i++ {
		sum2 += (float64(arr[i]) - mean) * (float64(arr[i]) - mean)
	}

	variance := sum2 / float64(len(arr))

	stdDev := math.Sqrt(variance)

	return mean, stdDev
}

func getAdjustedNumbers(arr []float64, mean float64, stdDev float64) (float64, float64) {
	var zArr []float64

	for x := range arr {
		z := (float64(arr[x]) - mean) / stdDev
		if math.Abs(z) > 1 {
			zArr = append(zArr, mean+signum(z)*math.Log10(math.Abs(z))*stdDev)
		} else {
			zArr = append(zArr, arr[x])
		}
	}

	return findStdDev(zArr)
}

func (num *appNumbers) getAppNumbers(nostalgia Nostalgia) error {
	var (
		arrDau              []float64
		arrRequests         []float64
		arrResponses        []float64
		arrImpressions      []float64
		dataWindow          int
		movingAverageWindow int
	)
	dataWindow = 15
	movingAverageWindow = 3 // should be less than dataWindow

	data, err := nostalgia.FetchAppDataForRange(num.app.ID, num.app.Date, dataWindow+10)

	if err != nil {
		return err
	}

	if len(data.Result) < dataWindow {
		return errors.New("Not Sufficient data for Anomaly detection")
	}

	sort.Sort(data)

	for i := 0; i < dataWindow; i++ {
		arrDau = append(arrDau, float64(data.Result[i].Dau))
		arrRequests = append(arrRequests, float64(data.Result[i].Requests))
		arrResponses = append(arrResponses, float64(data.Result[i].Responses))
		arrImpressions = append(arrImpressions, float64(data.Result[i].Impressions))
	}

	meanDau, stdDau := findStdDev(arrDau)
	meanRequests, stdRequests := findStdDev(arrRequests)
	meanResponses, stdResponses := findStdDev(arrResponses)
	meanImpressions, stdImpressions := findStdDev(arrImpressions)

	num.meanDau, num.stdDau = getAdjustedNumbers(arrDau, meanDau, stdDau)
	num.meanRequests, num.stdRequests = getAdjustedNumbers(arrRequests, meanRequests, stdRequests)
	num.meanResponses, num.stdResponses = getAdjustedNumbers(arrResponses, meanResponses, stdResponses)
	num.meanImpressions, num.stdImpressions = getAdjustedNumbers(arrImpressions, meanImpressions, stdImpressions)

	num.app.getAppData(data, movingAverageWindow)

	return nil
}

func (app *App) getAppData(data *NostalgiaResponse, window int) {

	for i := 0; i < window; i++ {
		app.Dau += data.Result[i].Dau
		app.Requests += data.Result[i].Requests
		app.Responses += data.Result[i].Responses
		app.Impressions += data.Result[i].Impressions
	}
	app.Dau /= window
	app.Requests /= window
	app.Responses /= window
	app.Responses /= window

}
