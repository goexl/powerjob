package builder

import (
	"context"

	"github.com/goexl/powerjob/internal/callback"
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

func (j *Job) Id(id *int64) (job *Job) {
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

func (j *Job) Http(description string) (job *Job) {
	j.params.Description = description
	job = j

	return
}

func (j *Job) Do(ctx context.Context) error {
	return j.params.Post(ctx, j.core.params, "saveJob", j.params, &response.Job{
		Id: j.params.Id,
	})
}
