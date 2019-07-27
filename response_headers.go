package rq

func (resp *Response) GetHeader(name string) string {
	return resp.Headers.Get(name)
}

func (resp *Response) HasHeader(name string) bool {
	return resp.Headers.Get(name) != ""
}

func (resp *Response) LookupHeader(name string) (string, bool) {
	value := resp.Headers.Get(name)
	return value, value != ""
}
