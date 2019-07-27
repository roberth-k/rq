package rq

import (
	"encoding/json"
	"github.com/pkg/errors"
	"strings"
)

type Unmarshaller func(response *Response, value interface{}) error

func UnmarshalJSON(response *Response, value interface{}) error {
	if !strings.HasPrefix(response.GetHeader("Content-Type"), "application/json") {
		return errors.New("expected content-type: application/json")
	}

	data, err := response.ReadAll()
	if err != nil {
		return err
	}

	return json.Unmarshal(data, value)
}
