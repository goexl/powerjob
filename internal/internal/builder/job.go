package builder

import (
	"context"
	"encoding/json"
	"time"

	"github.com/goexl/gox"
	"github.com/goexl/powerjob/internal/callback"
	"github.com/goexl/powerjob/internal/internal/constant"
	"github.com/goexl/powerjob/internal/internal/powerjob/request"
	"github.com/goexl/powerjob/internal/internal/powerjob/response"
)

type Job struct {
	core   *Core
	params *request.Job
}

func NewJob(post callback.Post) *Job {
	return &Job{
		core:   NewCore(),
		params: request.NewJob(post),
	}
}

func (j *Job) Id(id int64) (job *Job) {
	j.params.Id = id
	job = j

	return
}

func (j *Job) Name(name string) (job *Job) {
	j.params.Name = name
	job = j

	return
}

func (j *Job) Description(description string) (job *Job) {
	j.params.Description = description
	job = j

	return
}

func (j *Job) Standalone() (job *Job) {
	j.params.ExpressType = constant.Standalone
	job = j

	return
}

func (j *Job) Random() (job *Job) {
	j.params.DispatchStrategy = constant.Random
	job = j

	return
}

func (j *Job) Health() (job *Job) {
	j.params.DispatchStrategy = "HEALTH_FIRST"
	job = j

	return
}

func (j *Job) Cron(cron string) (job *Job) {
	j.params.ExpressType = "CRON"
	j.params.Express = cron
	job = j

	return
}

func (j *Job) Rate(duration time.Duration) (job *Job) {
	j.params.ExpressType = "FIXED_RATE"
	j.params.Express = gox.ToString(duration.Microseconds())
	job = j

	return
}

func (j *Job) Delay(duration time.Duration) (job *Job) {
	j.params.ExpressType = constant.Delay
	j.params.Express = gox.ToString(duration.Microseconds())
	job = j

	return
}

func (j *Job) Http(url string) (http *Http) {
	j.params.ProcessorType = "BUILT_IN"
	j.params.ProcessorInfo = "tech.powerjob.official.processors.impl.HttpProcessor"
	http = NewHttp(j, j.params, url)

	return
}

func (j *Job) Instance() *Instance {
	return NewInstance(j, j.params)
}

func (j *Job) Do(ctx context.Context) (id int64, err error) {
	job := new(response.Job)
	if bytes, me := json.Marshal(j.params.Data); nil != me {
		err = me
	} else {
		j.params.Params = string(bytes)
		err = j.params.Post(ctx, j.core.params, "saveJob", j.params, job)
	}
	if nil == err {
		id = job.Id
	}

	return
}
