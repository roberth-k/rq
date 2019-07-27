package rq

import (
	"fmt"
	"net/url"
)

func (req Request) MapURL(f func(url url.URL) url.URL) Request {
	if req.err != nil {
		return req
	}

	req.URL = f(req.URL)
	return req
}

func (req Request) SetURL(value string) Request {
	if req.err != nil {
		return req
	}

	u, err := url.Parse(value)
	if err != nil {
		req.err = err
		return req
	}

	req.URL = *u
	return req
}

func (req Request) JoinURL(segments ...string) Request {
	return req.MapURL(func(url url.URL) url.URL {
		for _, segment := range segments {
			u, err := url.Parse(segment)
			if err != nil {
				req.err = err
				return url
			}

			url = *u
		}
		return url
	})
}

func (req Request) JoinFormatURL(segment string, args ...interface{}) Request {
	return req.JoinURL(fmt.Sprintf(segment, args...))
}
