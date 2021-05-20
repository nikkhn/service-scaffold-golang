package main

import  "net/http"

type server struct{}

var version = "unknown"
var buildDate = "unknown"

func main() {
	healthzHandler := HealthZ{}
	ima := ImageMatching{}
    http.Handle("/healthz", &healthzHandler)
	http.Handle("/predict", &ima)
    http.ListenAndServe(":8000", nil)
}