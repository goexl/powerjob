package request

import (
	"github.com/goexl/powerjob/internal/callback"
	"github.com/goexl/powerjob/internal/internal/powerjob"
)

type Job struct {
	powerjob.Request

	Id            *int64 `json:"jobId,string,omitempty"`
	Name          string `json:"jobName,omitempty"`
	Description   string `json:"jobDescription,omitempty"`
	Params        string `json:"jobParams,omitempty"`
	ExpressType   uint8  `json:"timeExpressionType,omitempty"`
	Express       string `json:"timeExpression,omitempty"`
	ProcessorType uint8  `json:"processorType,omitempty"`
	ProcessorInfo string `json:"processorInfo,omitempty"`

	Post callback.Post `json:"-"`
}

func NewJob(post callback.Post) *Job {
	return &Job{
		Post: post,
	}
}
