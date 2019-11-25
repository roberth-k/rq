package rq

import (
	"fmt"
	"net/url"
	"strings"
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

func (req Request) Path(segments ...string) Request {
	return req.MapURL(func(url url.URL) url.URL {
		s := strings.Join(segments, "/")

		if strings.HasPrefix(s, "/") {
			url.Path = s
		} else {
			url.Path += "/" + s
		}

		return url
	})
}

func (req Request) Pathf(segment string, args ...interface{}) Request {
	return req.Path(fmt.Sprintf(segment, args...))
}

func (req Request) AddQueryf(name string, value string, args ...interface{}) Request {
	if len(args) > 0 {
		value = fmt.Sprintf(value, args...)
	}

	query := req.URL.Query()
	query.Add(name, value)
	req.URL.RawQuery = query.Encode()
	return req
}

func (req Request) GetQuery(name string) string {
	return req.URL.Query().Get(name)
}

func (req Request) RemoveQuery(name string) Request {
	query := req.URL.Query()
	query.Del(name)
	req.URL.RawQuery = query.Encode()
	return req
}

func (req Request) SetQueryf(name string, value string, args ...interface{}) Request {
	if len(args) > 0 {
		value = fmt.Sprintf(value, args...)
	}

	query := req.URL.Query()
	query.Set(name, value)
	req.URL.RawQuery = query.Encode()
	return req
}

func (req Request) ReplaceQuery(query url.Values) Request {
	req.URL.RawQuery = query.Encode()
	return req
}

func (req Request) GetFragment() string {
	return req.URL.Fragment
}

func (req Request) SetFragment(fragment string) Request {
	req.URL.Fragment = fragment
	return req
}
