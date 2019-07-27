package rq

import (
	"context"
	"io"
	"net/http"
)

func (req *Request) getClientOrDefault() *http.Client {
	if req.Client == nil {
		return http.DefaultClient
	}

	return req.Client
}

func (req *Request) getMarshallerOrDefault() Marshaller {
	if req.Marshaller == nil {
		return JSONMarshaller
	}

	return req.Marshaller
}

func (req *Request) getUnmarshallerOrDefault() Unmarshaller {
	if req.Unmarshaller == nil {
		return UnmarshalJSON
	}

	return req.Unmarshaller
}

func (req *Request) getContextOrDefault() context.Context {
	if req.Context == nil {
		return context.TODO()
	}

	return req.Context
}

func (req Request) WithMarshaller(marshaller Marshaller) Request {
	req.Marshaller = marshaller
	return req
}

func (req Request) WithUnmarshaller(unmarshaller Unmarshaller) Request {
	req.Unmarshaller = unmarshaller
	return req
}

func (req Request) WithBody(reader io.Reader) Request {
	req.Body = reader
	return req
}

func (req Request) URLToString() string {
	return req.URL.String()
}
