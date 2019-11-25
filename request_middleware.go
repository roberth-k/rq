package rq

func (req Request) AddRequestMiddlewares(middlewares ...RequestMiddleware) Request {
	if len(middlewares) == 0 {
		return req
	}

	out := make(
		[]RequestMiddleware,
		len(req.RequestMiddlewares),
		len(req.RequestMiddlewares)+len(middlewares))
	copy(out, req.RequestMiddlewares)
	out = append(out, middlewares...)
	req.RequestMiddlewares = out
	return req
}

func (req Request) SetRequestMiddlewares(middlewares ...RequestMiddleware) Request {
	req.RequestMiddlewares = middlewares
	return req
}

func (req Request) AddResponseMiddlewares(middlewares ...ResponseMiddleware) Request {
	if len(middlewares) == 0 {
		return req
	}

	out := make(
		[]ResponseMiddleware,
		len(req.ResponseMiddlewares),
		len(req.ResponseMiddlewares)+len(middlewares))
	copy(out, req.ResponseMiddlewares)
	out = append(out, middlewares...)
	req.ResponseMiddlewares = out
	return req
}

func (req Request) SetResponseMiddlewares(middlewares ...ResponseMiddleware) Request {
	req.ResponseMiddlewares = middlewares
	return req
}
