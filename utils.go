package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

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

func (ret *nostalgiaResponse) getNostalgiaResponse() {
	response, _ := ioutil.ReadFile("response.json")

	err := json.Unmarshal([]byte(response), &ret)
	if err != nil {
		fmt.Println(err)
	}

}

func initAnomaly() AnomalyDetector {
	svc := anomalyDetector{}
	return svc
}
