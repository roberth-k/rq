package rq_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/tetratom/rq"
	"strings"
	"testing"
)

func TestBasicHTTPMethods(t *testing.T) {
	t.Parallel()

	tests := []struct {
		method   string
		function func(rq.Request, context.Context) (rq.Response, error)
	}{
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
			req := HTTPBin().JoinURL(lcmethod)
			// todo: separate test of url joins
			//require.Equal(t, "http://httpbin.org/"+lcmethod, req.URL.String())
			rep, err := test.function(req, context.TODO())
			require.NoError(t, err)
			require.Equal(t, 200, rep.Status())
		})
	}
}
