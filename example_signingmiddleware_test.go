package rq_test

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/tetratom/rq"
	"io"
	"io/ioutil"
)

// NewRequestSigningMiddleware returns a rq request middleware that generates
// an example digest based on a secret value and the request body (if any).
// The result is the X-Request-Signature header, a hex-encoded MD5 hash.
//
// If the request has a body, it will be fully buffered into memory.
func NewRequestSigningMiddleware(secret string) rq.RequestMiddleware {
	return func(req rq.Request) rq.Request {
		digest := md5.New()
		io.WriteString(digest, secret)

		if req.Body != nil {
			data, _ := ioutil.ReadAll(req.Body.Reader())
			digest.Write(data)
			req = req.SetBodyBytes(data)
		}

		return req.
			SetHeader(
				"X-Request-Signature",
				hex.EncodeToString(digest.Sum(nil)))
	}
}

// NewResponseVerifyingMiddleware returns a rq response middleware that converts
// a successful response into a failure if the response did not include the
// X-Response-Signature header.
func NewResponseVerifyingMiddleware() rq.ResponseMiddleware {
	return func(req rq.Request, rep rq.Response, err error) (rq.Response, error) {
		if err != nil {
			return rep, err
		}

		if !rep.HasHeader("X-Response-Signature") {
			return rep, errors.New("no response signature")
		}

		return rep, nil
	}
}

func Example_signingMiddleware() {
	api := rq.
		Begin("http://httpbin.org").
		Path("anything").
		AddRequestMiddlewares(NewRequestSigningMiddleware("much secure"))

	var response HTTPBinResponse
	rep, _ := api.SetBodyString("very covert").GET(context.TODO())
	rep.UnmarshalJSON(&response)
	fmt.Println("first request")
	fmt.Println("signature:", response.Headers["X-Request-Signature"])
	fmt.Println("body:", response.Data)

	_, err := api.
		AddResponseMiddlewares(NewResponseVerifyingMiddleware()).
		GET(context.TODO())
	fmt.Println("second request")
	fmt.Println("error:", err) // error: no response signature

	// Output:
	// first request
	// signature: 121aa41cde97e6d27d468e743a8c05a6
	// body: very covert
	// second request
	// error: no response signature
}
