package main

import (
	"errors"
	"math"
	"sort"
)

func findStdDev(arr []int) (float64, float64) {

	sum := 0
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

func (num *appNumbers) getAppNumbers(nostalgia Nostalgia) error {

	var arrDau []int
	var arrRequests []int
	var arrResponses []int
	var arrImpressions []int

	data, err := nostalgia.FetchAppDataForRange(num.app.ID, num.app.Date, 25)

	if err != nil {
		return err
	}

	sort.Sort(data)

	for i := 0; i < len(data.Result); i++ {
		arrDau = append(arrDau, data.Result[i].Dau)
		arrRequests = append(arrRequests, data.Result[i].Requests)
		arrResponses = append(arrResponses, data.Result[i].Responses)
		arrImpressions = append(arrImpressions, data.Result[i].Impressions)
	}

	num.meanDau, num.stdDau = findStdDev(arrDau)
	num.meanRequests, num.stdRequests = findStdDev(arrRequests)
	num.meanResponses, num.stdResponses = findStdDev(arrResponses)
	num.meanImpressions, num.stdImpressions = findStdDev(arrImpressions)

	if len(data.Result) < 3 {
		return errors.New("Not Sufficient data for Anomaly detection")
	}

	num.app.getAppData(data, 3)

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
