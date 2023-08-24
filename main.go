package main

import (
	"os"
	"strconv"

	"github.com/official-stallion/stallion"
)

const (
	defaultPort    = "7025"
	defaultMetrics = 7026
)

func main() {
	port := defaultPort
	if value, ok := os.LookupEnv("ST_SERVER_PORT"); ok {
		port = value
	}

	metrics := defaultMetrics
	if value, ok := os.LookupEnv("ST_METRICS_PORT"); ok {
		metrics, _ = strconv.Atoi(value)
	}

	user, ok := os.LookupEnv("ST_USER")
	if !ok {
		user = ""
	}

	pass, ok := os.LookupEnv("ST_PASSWORD")
	if !ok {
		pass = ""
	}

	if err := stallion.NewServer(":"+port, metrics, user, pass); err != nil {
		panic(err)
	}
}
