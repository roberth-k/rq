package rq

import (
	"io/ioutil"
)

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
