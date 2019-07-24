package rq

import (
	"net/http"
)

func DoRequest(req Request) (*Response, error) {
	// todo; optional request body
	r, err := http.NewRequest(req.Method, req.URL.String(), req.Body)
	if err != nil {
		return nil, err
	}

	r = r.WithContext(req.Context)

	response, err := req.Client.Do(r)
	if err != nil {
		return nil, err
	}

	return NewResponse(response), nil
}
