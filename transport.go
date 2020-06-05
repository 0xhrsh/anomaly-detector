package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type findAnomalyRequest struct {
	ID   string `json:"id"`
	Date string `json:"date"`
}

type findAnomalyResponse struct {
	Code int    `json:"code"`
	Err  string `json:"err,omitempty"`
}

func makeFindAnomalyEndpoint(svc AnomalyDetector, conf config) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(findAnomalyRequest)
		v, err := svc.FindAnomaly(req.ID, req.Date, conf)
		if err != nil {
			return findAnomalyResponse{v, err.Error()}, nil
		}
		return findAnomalyResponse{v, ""}, nil
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
