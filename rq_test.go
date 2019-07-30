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
