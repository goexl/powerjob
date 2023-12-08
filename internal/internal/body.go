package internal

import (
	"github.com/go-resty/resty/v2"
	"github.com/goexl/powerjob/internal/param"
)

type Body func(request *resty.Request, worker *param.Worker) error
