package rq_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tetratom/rq"
	"testing"
)

func TestBegin(t *testing.T) {
	t.Parallel()
	req := rq.Begin("http://example.com")
	require.Equal(t, "http://example.com", req.URL.String())
}

func TestSetURL(t *testing.T) {
	t.Parallel()
	req := rq.SetURL("http://example.com")
	require.Equal(t, "http://example.com", req.URL.String())
}

func TestMixedQueryAndURL(t *testing.T) {
	t.Parallel()

	r := rq.Begin("").
		SetURL("https://example.com:123").
		SetQueryf("u", "1").
		Path("/foo/bar").
		SetQueryf("t", "3")
	expect := "https://example.com:123/foo/bar?t=3&u=1"
	require.Equal(t, expect, r.URL.String())
}
