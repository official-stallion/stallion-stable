package main

import (
	"os"

	"github.com/official-stallion/stallion"
)

const (
	defaultPort = "7025"
	defaultMetrics = "7026"
)

func main() {
	port := defaultPort
	if value, ok := os.LookupEnv("ST_SERVER_PORT"); ok {
		port = value
	}

	metrics := defaultMetrics
	if value, ok := os.LookupEnv("ST_METRICS_PORT"); ok {
		metrics = value
	} 

	if err := stallion.NewServer(":" + port, metrics); err != nil {
		panic(err)
	}
}
