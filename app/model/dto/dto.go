package dto

import "encoding/json"

type Response map[string]interface{}

func (r *Response) Get(key string) (interface{}, bool) {
	value, ok := (*r)[key]
	return value, ok
}

func (r *Response) ToString() (string, error) {
	str, error := json.Marshal(r)
	return string(str), error
}
