package rq

import (
	"context"
)

func (req Request) Map(mapper func(Request) Request) Request {
	return mapper(req)
}

func (req Request) WithUnmarshaller(unmarshaller Unmarshaller) Request {
	req.Unmarshaller = unmarshaller
	return req
}

func (req Request) URLToString() string {
	return req.URL.String()
}

func (req Request) GetError() error {
	return req.err
}

func (req Request) HasError() bool {
	return req.err != nil
}

func (req Request) SetError(err error) Request {
	req.err = err
	return req
}

func (req Request) GetContext() context.Context {
	if req.ctx == nil {
		return context.TODO()
	}

	return req.ctx
}
