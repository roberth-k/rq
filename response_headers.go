package rq

func (resp *Response) GetHeader(name string) string {
	return resp.Underlying.Header.Get(name)
}

func (resp *Response) HasHeader(name string) bool {
	return resp.Underlying.Header.Get(name) != ""
}

func (resp *Response) LookupHeader(name string) (string, bool) {
	value := resp.Underlying.Header.Get(name)
	return value, value != ""
}
