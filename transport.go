package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
)

type findAnomalyRequest struct {
	ID        string `json:"id"`
	StartDate string `json:"start"`
	EndDate   string `json:"end"`
}

type appResponse struct {
	AnomalyDau         int       `json:"dau"`
	AnomalyResponses   int       `json:"responses"`
	AnomalyRequests    int       `json:"requests"`
	AnomalyImpressions int       `json:"impressions"`
	AnomalyTime        time.Time `json:"time"`
	Err                string    `json:"err,omitempty"`
}

type findAnomalyResponse struct {
	Response []appResponse `json:"response"`
	Err      string        `json:"err,omitempty"`
}

func makeFindAnomalyEndpoint(svc AnomalyDetector) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(findAnomalyRequest)
		resp, err := svc.FindAnomaly(req.ID, req.StartDate, req.EndDate)
		ret := findAnomalyResponse{
			Response: resp,
		}
		if err != nil {
			ret.Err = fmt.Sprint(err)
		}
		return ret, nil

	}
}

func decodeFindAnomalyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request findAnomalyRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
