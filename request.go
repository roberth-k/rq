package rq

import (
	"context"
	"net/http"
	url2 "net/url"
)

const (
	DefaultMethod = "GET"
)

func Begin() Request {
	return Request{
		Method: DefaultMethod,
		Client: http.DefaultClient,
	}
}

type Header struct {
	Name  string
	Value string
}

type RequestMiddleware func(Request) Request

type Marshaller func(Request, interface{}) ([]byte, error)

type Request struct {
	URL                 url2.URL
	Method              string
	Headers             []Header
	Client              *http.Client
	RequestMiddlewares  []RequestMiddleware
	ResponseMiddlewares []ResponseMiddleware
	Marshaller          Marshaller
	Context             context.Context
	err                 error
}

func (req Request) Do(ctx context.Context, body interface{}) (*Response, error) {
	if req.err != nil {
		return nil, req.err
	}

	req.Context = ctx

	for _, middleware := range req.RequestMiddlewares {
		req = middleware(req)
	}

	return DoRequest(req)
}

func (req Request) GET(ctx context.Context) (*Response, error) {
	req.Method = "GET"
	return req.Do(ctx, nil)
}

func (req Request) POST(ctx context.Context, v interface{}) (*Response, error) {
	panic("not implemented")
}

func (req Request) MapURL(f func(url url2.URL) url2.URL) Request {
	if req.err != nil {
		return req
	}

	req.URL = f(req.URL)
	return req
}

func (req Request) SetURL(url string) Request {
	if req.err != nil {
		return req
	}

	u, err := url2.Parse(url)
	if err != nil {
		req.err = err
		return req
	}

	req.URL = *u
	return req
}

func (req Request) JoinURL(segments ...string) Request {
	return req.MapURL(func(url url2.URL) url2.URL {
		for _, segment := range segments {
			u, err := url.Parse(segment)
			if err != nil {
				req.err = err
				return url
			}

			url = *u
		}
		return url
	})
}

func (req Request) MapHeaders(f func([]Header) []Header) Request {
	req.Headers = f(req.Headers)
	return req
}

func (req Request) WithRequestMiddlewares(middleware RequestMiddleware) Request {
	req.RequestMiddlewares = append(req.RequestMiddlewares, middleware)
	return req
}

func (req Request) WithResponseMiddleware(middleware ResponseMiddleware) Request {
	req.ResponseMiddlewares = append(req.ResponseMiddlewares, middleware)
	return req
}

func (req Request) WithMarshaller(marshaller Marshaller) Request {
	req.Marshaller = marshaller
	return req
}
