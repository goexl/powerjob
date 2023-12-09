package param

import (
	"time"
)

type Instance struct {
	Max   int
	Retry int
	Limit time.Duration
}

func NewInstance() *Instance {
	return &Instance{}
}
