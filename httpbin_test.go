package rq_test

import (
	"github.com/tetratom/rq"
)

type HTTPBinResponse struct {
	Data    string            `json:"data"`
	Headers map[string]string `json:"headers"`
}

var httpbin = rq.Begin().SetURL("https://httpbin.org")
