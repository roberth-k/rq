package rq

import (
	"context"
)

func (req Request) SetMethod(method string) Request {
	req.Method = method
	return req
}

func (req Request) DELETE(ctx context.Context) (Response, error) {
	req.Method = "DELETE"
	return req.Do(ctx)
}

func (req Request) GET(ctx context.Context) (Response, error) {
	req.Method = "GET"
	return req.Do(ctx)
}

func (req Request) PATCH(ctx context.Context) (Response, error) {
	req.Method = "PATCH"
	return req.Do(ctx)
}

func (req Request) POST(ctx context.Context) (Response, error) {
	req.Method = "POST"
	return req.Do(ctx)
}

func (req Request) PUT(ctx context.Context) (Response, error) {
	req.Method = "PUT"
	return req.Do(ctx)
}
