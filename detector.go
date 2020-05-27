package main

import (
	"database/sql"
	"fmt"
	"log"
	"math"
)

var readNumbers = `SELECT *
FROM data
WHERE id = ($1);`

var readApp = `SELECT * FROM data where id = ($1) and date = ($2);`

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
	meanDau      float64
	stdDau       float64
	meanRequests float64
	stdRequests  float64
	meanResponse float64
	stdResponse  float64
}

func getAppData(ID string, Date string, db *sql.DB) (appNumbers, App) {

	rows, err := db.Query(readNumbers, ID)
	if err != nil {
		log.Panic(err)
	}

	var arrDau []int
	var arrRequests []int
	var arrResponse []int
	var ret appNumbers
	var app App

	for rows.Next() {
		err = rows.Scan(&app.Date, &app.ID, &app.Dau, &app.Requests, &app.Response)
		arrDau = append(arrDau, app.Dau)
		arrRequests = append(arrRequests, app.Requests)
		arrResponse = append(arrResponse, app.Response)
	}
	rows.Close()

	ret.meanDau, ret.stdDau = findStdDev(arrDau)
	ret.meanRequests, ret.stdRequests = findStdDev(arrRequests)
	ret.meanResponse, ret.stdResponse = findStdDev(arrResponse)

	row := db.QueryRow(readApp, ID, Date)
	err = row.Scan(&app.Date, &app.ID, &app.Dau, &app.Requests, &app.Response)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(err)
		}
	}

	return ret, app
}
