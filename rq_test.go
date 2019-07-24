package rq_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/tetratom/rq"
	"strings"
	"testing"
)

var httpbin = rq.Begin().SetURL("https://httpbin.org")

func TestSimpleMethods(t *testing.T) {
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
