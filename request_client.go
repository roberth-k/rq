package rq

import (
	"net/http"
)

func (req Request) SetClient(client *http.Client) Request {
	req.Client = client
	return req
}
