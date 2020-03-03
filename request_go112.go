// +build !go1.13

package rq

import (
	"io"
	"net/http"
)

func newHttpRequest(req *Request, r io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(req.Method, req.URL.String(), r)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(req.GetContext())
	return req, nil
}
