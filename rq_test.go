package rq_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/tetratom/rq"
	"testing"
)

var httpbin = rq.Begin().SetURL("https://httpbin.org")

func TestSimpleGET(t *testing.T) {
	req := httpbin.JoinURL("/get")
	require.Equal(t, "GET", req.Method)
	require.Equal(t, "https://httpbin.org/get", req.URL.String())

	rep, err := req.GET(context.Background())
	require.NoError(t, err)
	require.Equal(t, 200, rep.Status)
}
