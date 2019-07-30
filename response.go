package rq

func (resp Response) Map(f func(Response) Response) Response {
	return f(resp)
}
