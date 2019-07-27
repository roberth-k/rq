package rq

import (
	"encoding/json"
)

func NOOPMarshaller(req Request, _ interface{}) Request {
	return req
}

func JSONMarshaller(req Request, value interface{}) Request {
	data, err := json.Marshal(value)
	if err != nil {
		return req.SetError(err)
	}

	return req.
		SetBodyBytes(data).
		SetHeader("Content-Type", "application/json; charset=utf-8")
}
