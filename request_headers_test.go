package rq_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestRequestHeaders(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	req := httpbin.JoinURL("/anything").SetHeader("Foo", "Bar").SetHeader("Bax", "Baz")
	require.Equal(t, http.Header{"Foo": []string{"Bar"}, "Bax": []string{"Baz"}}, req.HeaderMap())
	rep, err := req.GET(ctx)
	require.NoError(t, err)
	require.Equal(t, 200, rep.Status)
	var result HTTPBinResponse
	require.NoError(t, rep.Unmarshal(&result))
	require.Equal(t, "Bar", result.Headers["Foo"])
	require.Equal(t, "Baz", result.Headers["Bax"])
}
