package rq_test

import (
	"github.com/tetratom/rq"
	"os"
)

type HTTPBinResponse struct {
	Authenticated bool              `json:"authenticated"`
	Data          string            `json:"data"`
	Headers       map[string]string `json:"headers"`
	Token         string            `json:"token"`
	User          string            `json:"user"`
}

func HTTPBin() rq.Request {
	return rq.Begin().SetURL(HTTPBinURL())
}

func HTTPBinURL() string {
	if url, ok := os.LookupEnv("HTTPBIN_URL"); ok {
		return url
	}

	return "http://httpbin.org"
}
