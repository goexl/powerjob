package builder

import (
	"github.com/goexl/powerjob/internal/param"
)

type Core struct {
	params *param.Core
}

func NewCore() *Core {
	return &Core{
		params: param.NewCore(),
	}
}

func (c *Core) Label(label string) (core *Core) {
	c.params.Label = label
	core = c

	return
}
