package rq

import (
	"net/http"
)

func DoRequest(req Request) (*Response, error) {
	r, err := http.NewRequest(req.Method, req.URL.String(), req.Body)
	if err != nil {
		return nil, err
	}

	for _, header := range req.Headers {
		r.Header.Add(header.Name, header.Value)
	}

	r = r.WithContext(req.getContextOrDefault())

	result, err := req.getClientOrDefault().Do(r)
	if err != nil {
		return nil, err
	}

	response := Response{
		request: req,
		Body:    result.Body,
		Headers: result.Header,
		Status:  result.StatusCode,
	}

	for _, middleware := range req.ResponseMiddlewares {
		response, err = middleware(response, err)
	}

	return &response, err
}
