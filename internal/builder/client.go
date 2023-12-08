package builder

import (
	"github.com/goexl/powerjob/internal/core"
	"github.com/goexl/powerjob/internal/param"
)

type Client struct {
	params *param.Client
}

func NewClient() *Client {
	return &Client{
		params: param.NewClient(),
	}
}

func (c *Client) Worker(endpoint string, app string) *Worker {
	return NewWorker(endpoint, app, c)
}

func (c *Client) Build() *core.Client {
	return core.NewClient(c.params)
}
