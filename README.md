<h1 align="center">github.com/tetratom/rq</h1>
<p align="center">
  <a href="https://godoc.org/github.com/tetratom/rq">
    <img src="https://godoc.org/github.com/tetratom/rq?status.svg" alt="GoDoc">
  </a>
  <a href="https://circleci.com/gh/tetratom/rq">
    <img src="https://img.shields.io/circleci/build/gh/tetratom/rq/master" alt="CircleCI">
  </a>
  <a href="https://codecov.io/gh/tetratom/rq">
    <img src="https://img.shields.io/codecov/c/github/tetratom/rq/master" alt="Codecov">
  </a>
</p>

_rq is a no-nonsense library for working with rest apis_

# highlights

- request constructors pass values: one `rq.Request` can be used as the basis of another without copy concerns
- easy access to common operations: `Path()`, `Queryf()`, `SetHeader()`, `Is2xx()`, and many more
- support for arbitrary request, response middleware, and marshallers
- context-first

# example

```go
import "github.com/tetratom/rq"

req := rq.
    Begin("https://httpbin.org").
    Path("get").
    Queryf("x", "y%d", 1)

rep, err := req.GET(context.Background())
switch {
case err != nil:
    panic(err)
case !rep.Is2xx():
    panic("bad response!")
}

log.Printf("response: %s", rep.ReadStringOrPanic())
```
