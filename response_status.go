package rq

func (resp *Response) Is1xx() bool {
	return resp.Status >= 100 && resp.Status < 200
}

func (resp *Response) Is2xx() bool {
	return resp.Status >= 200 && resp.Status < 300
}

func (resp *Response) Is3xx() bool {
	return resp.Status >= 300 && resp.Status < 400
}

func (resp *Response) Is4xx() bool {
	return resp.Status >= 400 && resp.Status < 500
}

func (resp *Response) Is5xx() bool {
	return resp.Status >= 500 && resp.Status < 600
}
