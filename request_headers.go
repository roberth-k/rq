package rq

import (
	"net/http"
)

func (req Request) HeaderMap() http.Header {
	m := make(map[string][]string, len(req.Headers))
	for _, header := range req.Headers {
		m[header.Name] = append(m[header.Name], header.Value)
	}
	return m
}

func (req Request) AddHeader(name string, value string) Request {
	req.Headers = append(req.Headers, Header{Name: name, Value: value})
	return req
}

func (req Request) SetHeader(name string, value string) Request {
	return req.RemoveHeader(name).AddHeader(name, value)
}

func (req Request) RemoveHeader(name string) Request {
	newHeaders := make([]Header, 0, len(req.Headers))

	for _, header := range req.Headers {
		if header.Name == name {
			continue
		}

		newHeaders = append(newHeaders, header)
	}

	if len(newHeaders) == len(req.Headers) {
		// don't replace the old slice unnecessarily
		return req
	}

	req.Headers = newHeaders
	return req
}
