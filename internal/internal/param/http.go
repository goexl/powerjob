package param

import (
	"github.com/goexl/gox/http"
)

type Http struct {
	Method  http.Method       `json:"method,omitempty"`
	Url     string            `json:"url,omitempty"`
	Timeout int               `json:"timeout,omitempty"`
	Mime    string            `json:"mediaType,omitempty"`
	Body    any               `json:"body,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
}

func NewHttp(url string) *Http {
	return &Http{
		Method:  http.MethodPost,
		Url:     url,
		Headers: make(map[string]string),
	}
}
