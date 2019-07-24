package rq

import (
	"io"
	"net/http"
)

type ResponseMiddleware func(Response, error) (Response, error)

type Response struct {
	Body   io.ReadCloser
	Status int
}

func NewResponse(response *http.Response) *Response {
	var rep Response
	rep.Body = response.Body
	rep.Status = response.StatusCode
	return &rep
}

func (resp *Response) Unmarshal(v interface{}) error {
	panic("not implemented")
}
