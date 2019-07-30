package rq_test

import (
	"context"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/tetratom/rq"
	"testing"
)

func TestRequest_Map(t *testing.T) {
	t.Parallel()
	req := rq.Request{}
	require.Equal(t, "", req.Method)
	req = req.Map(func(req rq.Request) rq.Request {
		req.Method = "test"
		return req
	})
	require.Equal(t, "test", req.Method)
}

func TestRequest_SetUnmarshaller(t *testing.T) {
	t.Parallel()
	req := rq.Request{}
	require.Nil(t, req.Unmarshaller)
	req = req.SetUnmarshaller(rq.UnmarshalJSON)
	require.NotNil(t, req.Unmarshaller)
}

func TestRequest_SetAndGetAndHasError(t *testing.T) {
	t.Parallel()
	req := rq.Request{}
	require.Nil(t, req.GetError())
	require.False(t, req.HasError())
	err := errors.New("test")
	req = req.SetError(err)
	require.Equal(t, err, req.GetError())
	require.True(t, req.HasError())
}

func TestRequest_GetContext(t *testing.T) {
	t.Parallel()
	req := rq.Request{}
	require.NotNil(t, req.GetContext())
	var ctx context.Context
	middleware := func(req rq.Request) rq.Request {
		ctx = req.GetContext()
		return req.SetError(errors.New("oops"))
	}
	req = req.AddRequestMiddlewares(middleware)
	_, err := req.Do(context.TODO())
	require.Error(t, err)
	require.NotNil(t, ctx)
}

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
