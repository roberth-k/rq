// rq is a wrapper for go's http client.
package rq

import (
	"context"
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
	Body                RequestBody
	RequestMiddlewares  []RequestMiddleware
	ResponseMiddlewares []ResponseMiddleware
	Marshaller          Marshaller
	Unmarshaller        Unmarshaller
	err                 error
	ctx                 context.Context
}

// ResponseMiddleware intercepts a response. It is also called in case of a
// request failure, in which case `rep` may be empty, and `err` non-nil. The
// middleware is expected to return (rep, err) immediately if it is not designed
// to handle an erroneous response.
//
// The `req` passed to the middleware is exactly the same as was used for the
// original request. Use req.GetContext() to obtain the request context.
type ResponseMiddleware func(req Request, rep Response, err error) (Response, error)

type Unmarshaller func(Response, interface{}) error

type Response struct {
	Underlying        *http.Response
	Unmarshaller      Unmarshaller
	IgnoreContentType bool
}

// Begin creates an otherwise empty rq.Request with the given base URL.
func Begin(url string) Request {
	return Request{}.SetURL(url)
}

// SetURL is an alias for Begin().
func SetURL(url string) Request {
	return Begin(url)
}
