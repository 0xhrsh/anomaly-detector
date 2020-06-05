package main

import (
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/vrischmann/envconfig"
)

func main() {
	var conf Config
	if err := envconfig.Init(&conf); err != nil {
		log.Fatalln(err)
	}

	findAnomalyHandler := httptransport.NewServer(
		makeFindAnomalyEndpoint(NewAnomalyDetector(conf)),
		decodeFindAnomalyRequest,
		encodeResponse,
	)

	http.Handle("/anm", findAnomalyHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
