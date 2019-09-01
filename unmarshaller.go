package rq

import (
	"encoding/json"
	"encoding/xml"
	"github.com/pkg/errors"
	"strings"
)

func UnmarshalJSON(response Response, value interface{}) error {
	if !strings.HasPrefix(response.GetHeader("Content-Type"), "application/json") {
		return errors.New("expected content-type: application/json")
	}

	data, err := response.ReadAll()
	if err != nil {
		return err
	}

	return json.Unmarshal(data, value)
}

func UnmarshalXML(response Response, value interface{}) error {
	if !strings.HasPrefix(response.GetHeader("Content-Type"), "text/xml") {
		return errors.New("expected content-type: text/xml")
	}

	data, err := response.ReadAll()
	if err != nil {
		return err
	}

	return xml.Unmarshal(data, value)
}
