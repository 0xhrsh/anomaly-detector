package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func initAnomaly() AnomalyDetector {
	svc := anomalyDetector{}
	return svc
}

// App contains all fields of app
type App struct {
	Date        time.Time `json:"date"`
	ID          string    `json:"app"`
	Dau         int       `json:"dau"`
	Requests    int       `json:"requests"`
	Response    int       `json:"responses"`
	Impressions int       `json:"impressions"`
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

func (nResp *nostalgiaResponse) getNostalgiaResponse(ID string, Date time.Time, window int) error {

	client := &http.Client{}
	godotenv.Load("auth.env")

	req, err := http.NewRequest("GET", "http://go.greedygame.com/v3/nostalgia/report?from=2020-05-19&to=2020-05-25&dim=date,app&metrics=ad_responses,impressions,dau", nil) //
	if err != nil {
		return err
	}

	req.Header.Set("User-Id", os.Getenv("USER_ID"))
	req.Header.Set("Auth-Token", os.Getenv("AUTH_TOKEN"))
	res, _ := client.Do(req)
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&nResp)

	if err != nil {
		return err
	}

	return nil

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
