package rq_test

import (
	"context"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/tetratom/rq"
	"testing"
)

func TestRequest_Do(t *testing.T) {
	t.Parallel()
	//emptyRequest := rq.Request{}
	emptyResponse := rq.Response{}
	testerr := errors.New("test")

	t.Run("return immediately when error is non-nil", func(t *testing.T) {
		req := rq.Request{}.SetError(testerr)
		rep, err := req.Do(context.TODO())
		require.Equal(t, testerr, err)
		require.Equal(t, emptyResponse, rep)
	})

	t.Run("return when a request middleware errors", func(t *testing.T) {
		counter := 0
		middleware := func(req rq.Request) rq.Request {
			if counter == 1 {
				counter++
				return req.SetError(testerr)
			}

			counter++
			return req
		}

		req := rq.Request{}.AddRequestMiddlewares(middleware, middleware)
		rep, err := req.Do(context.TODO())
		require.Equal(t, testerr, err)
		require.Equal(t, emptyResponse, rep)
		require.Equal(t, 2, counter)
	})

	t.Run("return when instantiating the inner request fails", func(t *testing.T) {
		req := rq.Request{}.SetMethod("\n")
		rep, err := req.Do(context.TODO())
		require.Error(t, err)
		require.Equal(t, emptyResponse, rep)
	})
}
