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
	ID   string `json:"id"`
	Date string `json:"date"`
}

type findAnomalyResponse struct {
	AnomalyDau         int       `json:"dau,omitempty"`
	AnomalyResponses   int       `json:"responses,omitempty"`
	AnomalyRequests    int       `json:"requests,omitempty"`
	AnomalyImpressions int       `json:"impressions,omitempty"`
	AnomalyTime        time.Time `json:"time,omitempty"`
	Err                string    `json:"err,omitempty"`
}

func makeFindAnomalyEndpoint(svc AnomalyDetector) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(findAnomalyRequest)
		resp, err := svc.FindAnomaly(req.ID, req.Date)
		if err != nil {
			resp.Err = fmt.Sprint(err)
			return resp, nil
		}
		return resp, nil
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
