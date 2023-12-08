package param

import (
	"github.com/goexl/gox/http"
)

type Http struct {
	Method http.Method `json:"method,omitempty"`
	Url    string      `json:"url,omitempty"`
}
