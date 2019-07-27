package rq_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/tetratom/rq"
	"net/http"
	"testing"
)

func TestRequestHeaders(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	req := HTTPBin().JoinURL("/anything").SetHeader("Foo", "Bar").SetHeader("Bax", "Baz")
	require.Equal(t, http.Header{"Foo": []string{"Bar"}, "Bax": []string{"Baz"}}, req.HeaderMap())
	rep, err := req.GET(ctx)
	require.NoError(t, err)
	require.Equal(t, 200, rep.Status())
	var result HTTPBinResponse
	require.NoError(t, rep.Unmarshal(&result))
	require.Equal(t, "Bar", result.Headers["Foo"])
	require.Equal(t, "Baz", result.Headers["Bax"])
}

func TestRequest_HeaderMap(t *testing.T) {
	t.Parallel()

	req := rq.Request{}
	require.Equal(t, http.Header{}, req.HeaderMap())

	req.Headers = []rq.Header{
		{"Single1", "a"},
		{"Multiple", "first"},
		{"Single2", "b"},
		{"Multiple", "second"},
	}
	require.Equal(t, http.Header{
		"Single1":  []string{"a"},
		"Single2":  []string{"b"},
		"Multiple": []string{"first", "second"},
	}, req.HeaderMap())
}

func TestRequest_AddHeader(t *testing.T) {
	t.Parallel()

	req := rq.Request{}
	require.Len(t, req.Headers, 0)

	req = req.AddHeader("Test1", "a")
	require.Equal(t, []rq.Header{{"Test1", "a"}}, req.Headers)

	req = req.AddHeader("Test1", "b")
	require.Equal(t, []rq.Header{{"Test1", "a"}, {"Test1", "b"}}, req.Headers)

	req = req.AddHeader("Test2", "%s", "c")
	require.Equal(t, []rq.Header{{"Test1", "a"}, {"Test1", "b"}, {"Test2", "c"}}, req.Headers)
}

func TestRequest_GetHeader(t *testing.T) {
	t.Parallel()

	req := rq.Begin().
		AddHeader("Test1", "a").
		AddHeader("Test2", "b").
		AddHeader("Test1", "c")

	require.Equal(t, "a; c", req.GetHeader("Test1"))
	require.Equal(t, "a; c", req.GetHeader("test1"))
	require.Equal(t, "a; c", req.GetHeader("TeSt1"))
	require.Equal(t, "b", req.GetHeader("TEST2"))
	require.Equal(t, "", req.GetHeader("Nothing"))
}

func TestRequest_SetHeader(t *testing.T) {
	t.Parallel()

	req := rq.Begin().
		AddHeader("Test1", "a").
		AddHeader("Test2", "b").
		AddHeader("TEST1", "%s", "c")
	require.Equal(t, []rq.Header{{"Test1", "a"}, {"Test2", "b"}, {"Test1", "c"}}, req.Headers)

	req = req.SetHeader("Test1", "d")
	require.Equal(t, []rq.Header{{"Test2", "b"}, {"Test1", "d"}}, req.Headers)
}

func TestRequest_RemoveHeader(t *testing.T) {
	t.Parallel()

	req := rq.Begin().
		AddHeader("Test1", "a").
		AddHeader("Test2", "b").
		AddHeader("Test1", "c")
	require.Equal(t, []rq.Header{{"Test1", "a"}, {"Test2", "b"}, {"Test1", "c"}}, req.Headers)

	req1 := req.RemoveHeader("Test1")
	require.Equal(t, []rq.Header{{"Test2", "b"}}, req1.Headers)

	req1 = req.RemoveHeader("TeSt1")
	require.Equal(t, []rq.Header{{"Test2", "b"}}, req1.Headers)

	req2 := req.RemoveHeader("Test2")
	require.Equal(t, []rq.Header{{"Test1", "a"}, {"Test1", "c"}}, req2.Headers)
}

func TestRequest_SetBasicAuth(t *testing.T) {
	t.Parallel()
	req := rq.Begin().SetBasicAuth("johndoe", "password123")
	require.Equal(t, "Basic am9obmRvZTpwYXNzd29yZDEyMw==", req.GetHeader("authorization"))
}

func TestRequest_SetBearerToken(t *testing.T) {
	t.Parallel()
	req := rq.Begin().SetBearerToken("mytoken")
	require.Equal(t, "Bearer mytoken", req.GetHeader("authorization"))
}
