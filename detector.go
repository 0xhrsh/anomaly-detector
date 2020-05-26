package main

import (
	"log"
	"math"
)

var readSQL = `SELECT *
FROM data
WHERE id = ($1);`

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

func getAppData(ID string) []float64 {
	db := connectToServer()
	defer db.Close()

	rows, err := db.Query(readSQL, ID)
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()

	var arrDau []int
	var arrRequests []int
	var arrResponse []int

	for rows.Next() {
		var obj app
		err = rows.Scan(&obj.Date, &obj.ID, &obj.Dau, &obj.Requests, &obj.Response)
		arrDau = append(arrDau, obj.Dau)
		arrRequests = append(arrRequests, obj.Requests)
		arrResponse = append(arrResponse, obj.Response)
	}

	var ret []float64

	a, b := findStdDev(arrDau)
	ret = append(ret, a)
	ret = append(ret, b)

	a, b = findStdDev(arrRequests)
	ret = append(ret, a)
	ret = append(ret, b)

	a, b = findStdDev(arrResponse)
	ret = append(ret, a)
	ret = append(ret, b)

	return ret
}
