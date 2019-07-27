package rq

import (
	"net/http"
)

func (req Request) MapHeaders(f func([]Header) []Header) Request {
	req.Headers = f(req.Headers)
	return req
}

func (req Request) HeaderMap() http.Header {
	m := make(map[string][]string, len(req.Headers))
	for _, header := range req.Headers {
		m[header.Name] = append(m[header.Name], header.Value)
	}
	return m
}

func (req Request) SetHeader(name string, values ...string) Request {
	// todo: remove existing values

	for _, value := range values {
		req.Headers = append(req.Headers, Header{
			Name:  name,
			Value: value,
		})
	}

	return req
}
