package rq

import (
	"bytes"
)

func (req Request) Marshal(v interface{}) Request {
	req1, err := req.getMarshallerOrDefault()(req, v)
	if err != nil {
		req.err = err
		return req
	}

	return req1
}

func (req Request) WithBodyAsBytes(data []byte) Request {
	req.Body = bytes.NewReader(data)
	return req
}

func (req Request) WithBodyFromJSON(v interface{}) Request {
	req.Marshaller = JSONMarshaller
	return req.Marshal(v)
}
