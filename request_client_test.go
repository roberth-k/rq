package rq_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tetratom/rq"
	"net/http"
	"testing"
)

func TestRequest_SetClient(t *testing.T) {
	t.Parallel()
	var client http.Client
	req := rq.Begin("")
	require.Equal(t, req.Client, http.DefaultClient)
	req = req.SetClient(&client)
	require.Equal(t, req.Client, &client)
}
