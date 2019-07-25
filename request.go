package rq

import (
	"bytes"
	"context"
	"io"
	"net/http"
	url2 "net/url"
	"time"
)

func Begin() Request {
	return Request{
		Method:     "GET",
		Client:     http.DefaultClient,
		Marshaller: JSONMarshaller,
		Context:    context.TODO(),
	}
}

type Header struct {
	Name  string
	Value string
}

type RequestMiddleware func(Request) Request

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

func (req Request) HeaderMap() http.Header {
	m := make(map[string][]string, len(req.Headers))
	for _, header := range req.Headers {
		m[header.Name] = append(m[header.Name], header.Value)
	}
	return m
}

func (req Request) SetHeader(name string, values ...string) Request {
	// todo: remove existing values

	for _, value := range values {
		req.Headers = append(req.Headers, Header{
			Name:  name,
			Value: value,
		})
	}

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

func (req Request) Marshal(v interface{}) Request {
	req1, err := req.Marshaller(req, v)
	if err != nil {
		req.err = err
		return req
	}

	return req1
}

func (req Request) WithBody(reader io.Reader) Request {
	req.Body = reader
	return req
}

func (req Request) WithBodyAsBytes(data []byte) Request {
	req.Body = bytes.NewReader(data)
	return req
}

func (req Request) WithBodyFromJSON(v interface{}) Request {
	req.Marshaller = JSONMarshaller
	return req.Marshal(v)
}

func (req Request) MapClient(f func(*http.Client) *http.Client) Request {
	req.Client = f(req.Client)
	return req
}

func (req Request) WithTimeout(timeout time.Duration) Request {
	return req.MapClient(func(client *http.Client) *http.Client {
		client.Timeout = timeout
		return client
	})
}
