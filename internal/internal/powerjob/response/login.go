package response

import (
	"github.com/goexl/powerjob/internal/internal/powerjob"
)

type Login struct {
	powerjob.Response

	Appid int64 `json:"data,omitempty"`
}
