package rest

import (
	"net/http"
)

type ResponseWrapper struct {
	Status   string   `json:"status"`
	Code     int      `json:"code"`
	Messages []string `json:"messages"`
}

func NewResponseWrapper(code int, err error) ResponseWrapper {
	r := ResponseWrapper{}
	r.Status = http.StatusText(code)
	r.Code = code
	if err != nil {
		r.Messages = []string{err.Error()}
	}
	return r
}
