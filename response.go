package rq

import (
	"io"
	"io/ioutil"
	"net/http"
)

type ResponseMiddleware func(Request, Response, error) (Response, error)

type Response struct {
	request Request
	Body    io.ReadCloser
	Headers http.Header
	Status  int
}

func (resp *Response) Close() error {
	return resp.Body.Close()
}

func (resp *Response) ReadAll() ([]byte, error) {
	defer resp.Close()
	return ioutil.ReadAll(resp.Body)
}

func (resp *Response) UnmarshalJSON(v interface{}) error {
	return UnmarshalJSON(resp, v)
}

func (resp *Response) Unmarshal(v interface{}) error {
	return resp.request.Unmarshaller(resp, v)
}

func (resp *Response) Is1xx() bool {
	return resp.Status >= 100 && resp.Status < 200
}

func (resp *Response) Is2xx() bool {
	return resp.Status >= 200 && resp.Status < 300
}

func (resp *Response) Is3xx() bool {
	return resp.Status >= 300 && resp.Status < 400
}

func (resp *Response) Is4xx() bool {
	return resp.Status >= 400 && resp.Status < 500
}

func (resp *Response) Is5xx() bool {
	return resp.Status >= 500 && resp.Status < 600
}

func (resp *Response) GetHeader(name string) string {
	return resp.Headers.Get(name)
}

func (resp *Response) HasHeader(name string) bool {
	return resp.Headers.Get(name) != ""
}

func (resp *Response) LookupHeader(name string) (string, bool) {
	value := resp.Headers.Get(name)
	return value, value != ""
}
