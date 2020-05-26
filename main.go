package main

import (
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

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
