package rq

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
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

type Marshaller func(Request, interface{}) (io.Reader, error)

type Request struct {
	URL                 url2.URL
	Method              string
	Headers             []Header
	Client              *http.Client
	Body                io.Reader
	RequestMiddlewares  []RequestMiddleware
	ResponseMiddlewares []ResponseMiddleware
	Marshaller          Marshaller
	Context             context.Context
	err                 error
}

func (req Request) Do(ctx context.Context) (*Response, error) {
	if req.err != nil {
		return nil, req.err
	}

	req.Context = ctx

	for _, middleware := range req.RequestMiddlewares {
		req = middleware(req)
	}

	return DoRequest(req)
}

func (req Request) DELETE(ctx context.Context) (*Response, error) {
	req.Method = "DELETE"
	return req.Do(ctx)
}

func (req Request) GET(ctx context.Context) (*Response, error) {
	req.Method = "GET"
	return req.Do(ctx)
}

func (req Request) PATCH(ctx context.Context) (*Response, error) {
	req.Method = "PATCH"
	return req.Do(ctx)
}

func (req Request) POST(ctx context.Context) (*Response, error) {
	req.Method = "POST"
	return req.Do(ctx)
}

func (req Request) PUT(ctx context.Context) (*Response, error) {
	req.Method = "PUT"
	return req.Do(ctx)
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

func (req Request) SetHeader(name string, values ...string) Request {
	panic("not implemented")
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

func (req Request) Marshal(v interface{}) Request {
	reader, err := req.Marshaller(req, v)
	if err != nil {
		req.err = err
		return req
	}

	req.Body = reader
	return req
}

func (req Request) WithBody(reader io.Reader) Request {
	req.Body = reader
	return req
}

func (req Request) WithBodyAsJSON(v interface{}) Request {
	req.Marshaller = func(_ Request, v interface{}) (io.Reader, error) {
		data, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		return bytes.NewReader(data), nil
	}

	req = req.Marshal(v)
	req = req.SetHeader("Content-Type", "application/json; charset=utf-8")
	return req
}
