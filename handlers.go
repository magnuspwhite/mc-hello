package main

import (
	"fmt"
	"net/http"
)

func HelloIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world, what a lovely day :)")
}

func WeatherIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "looks like it's classic british weather outside, better bring a coat/umbrella/sun screen/wellies/surf board")
}

func HealthcheckIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "healthy")
	w.WriteHeader(http.StatusOK)
}

func ReadinessIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ready")
	w.WriteHeader(http.StatusOK)
}
