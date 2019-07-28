package rq_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tetratom/rq"
	"net/url"
	"testing"
)

func TestRequest_AddQuery(t *testing.T) {
	t.Parallel()
	req := rq.Request{}
	require.Equal(t, "", req.URL.RawQuery)
	req = req.AddQuery("test1", "x%d", 1)
	require.Equal(t, "test1=x1", req.URL.RawQuery)
	req = req.AddQuery("test2", "b")
	require.Equal(t, "test1=x1&test2=b", req.URL.RawQuery)
}

func TestRequest_GetQuery(t *testing.T) {
	t.Parallel()
	req := rq.Request{}
	req.URL.RawQuery = "test1=a&test2=b"
	require.Equal(t, "a", req.GetQuery("test1"))
	require.Equal(t, "b", req.GetQuery("test2"))
}

func TestRequest_RemoveQuery(t *testing.T) {
	t.Parallel()
	req := rq.Request{}
	req.URL.RawQuery = "test1=a&test2=b&test1=c"
	require.Equal(t, "test2=b", req.RemoveQuery("test1").URL.RawQuery)
	require.Equal(t, "test1=a&test1=c", req.RemoveQuery("test2").URL.RawQuery)
}

func TestRequest_SetQuery(t *testing.T) {
	t.Parallel()
	req := rq.Request{}
	req.URL.RawQuery = "test1=a&test2=b&test1=c"
	require.Equal(t, "test1=d&test2=b", req.SetQuery("test1", "%s", "d").URL.RawQuery)
	require.Equal(t, "test1=a&test1=c&test2=d", req.SetQuery("test2", "d").URL.RawQuery)
}

func TestRequest_ReplaceQuery(t *testing.T) {
	t.Parallel()
	req := rq.Request{}
	req.URL.RawQuery = "test1=a&test2=b&test1=c"
	require.Equal(t, "test3=x", req.ReplaceQuery(url.Values{"test3": []string{"x"}}).URL.RawQuery)
}

func TestRequest_GetFragment(t *testing.T) {
	t.Parallel()
	req := rq.Request{}
	require.Equal(t, "", req.URL.Fragment)
	require.Equal(t, "", req.GetFragment())
	req.URL.Fragment = "test"
	require.Equal(t, "test", req.URL.Fragment)
	require.Equal(t, "test", req.GetFragment())
}

func TestRequest_SetFragment(t *testing.T) {
	t.Parallel()
	req := rq.Request{}
	require.Equal(t, "test", req.SetFragment("test").URL.Fragment)
}
