package rq

import (
	"bytes"
	"encoding/json"
)

func NOOPMarshaller(req Request, v interface{}) (Request, error) {
	return req, nil
}

func JSONMarshaller(req Request, v interface{}) (Request, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return req, err
	}

	req.Body = bytes.NewReader(data)
	req = req.SetHeader("Content-Type", "application/json; charset=utf-8")
	return req, nil
}
