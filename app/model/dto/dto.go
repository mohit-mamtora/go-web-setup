package dto

import "encoding/json"

type Request map[string]interface{}

func (r *Request) ToString() (string, error) {
	str, error := json.Marshal(r)
	return string(str), error
}

type Response map[string]interface{}

func (r *Response) ToString() (string, error) {
	str, error := json.Marshal(r)
	return string(str), error
}
