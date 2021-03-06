package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"HelloIndex",
		"GET",
		"/",
		HelloIndex,
	},
	Route{
		"WeatherIndex",
		"GET",
		"/weather",
		WeatherIndex,
	},
	Route{
		"HealthcheckIndex",
		"GET",
		"/healthcheck",
		HealthcheckIndex,
	},
	Route{
		"ReadinessIndex",
		"GET",
		"/readiness",
		ReadinessIndex,
	},
}
