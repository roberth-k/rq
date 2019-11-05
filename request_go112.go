// +build !go1.13

package rq

import (
	"io"
	"net/http"
)

func newHttpRequest(req *Request, r io.Reader) (*http.Request, error) {
	r, err := http.NewRequest(req.Method, req.URL.String(), r)
	if err != nil {
		return nil, err
	}

	r = r.WithContext(req.GetContext())
	return r, nil
}
