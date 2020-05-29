package main

import (
	"math"
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

type appNumbers struct {
	app          App
	meanDau      float64
	stdDau       float64
	meanRequests float64
	stdRequests  float64
	meanResponse float64
	stdResponse  float64
}

func (num *appNumbers) getAppNumbers() error {

	var arrDau []int
	var arrRequests []int
	var arrResponse []int

	var data nostalgiaResponse
	data.getNostalgiaResponse()

	for i := 0; i < len(data.Result); i++ {
		arrDau = append(arrDau, data.Result[i].Dau)
		arrRequests = append(arrRequests, data.Result[i].Requests)
		arrResponse = append(arrResponse, data.Result[i].Response)
	}

	num.meanDau, num.stdDau = findStdDev(arrDau)
	num.meanRequests, num.stdRequests = findStdDev(arrRequests)
	num.meanResponse, num.stdResponse = findStdDev(arrResponse)

	return nil
}

func (app *App) getAppData() error {
	var arrDau []int
	var arrRequests []int
	var arrResponse []int

	var data nostalgiaResponse
	data.getNostalgiaResponse()

	for i := 0; i < len(data.Result); i++ {
		arrDau = append(arrDau, data.Result[i].Dau)
		arrRequests = append(arrRequests, data.Result[i].Requests)
		arrResponse = append(arrResponse, data.Result[i].Response)
	}

	Dau, _ := findStdDev(arrDau)
	Requests, _ := findStdDev(arrRequests)
	Response, _ := findStdDev(arrResponse)

	app.Dau = int(Dau)
	app.Requests = int(Requests)
	app.Response = int(Response)

	return nil
}
