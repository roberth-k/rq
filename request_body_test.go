package rq_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/tetratom/rq"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestRequest_SetMarshaller(t *testing.T) {
	t.Parallel()

	req := rq.Request{}
	require.Nil(t, req.Marshaller)

	m := func(req rq.Request, _ interface{}) rq.Request {
		return req
	}

	req = req.SetMarshaller(m)
	require.NotNil(t, req.Marshaller)
}

func TestRequest_SetBody(t *testing.T) {
	t.Parallel()
	req := rq.Request{}
	require.Nil(t, req.Body)
	r := strings.NewReader("test")
	req = req.SetBodyReader(r)
	require.Equal(t, r, req.Body.Reader())
}

func TestRequest_SetBodyBytes(t *testing.T) {
	t.Parallel()
	req := rq.Request{}
	require.Nil(t, req.Body)
	expect := []byte("test")
	req = req.SetBodyBytes(expect)
	require.NotNil(t, req.Body)
	actual, err := ioutil.ReadAll(req.Body.Reader())
	require.NoError(t, err)
	require.Equal(t, expect, actual)
}

func TestRequest_SetBodyString(t *testing.T) {
	t.Parallel()
	req := rq.Request{}
	require.Nil(t, req.Body)
	expect := "test"
	req = req.SetBodyString(expect)
	require.NotNil(t, req.Body)
	actual, err := ioutil.ReadAll(req.Body.Reader())
	require.NoError(t, err)
	require.Equal(t, expect, string(actual))
}

func TestRequest_Marshal(t *testing.T) {
	t.Parallel()

	var calls int
	var value interface{}

	m := func(req rq.Request, v interface{}) rq.Request {
		calls++
		value = v

		return req.SetHeader("Test", "done")
	}

	req := rq.Request{}
	require.False(t, req.HasHeader("Test"))

	req = req.SetMarshaller(m)
	require.Zero(t, calls)

	actual := map[string]string{"x": "y"}

	req = req.Marshal(actual)
	require.Nil(t, req.Body)
	require.Equal(t, "done", req.GetHeader("Test"))
	require.Equal(t, 1, calls)
	require.Equal(t, actual, value)
}

func TestRequest_MarshalJSON(t *testing.T) {
	t.Parallel()

	actual := map[string]string{"x": "y"}
	req := rq.Request{}.MarshalJSON(actual)
	require.NotNil(t, req.Body)
	data, err := ioutil.ReadAll(req.Body.Reader())
	require.NoError(t, err)
	require.Equal(t, `{"x":"y"}`, string(data))
	require.Equal(t, "application/json; charset=utf-8", req.GetHeader("Content-Type"))
}

func TestRequest_MarshalForm(t *testing.T) {
	t.Parallel()

	actual := url.Values{"x": []string{"y", "z"}}
	req := rq.Request{}.MarshalForm(actual)
	require.NotNil(t, req.Body)
	data, err := ioutil.ReadAll(req.Body.Reader())
	require.NoError(t, err)
	require.Equal(t, "x=y&x=z", string(data))
	require.Equal(t, "application/x-www-form-urlencoded; charset=utf-8", req.GetHeader("Content-Type"))
}

func TestBasicMarshalling(t *testing.T) {
	t.Parallel()

	type Test struct {
		Foo int
	}

	ctx := context.Background()
	input := Test{Foo: 42}
	req := HTTPBin().Path("/anything").MarshalJSON(input)
	require.Equal(t, http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}, req.HeaderMap())
	rep, err := req.POST(ctx)
	require.NoError(t, err)
	require.Equal(t, 200, rep.Status())

	var result HTTPBinResponse
	err = rep.Unmarshal(&result)
	require.NoError(t, err)
	require.Equal(t, `{"Foo":42}`, result.Data)
}
