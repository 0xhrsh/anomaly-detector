package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

func initAnomaly() AnomalyDetector {
	svc := anomalyDetector{}
	return svc
}

// App contains all fields of app
type App struct {
	Date     time.Time `json:"date"`
	ID       string    `json:"app"`
	Dau      int       `json:"dau"`
	Requests int       `json:"requests"`
	Response int       `json:"responses"`
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

type nostalgiaResponse struct {
	Result []App `json:"result"`
}

func (nResp *nostalgiaResponse) getNostalgiaResponse(ID string, Date time.Time, window int) {
	response, _ := ioutil.ReadFile("response.json")

	err := json.Unmarshal([]byte(response), &nResp)
	if err != nil {
		fmt.Println(err)
	}

}

func (nResp nostalgiaResponse) Len() int {
	return len(nResp.Result)
}

func (nResp nostalgiaResponse) Swap(i, j int) {
	nResp.Result[i], nResp.Result[j] = nResp.Result[j], nResp.Result[i]
}

func (nResp nostalgiaResponse) Less(i, j int) bool {

	return nResp.Result[i].Date.After(nResp.Result[j].Date)
}
