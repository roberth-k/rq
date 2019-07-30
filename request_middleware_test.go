package rq_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tetratom/rq"
	"testing"
)

func noopRequestMiddleware(req rq.Request) rq.Request {
	return req
}

func noopResponseMiddleware(_ rq.Request, rep rq.Response, err error) (rq.Response, error) {
	return rep, err
}

func TestRequest_AddRequestMiddlewares(t *testing.T) {
	t.Parallel()
	req := rq.Request{}
	require.Empty(t, req.RequestMiddlewares)
	req = req.AddRequestMiddlewares(noopRequestMiddleware)
	require.Len(t, req.RequestMiddlewares, 1)
}

func TestRequest_SetRequestMiddlewares(t *testing.T) {
	t.Parallel()
	req := rq.Request{}
	req = req.AddRequestMiddlewares(noopRequestMiddleware)
	require.Len(t, req.RequestMiddlewares, 1)
	req = req.SetRequestMiddlewares(noopRequestMiddleware, noopRequestMiddleware)
	require.Len(t, req.RequestMiddlewares, 2)
}

func TestRequest_SetResponseMiddlewares(t *testing.T) {
	t.Parallel()
	req := rq.Request{}
	require.Empty(t, req.ResponseMiddlewares)
	req = req.AddResponseMiddlewares(noopResponseMiddleware)
	require.Len(t, req.ResponseMiddlewares, 1)
}

func TestRequest_AddResponseMiddlewares(t *testing.T) {
	t.Parallel()
	req := rq.Request{}
	req = req.AddResponseMiddlewares(noopResponseMiddleware)
	require.Len(t, req.ResponseMiddlewares, 1)
	req = req.SetResponseMiddlewares(noopResponseMiddleware, noopResponseMiddleware)
	require.Len(t, req.ResponseMiddlewares, 2)
}
