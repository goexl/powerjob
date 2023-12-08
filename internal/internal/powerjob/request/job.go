package request

import (
	"github.com/goexl/powerjob/internal/callback"
	"github.com/goexl/powerjob/internal/internal/constant"
	"github.com/goexl/powerjob/internal/internal/powerjob"
)

type Job struct {
	powerjob.Request

	Id            int64  `json:"jobId,string,omitempty"`
	Name          string `json:"jobName,omitempty"`
	Description   string `json:"jobDescription,omitempty"`
	Params        string `json:"jobParams,omitempty"`
	ExecuteType   string `json:"executeType,omitempty" validate:"required"`
	ExpressType   string `json:"timeExpressionType,omitempty" validate:"required"`
	Express       string `json:"timeExpression,omitempty" validate:"required"`
	ProcessorType string `json:"processorType,omitempty" validate:"required"`
	ProcessorInfo string `json:"processorInfo,omitempty" validate:"required"`

	Post callback.Post `json:"-"`
	Data any           `json:"-"`
}

func NewJob(post callback.Post) *Job {
	return &Job{
		Name:        constant.Unknown,
		ExecuteType: constant.Standalone,
		ExpressType: constant.Delay,
		Express:     "1",
		Description: constant.Unknown,
		Post:        post,
	}
}
