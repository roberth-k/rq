package rq

import (
	"bytes"
	"encoding/json"
	"io"
)

type Marshaller func(Request, interface{}) (io.Reader, error)

func NOOPMarshaller(_ Request, v interface{}) (io.Reader, error) {
	return nil, nil
}

func JSONMarshaller(_ Request, v interface{}) (io.Reader, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(data), nil
}
