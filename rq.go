package rq

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

type RequestMiddleware func(Request) Request

type Marshaller func(Request, interface{}) (Request, error)

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

type Unmarshaller func(response *Response, value interface{}) error

type Response struct {
	request Request
	Body    io.ReadCloser
	Headers http.Header
	Status  int
}

func Begin() Request {
	return Request{}
}
