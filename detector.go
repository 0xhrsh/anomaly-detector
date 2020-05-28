package main

import (
	"database/sql"
	"errors"
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

func (num *appNumbers) getAppNumbers(db *sql.DB) error {
	var readNumbers = `SELECT * FROM data WHERE id = ($1) and date < ($2);`

	rows, err := db.Query(readNumbers, num.app.ID, num.app.Date)
	if err != nil {
		return err
	}

	var app App
	var arrDau []int
	var arrRequests []int
	var arrResponse []int

	for rows.Next() {
		err = rows.Scan(&app.Date, &app.ID, &app.Dau, &app.Requests, &app.Response)
		arrDau = append(arrDau, app.Dau)
		arrRequests = append(arrRequests, app.Requests)
		arrResponse = append(arrResponse, app.Response)
	}
	rows.Close()

	num.meanDau, num.stdDau = findStdDev(arrDau)
	num.meanRequests, num.stdRequests = findStdDev(arrRequests)
	num.meanResponse, num.stdResponse = findStdDev(arrResponse)

	return nil
}

func (app *App) getAppData(db *sql.DB) error {

	var readApp = `SELECT * FROM data where id = ($1) and date = ($2);`

	row := db.QueryRow(readApp, app.ID, app.Date)
	err := row.Scan(&app.Date, &app.ID, &app.Dau, &app.Requests, &app.Response)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("No Data for App: " + app.ID + " for date: " + app.Date)
		}
		return err
	}
	return nil
}
