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

# Example

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
