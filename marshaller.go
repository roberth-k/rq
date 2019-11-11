package rq

import (
	"encoding/json"
	"encoding/xml"
	"net/url"
)

func NOOPMarshaller(req Request, _ interface{}) Request {
	return req
}

func JSONMarshaller(req Request, value interface{}) Request {
	data, err := json.Marshal(value)
	if err != nil {
		return req.SetError(err)
	}

	return req.
		SetBodyBytes(data).
		SetHeader("Content-Type", "application/json; charset=utf-8")
}

func XMLMarshaller(req Request, value interface{}) Request {
	data, err := xml.Marshal(value)
	if err != nil {
		return req.SetError(err)
	}

	return req.
		SetBodyBytes(data).
		SetHeader("Content-Type", "text/xml; charset=utf-8")
}

func FormMarshaller(req Request, value interface{}) Request {
	data := value.(url.Values)
	return req.
		SetBodyString(data.Encode()).
		SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
}
