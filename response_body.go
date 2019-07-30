package rq

import (
	"io/ioutil"
)

func (resp Response) Close() error {
	return resp.Underlying.Body.Close()
}

func (resp Response) ReadAll() ([]byte, error) {
	defer resp.Close()
	return ioutil.ReadAll(resp.Underlying.Body)
}

func (resp Response) ReadString() (string, error) {
	data, err := resp.ReadAll()
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (resp Response) ReadStringOrPanic() string {
	data, err := resp.ReadString()
	if err != nil {
		panic(err)
	}

	return data
}

func (resp Response) UnmarshalJSON(v interface{}) error {
	return UnmarshalJSON(resp, v)
}

func (resp Response) Unmarshal(v interface{}) error {
	if resp.Unmarshaller == nil {
		return resp.UnmarshalJSON(v)
	}

	return resp.Unmarshaller(resp, v)
}
