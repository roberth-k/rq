package rq

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/textproto"
	"strings"
)

func (req Request) HeaderMap() http.Header {
	m := make(map[string][]string, len(req.Headers))
	for _, header := range req.Headers {
		m[header.Name] = append(m[header.Name], header.Value)
	}
	return m
}

func (req Request) AddHeader(name string, value string, args ...interface{}) Request {
	name = textproto.CanonicalMIMEHeaderKey(name)

	if len(args) > 0 {
		value = fmt.Sprintf(value, args...)
	}

	req.Headers = append(req.Headers, Header{Name: name, Value: value})
	return req
}

func (req Request) GetHeader(name string) string {
	name = textproto.CanonicalMIMEHeaderKey(name)
	values := make([]string, 0, 1)

	for _, header := range req.Headers {
		if header.Name == name {
			values = append(values, header.Value)
		}
	}

	return strings.Join(values, "; ")
}

func (req Request) SetHeader(name string, value string, args ...interface{}) Request {
	return req.RemoveHeader(name).AddHeader(name, value, args...)
}

func (req Request) RemoveHeader(name string) Request {
	name = textproto.CanonicalMIMEHeaderKey(name)
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

func (req Request) SetBasicAuth(username, password string) Request {
	concatenated := username + ":" + password
	credentials := base64.StdEncoding.EncodeToString([]byte(concatenated))
	return req.SetHeader("Authorization", "Basic %s", credentials)
}

func (req Request) SetBearerToken(token string) Request {
	return req.SetHeader("Authorization", "Bearer %s", token)
}
