package powerjob

import (
	"github.com/goexl/powerjob/internal/builder"
	"github.com/goexl/powerjob/internal/core"
)

type Client = core.Client

func New() *builder.Client {
	return builder.NewClient()
}
