package rq

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

type RequestMiddleware func(Request) Request

// Marshaller is a function for marshalling the given value into the given
// request. The marshaller is expected to set any additional relevant properties,
// such as the Content-Type header.
//
// If marshalling fails, the error of the resultant request should be set using
// SetError().
type Marshaller func(Request, interface{}) Request

type Header struct {
	Name  string
	Value string
}

type Request struct {
	URL                 url.URL
	Method              string
	Headers             []Header
	Client              *http.Client
	Body                io.Reader
	RequestMiddlewares  []RequestMiddleware
	ResponseMiddlewares []ResponseMiddleware
	Marshaller          Marshaller
	Unmarshaller        Unmarshaller
	Context             context.Context
	err                 error
}

type ResponseMiddleware func(Request, Response, error) (Response, error)

type Unmarshaller func(Response, interface{}) error

type Response struct {
	response     *http.Response
	Unmarshaller Unmarshaller
}

func Begin(segments ...string) Request {
	return Request{}.JoinURL(segments...)
}
