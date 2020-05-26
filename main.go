package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// AnomalyDetector provides operations on strings.
type AnomalyDetector interface {
	FindAnomaly(string) (int, error)
}

// appInfo is a concrete implementation of AnomalyDetector
type anomalyDetector struct {
	apps map[string]app
}

type app struct {
	Date     string `json:"date"`
	ID       string `json:"app"`
	Dau      int    `json:"dau"`
	Requests int    `json:"requests"`
	Response int    `json:"responses"`
}

// FindAnomaly abs
func (svc anomalyDetector) FindAnomaly(s string) (int, error) {

	if s == "" {
		return 404, nil
	}

	_, prs := svc.apps[s]
	if !prs {
		return 500, nil
	}
	return 200, ErrEmpty

	// return strings.ToUpper(s), nil
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

// For each method, we define request and response structs
type findAnomalyRequest struct {
	S string `json:"s"`
}

type findAnomalyResponse struct {
	V   int    `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func makeFindAnomalyEndpoint(svc AnomalyDetector) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(findAnomalyRequest)
		v, err := svc.FindAnomaly(req.S)
		if err != nil {
			return findAnomalyResponse{v, err.Error()}, nil
		}
		return findAnomalyResponse{v, ""}, nil
	}
}

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
func main() {
	svc := anomalyDetector{}

	findAnomalyHandler := httptransport.NewServer(
		makeFindAnomalyEndpoint(svc),
		decodeFindAnomalyRequest,
		encodeResponse,
	)

	http.Handle("/anm", findAnomalyHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
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
