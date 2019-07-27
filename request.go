package rq

import (
	"context"
	"net/http"
)

func (req Request) Do(ctx context.Context) (*Response, error) {
	if req.err != nil {
		return nil, req.err
	}

	req.Context = ctx

	for _, middleware := range req.RequestMiddlewares {
		req = middleware(req)
	}

	r, err := http.NewRequest(req.Method, req.URL.String(), req.Body)
	if err != nil {
		return nil, err
	}

	for _, header := range req.Headers {
		r.Header.Add(header.Name, header.Value)
	}

	r = r.WithContext(req.getContextOrDefault())

	response, err := req.getClientOrDefault().Do(r)
	if err != nil {
		return nil, err
	}

	result := Response{
		response:     response,
		Unmarshaller: req.Unmarshaller,
	}

	for _, middleware := range req.ResponseMiddlewares {
		result, err = middleware(req, result, err)
	}

	return &result, err
}
