package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type findAnomalyRequest struct {
	ID        string `json:"id"`
	StartDate string `json:"start"`
	EndDate   string `json:"end"`
}

// AppResponse contains the information of anomaly for an app for a particular date
type AppResponse struct {
	AnomalyDau         int          `json:"dau"`
	AnomalyResponses   int          `json:"responses"`
	AnomalyRequests    int          `json:"requests"`
	AnomalyImpressions int          `json:"impressions"`
	AnomalyTime        string       `json:"time"`
	CodeChanges        []CommitInfo `json:"commits,omitempty"`
	Err                string       `json:"err,omitempty"`
}

type findAnomalyResponse struct {
	Response []AppResponse `json:"response"`
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
