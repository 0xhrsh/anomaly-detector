package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func initAnomaly() AnomalyDetector {
	svc := anomalyDetector{}
	return svc
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
