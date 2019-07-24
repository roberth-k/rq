package rq

import (
	"encoding/json"
	"io"
	"io/ioutil"
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
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.Unmarshal(data, v)
}
