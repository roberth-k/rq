package rq_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tetratom/rq"
	"net/http"
	"testing"
)

func TestResponse_Map(t *testing.T) {
	t.Parallel()
	rep := rq.Response{}
	require.Nil(t, rep.Underlying)
	var httpResponse http.Response
	rep = rep.Map(func(rep rq.Response) rq.Response {
		rep.Underlying = &httpResponse
		return rep
	})
	require.Equal(t, &httpResponse, rep.Underlying)
}
