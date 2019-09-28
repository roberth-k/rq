package rq

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

func UnmarshalJSON(response Response, value interface{}) error {
	if !response.IgnoreContentType &&
		!strings.HasPrefix(response.GetHeader("Content-Type"), "application/json") {

		return errors.New("expected content-type: application/json")
	}

	data, err := response.ReadAll()
	if err != nil {
		return err
	}

	return json.Unmarshal(data, value)
}

func UnmarshalXML(response Response, value interface{}) error {
	if !response.IgnoreContentType &&
		!strings.HasPrefix(response.GetHeader("Content-Type"), "text/xml") &&
		!strings.HasPrefix(response.GetHeader("Content-Type"), "application/xml") {

		return errors.New("expected content-type: text/xml or application/xml")
	}

	data, err := response.ReadAll()
	if err != nil {
		return err
	}

	return xml.Unmarshal(data, value)
}
