package param

import (
	"github.com/goexl/powerjob/internal/internal/constant"
)

type Core struct {
	Label string
}

func NewCore() *Core {
	return &Core{
		Label: constant.LabelSystem,
	}
}
