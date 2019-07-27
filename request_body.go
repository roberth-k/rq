package rq

import (
	"bytes"
	"io"
	"strings"
)

func (req Request) SetMarshaller(marshaller Marshaller) Request {
	req.Marshaller = marshaller
	return req
}

func (req Request) SetBody(reader io.Reader) Request {
	req.Body = reader
	return req
}

func (req Request) SetBodyBytes(data []byte) Request {
	req.Body = bytes.NewReader(data)
	return req
}

func (req Request) SetBodyString(data string) Request {
	req.Body = strings.NewReader(data)
	return req
}

// Marshal sets the body of the request by marshalling the given value using
// the Marshaller set on the request object. If none is set, it is marshalled
// using the MarshalJSON() method of the request.
func (req Request) Marshal(value interface{}) Request {
	if req.Marshaller == nil {
		return req.MarshalJSON(value)
	}

	return req.Marshaller(req, value)
}

// MarshalJSON marshals the given value using the default JSON marshaller.
// The Content-Type of the request will be set to "application/json; charset=utf-8".
func (req Request) MarshalJSON(value interface{}) Request {
	req.Marshaller = JSONMarshaller
	return req.Marshal(value)
}
