package rq

func (resp Response) Status() int {
	return resp.Underlying.StatusCode
}

func (resp Response) Is1xx() bool {
	status := resp.Status()
	return status >= 100 && status < 200
}

func (resp Response) Is2xx() bool {
	status := resp.Status()
	return status >= 200 && status < 300
}

func (resp Response) Is3xx() bool {
	status := resp.Status()
	return status >= 300 && status < 400
}

func (resp Response) Is4xx() bool {
	status := resp.Status()
	return status >= 400 && status < 500
}

func (resp Response) Is5xx() bool {
	status := resp.Status()
	return status >= 500 && status < 600
}
