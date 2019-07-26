package rq

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type ResponseMiddleware func(Response, error) (Response, error)

type Response struct {
	Body    io.ReadCloser
	Headers http.Header
	Status  int
}

func NewResponse(response *http.Response) *Response {
	var rep Response
	rep.Body = response.Body
	rep.Headers = response.Header
	rep.Status = response.StatusCode
	return &rep
}

func (resp *Response) Close() error {
	return resp.Body.Close()
}

func (resp *Response) ReadAll() ([]byte, error) {
	defer resp.Close()
	return ioutil.ReadAll(resp.Body)
}

func (resp *Response) IsBodyJSON() bool {
	return strings.HasPrefix(resp.GetHeader("content-type"), "application/json")
}

func (resp *Response) Unmarshal(v interface{}) error {
	data, err := resp.ReadAll()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
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
