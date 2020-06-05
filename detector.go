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

func (num *appNumbers) getAppNumbers(conf config) error {

	var arrDau []int
	var arrRequests []int
	var arrResponse []int

	var data nostalgiaResponse
	err := data.getNostalgiaResponse(num.app.ID, num.app.Date, 25, conf)

	if err != nil {
		return err
	}

	sort.Sort(data)

	for i := 0; i < len(data.Result); i++ {
		arrDau = append(arrDau, data.Result[i].Dau)
		arrRequests = append(arrRequests, data.Result[i].Requests)
		arrResponse = append(arrResponse, data.Result[i].Response)
	}

	num.meanDau, num.stdDau = findStdDev(arrDau)
	num.meanRequests, num.stdRequests = findStdDev(arrRequests)
	num.meanResponse, num.stdResponse = findStdDev(arrResponse)

	if len(data.Result) < 3 {
		return errors.New("Not Sufficient data for Anomaly detection")
	}

	num.app.getAppData(data, 3)

	return nil
}

func (app *App) getAppData(data nostalgiaResponse, window int) {

	for i := 0; i < window; i++ {
		app.Dau += data.Result[i].Dau
		app.Requests += data.Result[i].Requests
		app.Response += data.Result[i].Response
	}
	app.Dau /= window
	app.Requests /= window
	app.Response /= window

}
