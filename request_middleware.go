package rq

func (req Request) AddRequestMiddlewares(middlewares ...RequestMiddleware) Request {
	req.RequestMiddlewares = append(req.RequestMiddlewares, middlewares...)
	return req
}

func (req Request) SetRequestMiddlewares(middlewares ...RequestMiddleware) Request {
	req.RequestMiddlewares = middlewares
	return req
}

func (req Request) AddResponseMiddlewares(middlewares ...ResponseMiddleware) Request {
	req.ResponseMiddlewares = append(req.ResponseMiddlewares, middlewares...)
	return req
}

func (req Request) SetResponseMiddlewares(middlewares ...ResponseMiddleware) Request {
	req.ResponseMiddlewares = middlewares
	return req
}
