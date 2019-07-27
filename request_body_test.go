package rq_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestBasicMarshalling(t *testing.T) {
	t.Parallel()

	type Test struct {
		Foo int
	}

	ctx := context.Background()
	input := Test{Foo: 42}
	req := httpbin.JoinURL("/anything").WithBodyFromJSON(input)
	require.Equal(t, http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}, req.HeaderMap())
	rep, err := req.POST(ctx)
	require.NoError(t, err)
	require.Equal(t, 200, rep.Status)

	var result HTTPBinResponse
	err = rep.Unmarshal(&result)
	require.NoError(t, err)
	require.Equal(t, `{"Foo":42}`, result.Data)
}
