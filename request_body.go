package rq

import (
	"bytes"
	"io"
	"net/url"
	"strings"
)

// RequestBody produces readers for requests. Any time a request is executed,
// the body is set to be the reader returned by Reader(). If a RequestBody is
// to be reused, the implementation must ensure that every call to Reader()
// returns a new io.Reader that starts from the beginning.
type RequestBody interface {
	Reader() io.Reader
}

// RequestBodyBytes is a reusable request body of bytes.
type RequestBodyBytes struct {
	Data []byte
}

func (rbb *RequestBodyBytes) Reader() io.Reader {
	return bytes.NewReader(rbb.Data)
}

type RequestBodyString struct {
	Data string
}

func (rbb *RequestBodyString) Reader() io.Reader {
	return strings.NewReader(rbb.Data)
}

type RequestBodyReader struct {
	BodyReader io.Reader
}

func (rbb *RequestBodyReader) Reader() io.Reader {
	return rbb.BodyReader
}

func (req Request) SetMarshaller(marshaller Marshaller) Request {
	req.Marshaller = marshaller
	return req
}

func (req Request) SetBody(body RequestBody) Request {
	req.Body = body
	return req
}

func (req Request) SetBodyBytes(data []byte) Request {
	req.Body = &RequestBodyBytes{data}
	return req
}

func (req Request) SetBodyString(data string) Request {
	req.Body = &RequestBodyString{data}
	return req
}

func (req Request) SetBodyReader(r io.Reader) Request {
	req.Body = &RequestBodyReader{r}
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

func (req Request) MarshalForm(value url.Values) Request {
	req.Marshaller = FormMarshaller
	return req.Marshal(value)
}
