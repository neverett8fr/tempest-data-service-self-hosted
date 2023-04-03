package application

import "encoding/json"

type Response struct {
	Errors []error     `json:"errors"`
	Data   interface{} `json:"data"`
}

func NewResponse(data interface{}, err ...error) *Response {

	return &Response{
		Errors: err,
		Data:   data,
	}
}

type FileMetadata struct {
	Extension string `json:"extension"`
	Name      string `json:"name"`
	Size      int    `json:"size"`
}

type File struct {
	Key      string          `json:"key"`
	User     string          `json:"username"`
	Metadata FileMetadata    `json:"metadata"`
	Data     json.RawMessage `json:"data"`
}
