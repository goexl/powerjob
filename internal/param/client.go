package param

import (
	"github.com/goexl/http"
)

type Client struct {
	Http    *http.Client
	Workers map[string]*Worker
}

func NewClient() *Client {
	return &Client{
		Http:    http.New().Build(),
		Workers: make(map[string]*Worker, 1),
	}
}
