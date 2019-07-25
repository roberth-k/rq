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
	t.Parallel()
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
	Data    string            `json:"data"`
	Headers map[string]string `json:"headers"`
}

func TestBasicMarshalling(t *testing.T) {
	t.Parallel()

	type Test struct {
		Foo int
	}

	ctx := context.Background()
	input := Test{Foo: 42}
	req := httpbin.JoinURL("/anything").WithBodyFromJSON(input)
	require.Equal(t, http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}, req.HeaderMap())
	rep, err := req.POST(ctx)
	require.NoError(t, err)
	require.Equal(t, 200, rep.Status)

	var result HTTPBinResponse
	err = rep.Unmarshal(&result)
	require.NoError(t, err)
	require.Equal(t, `{"Foo":42}`, result.Data)
}

func TestRequestHeaders(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	req := httpbin.JoinURL("/anything").SetHeader("Foo", "Bar").SetHeader("Bax", "Baz")
	require.Equal(t, http.Header{"Foo": []string{"Bar"}, "Bax": []string{"Baz"}}, req.HeaderMap())
	rep, err := req.GET(ctx)
	require.NoError(t, err)
	require.Equal(t, 200, rep.Status)
	var result HTTPBinResponse
	require.NoError(t, rep.Unmarshal(&result))
	require.Equal(t, "Bar", result.Headers["Foo"])
	require.Equal(t, "Baz", result.Headers["Bax"])
}
