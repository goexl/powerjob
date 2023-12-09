package builder

import (
	"time"

	"github.com/goexl/powerjob/internal/internal/param"
	"github.com/goexl/powerjob/internal/internal/powerjob/request"
)

type Instance struct {
	job    *Job
	req    *request.Job
	params *param.Instance
}

func NewInstance(job *Job, req *request.Job) *Instance {
	return &Instance{
		job:    job,
		req:    req,
		params: param.NewInstance(),
	}
}

func (i *Instance) Max(max int) (instance *Instance) {
	i.params.Max = max
	instance = i

	return
}

func (i *Instance) Retry(retry int) (instance *Instance) {
	i.params.Retry = retry
	instance = i

	return
}

func (i *Instance) Limit(limit time.Duration) (instance *Instance) {
	i.params.Limit = limit
	instance = i

	return
}

func (i *Instance) Build() (job *Job) {
	i.req.InstanceTimeLimit = int(i.params.Limit.Seconds())
	i.req.InstanceRetryNum = i.params.Retry
	i.req.MaxInstanceNum = i.params.Max
	job = i.job

	return
}
