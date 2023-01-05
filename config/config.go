package config

import "os"

var (
	HttpPort     = "8080"
	MergesortUrl = ""
)

func init() {
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort != "" {
		HttpPort = httpPort
	}

	mergesortUrl := os.Getenv("MERGESORT_URL")
	if mergesortUrl != "" {
		MergesortUrl = mergesortUrl
	}
}
