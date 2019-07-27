package rq_test

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/tetratom/rq"
	"testing"
)

func TestResponseStatusMethods(t *testing.T) {
	t.Parallel()

	tests := []int{
		200, 299,
		300, 399,
		400, 499,
		500, 599,
	}

	all := []func(rq.Response) bool{
		rq.Response.Is1xx,
		rq.Response.Is2xx,
		rq.Response.Is3xx,
		rq.Response.Is4xx,
		rq.Response.Is5xx,
	}

	for _, code := range tests {
		code := code
		t.Run(fmt.Sprintf("%d", code), func(t *testing.T) {
			t.Parallel()
			req := HTTPBin().JoinFormatURL("status/%d", code)
			rep, err := req.GET(context.TODO())
			require.NoError(t, err)

			for i, f := range all {
				min, max := (i+1)*100, (i+2)*100

				if code >= min && code < max {
					require.True(t, f(rep))
				} else {
					require.False(t, f(rep))
				}
			}
		})
	}
}
