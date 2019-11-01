package rq

import (
	"context"
	"io"
	"net/http"
)

func (req Request) Map(mapper func(Request) Request) Request {
	return mapper(req)
}

func (req Request) SetUnmarshaller(unmarshaller Unmarshaller) Request {
	req.Unmarshaller = unmarshaller
	return req
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

func (req Request) GetContext() context.Context {
	if req.ctx == nil {
		return context.TODO()
	}

	return req.ctx
}

// Prepare builds the native http.Request that will be used for the HTTP request.
func (req Request) Prepare(ctx context.Context) (*http.Request, error) {
	if req.err != nil {
		return nil, req.err
	}

	req.ctx = ctx

	for _, middleware := range req.RequestMiddlewares {
		req = middleware(req)
		if req.err != nil {
			return nil, req.err
		}
	}

	var reader io.Reader
	if req.Body != nil {
		reader = req.Body.Reader()
	}

	r, err := http.NewRequest(req.Method, req.URL.String(), reader)
	if err != nil {
		return nil, err
	}

	for _, header := range req.Headers {
		r.Header.Add(header.Name, header.Value)
	}

	r = r.WithContext(req.GetContext())

	return r, nil
}

func (req Request) Do(ctx context.Context) (Response, error) {
	r, err := req.Prepare(ctx)
	if err != nil {
		return Response{}, err
	}

	client := req.Client
	if client == nil {
		client = http.DefaultClient
	}

	response, err := client.Do(r)
	if err != nil {
		return Response{}, err
	}

	result := Response{
		Underlying:   response,
		Unmarshaller: req.Unmarshaller,
	}

	for _, middleware := range req.ResponseMiddlewares {
		result, err = middleware(req, result, err)
	}

	return result, err
}
