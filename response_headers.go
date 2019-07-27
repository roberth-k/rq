package rq

func (resp *Response) GetHeader(name string) string {
	return resp.response.Header.Get(name)
}

func (resp *Response) HasHeader(name string) bool {
	return resp.response.Header.Get(name) != ""
}

func (resp *Response) LookupHeader(name string) (string, bool) {
	value := resp.response.Header.Get(name)
	return value, value != ""
}
