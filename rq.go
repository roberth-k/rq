// rq is a wrapper for go's http client.
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
	err                 error
	ctx                 context.Context
}

type ResponseMiddleware func(Request, Response, error) (Response, error)

type Unmarshaller func(Response, interface{}) error

type Response struct {
	Underlying   *http.Response
	Unmarshaller Unmarshaller
}

// Begin creates an otherwise empty rq.Request with the given base URL.
func Begin(url string) Request {
	return Request{}.SetURL(url)
}

// SetURL is an alias for Begin().
func SetURL(url string) Request {
	return Begin(url)
}
