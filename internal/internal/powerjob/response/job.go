package response

import (
	"github.com/goexl/powerjob/internal/internal/powerjob"
)

type Job struct {
	powerjob.Response

	Id int64 `json:"data,omitempty"`
}
