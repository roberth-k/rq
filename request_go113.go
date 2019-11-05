// +build go1.13

package rq

import (
	"io"
	"net/http"
)

func newHttpRequest(req *Request, r io.Reader) (*http.Request, error) {
	return http.NewRequestWithContext(req.GetContext(), req.Method, req.URL.String(), r)
}
