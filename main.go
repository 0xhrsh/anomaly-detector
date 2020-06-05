package main

import (
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/vrischmann/envconfig"
)

func main() {
	var conf config
	if err := envconfig.Init(&conf); err != nil {
		log.Fatalln(err)
	}

	findAnomalyHandler := httptransport.NewServer(
		makeFindAnomalyEndpoint(initAnomaly(), conf),
		decodeFindAnomalyRequest,
		encodeResponse,
	)

	http.Handle("/anm", findAnomalyHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
