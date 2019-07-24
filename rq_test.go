package rq_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/tetratom/rq"
	"net/http"
	"strings"
	"testing"
)

var httpbin = rq.Begin().SetURL("https://httpbin.org")

func TestBasicHTTPMethods(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		method string
		f      func(request rq.Request) (*rq.Response, error)
	}{
		{"DELETE", func(req rq.Request) (*rq.Response, error) { return req.DELETE(ctx) }},
		{"GET", func(req rq.Request) (*rq.Response, error) { return req.GET(ctx) }},
		{"PATCH", func(req rq.Request) (*rq.Response, error) { return req.PATCH(ctx) }},
		{"POST", func(req rq.Request) (*rq.Response, error) { return req.POST(ctx) }},
		{"PUT", func(req rq.Request) (*rq.Response, error) { return req.PUT(ctx) }},
	}

	for _, test := range tests {
		test := test
		t.Run(test.method, func(t *testing.T) {
			t.Parallel()

			lcmethod := strings.ToLower(test.method)
			req := httpbin.JoinURL(lcmethod)
			require.Equal(t, "https://httpbin.org/"+lcmethod, req.URL.String())
			rep, err := test.f(req)
			require.NoError(t, err)
			require.Equal(t, 200, rep.Status)
		})
	}
}

type HTTPBinResponse struct {
	Data string `json:"data"`
}

func TestBasicMarshalling(t *testing.T) {
	type Test struct {
		Foo int
	}

	ctx := context.Background()
	input := Test{Foo: 42}
	req := httpbin.JoinURL("/anything").WithBodyAsJSON(input)
	require.Equal(t, http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}, req.HeaderMap())
	rep, err := req.POST(ctx)
	require.NoError(t, err)
	require.Equal(t, 200, rep.Status)

	var result HTTPBinResponse
	err = rep.Unmarshal(&result)
	require.NoError(t, err)
	require.Equal(t, `{"Foo":42}`, result.Data)
}
