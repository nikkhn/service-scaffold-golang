package main

import  "net/http"

type server struct{}

var version = "unknown"
var buildDate = "unknown"

func main() {
	healthzHandler := HealthZ{}
    http.Handle("/healthz", &healthzHandler)
    http.ListenAndServe(":8000", nil)
}