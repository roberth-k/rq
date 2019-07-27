package rq

import (
	"context"
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

func (req *Request) getContextOrDefault() context.Context {
	if req.Context == nil {
		return context.TODO()
	}

	return req.Context
}

func (req Request) Map(mapper func(Request) Request) Request {
	return mapper(req)
}

func (req Request) WithUnmarshaller(unmarshaller Unmarshaller) Request {
	req.Unmarshaller = unmarshaller
	return req
}

func (req Request) URLToString() string {
	return req.URL.String()
}

func (req Request) GetError() error {
	return req.err
}

func (req Request) HasError() bool {
	return req.err != nil
}

func (req Request) SetError(err error) Request {
	req.err = err
	return req
}

func (req Request) SetContext(c context.Context) Request {
	req.Context = c
	return req
}
