// +build !go1.13

package rq

import (
	"io"
	"net/http"
)

func newHttpRequest(req *Request, r io.Reader) (*http.Request, error) {
	httpRequest, err := http.NewRequest(req.Method, req.URL.String(), r)
	if err != nil {
		return nil, err
	}

	httpRequest = httpRequest.WithContext(req.GetContext())
	return httpRequest, nil
}
