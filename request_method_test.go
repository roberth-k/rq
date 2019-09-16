package rq_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/tetratom/rq"
	"strings"
	"testing"
)

func TestRequest_SetMethod(t *testing.T) {
	t.Parallel()
	req := rq.Request{}
	require.Equal(t, "", req.Method)
	req = req.SetMethod("test")
	require.Equal(t, "test", req.Method)
}

func TestBasicHTTPMethods(t *testing.T) {
	t.Parallel()

	tests := []struct {
		method   string
		function func(rq.Request, context.Context) (rq.Response, error)
	}{
		{"", rq.Request.GET},
		{"DELETE", rq.Request.DELETE},
		{"GET", rq.Request.GET},
		{"PATCH", rq.Request.PATCH},
		{"POST", rq.Request.POST},
		{"PUT", rq.Request.PUT},
	}

	for _, test := range tests {
		test := test
		t.Run(test.method, func(t *testing.T) {
			t.Parallel()

			lcmethod := strings.ToLower(test.method)
			req := HTTPBin().Path(lcmethod)
			// todo: separate test of url joins
			//require.Equal(t, "http://httpbin.org/"+lcmethod, req.URL.String())
			rep, err := test.function(req, context.TODO())
			require.NoError(t, err)
			require.Equal(t, 200, rep.Status())
		})
	}
}
