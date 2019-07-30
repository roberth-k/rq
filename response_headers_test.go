package rq_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tetratom/rq"
	"net/http"
	"testing"
)

func TestResponseHeaders(t *testing.T) {
	t.Parallel()
	rep := rq.Response{}
	rep.Underlying = &http.Response{}
	rep.Underlying.Header = make(http.Header)
	require.Equal(t, "", rep.GetHeader("X-Test"))
	require.False(t, rep.HasHeader("X-Test"))
	s, b := rep.LookupHeader("X-Test")
	require.Equal(t, "", s)
	require.False(t, b)

	rep.Underlying.Header.Add("X-Test", "test")
	require.Equal(t, "test", rep.GetHeader("X-Test"))
	require.True(t, rep.HasHeader("X-Test"))
	s, b = rep.LookupHeader("X-Test")
	require.Equal(t, "test", s)
	require.True(t, b)
}
