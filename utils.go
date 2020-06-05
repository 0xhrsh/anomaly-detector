package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func initAnomaly() AnomalyDetector {
	svc := anomalyDetector{}

	return svc
}

type config struct {
	UserID    string `required:"true"`
	AuthToken string `required:"true"`
	Endpoint  string `required:"true"`
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
	app             App
	meanDau         float64
	stdDau          float64
	meanRequests    float64
	stdRequests     float64
	meanResponse    float64
	stdResponse     float64
	meanImpressions float64
	stdImpressions  float64
}

type nostalgiaResponse struct {
	Result []App `json:"result"`
}

func (nResp *nostalgiaResponse) getNostalgiaResponse(ID string, Date time.Time, window int, conf config) error {

	client := &http.Client{}

	url := fmt.Sprintf(conf.Endpoint+"?app_id=%s&from=%s&to=%s&dim=date,app&metrics=ad_responses,impressions,dau", ID, Date.AddDate(0, 0, -1*window).Format("2006-01-02"), Date.AddDate(0, 0, -1).Format("2006-01-02"))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Id", conf.UserID)
	req.Header.Set("Auth-Token", conf.AuthToken)
	res, err := client.Do(req)

	if err != nil {
		res.Body.Close()
		return err
	}
	defer res.Body.Close()

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
