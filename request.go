package rq

import (
	"context"
	"net/http"
)

func (req Request) Do(ctx context.Context) (Response, error) {
	if req.err != nil {
		return Response{}, req.err
	}

	req.ctx = ctx

	for _, middleware := range req.RequestMiddlewares {
		req = middleware(req)
		if req.err != nil {
			return Response{}, req.err
		}
	}

	r, err := http.NewRequest(req.Method, req.URL.String(), req.Body)
	if err != nil {
		return Response{}, err
	}

	for _, header := range req.Headers {
		r.Header.Add(header.Name, header.Value)
	}

	r = r.WithContext(req.GetContext())

	client := req.Client
	if client == nil {
		client = http.DefaultClient
	}

	response, err := client.Do(r)
	if err != nil {
		return Response{}, err
	}

	result := Response{
		response:     response,
		Unmarshaller: req.Unmarshaller,
	}

	for _, middleware := range req.ResponseMiddlewares {
		result, err = middleware(req, result, err)
	}

	return result, err
}
