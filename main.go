package main

import (
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {

	findAnomalyHandler := httptransport.NewServer(
		makeFindAnomalyEndpoint(initAnomaly()),
		decodeFindAnomalyRequest,
		encodeResponse,
	)

	http.Handle("/anm", findAnomalyHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
