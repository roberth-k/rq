package rq

func (resp *Response) getUnmarshallerOrDefault() Unmarshaller {
	if resp.Unmarshaller == nil {
		return UnmarshalJSON
	}

	return resp.Unmarshaller
}
