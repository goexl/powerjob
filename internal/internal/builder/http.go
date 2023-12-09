package builder

import (
	"time"

	"github.com/goexl/gox/http"
	"github.com/goexl/powerjob/internal/internal/param"
	"github.com/goexl/powerjob/internal/internal/powerjob/request"
)

type Http struct {
	job    *Job
	req    *request.Job
	params *param.Http
}

func NewHttp(job *Job, req *request.Job, url string) *Http {
	return &Http{
		job:    job,
		req:    req,
		params: param.NewHttp(url),
	}
}

func (h *Http) Get() *Http {
	return h.method(http.MethodGet)
}

func (h *Http) Post() *Http {
	return h.method(http.MethodPost)
}

func (h *Http) Put() *Http {
	return h.method(http.MethodPut)
}

func (h *Http) Delete() *Http {
	return h.method(http.MethodDelete)
}

func (h *Http) Timeout(timeout time.Duration) (http *Http) {
	h.params.Timeout = int(timeout.Seconds())
	http = h

	return
}

func (h *Http) Mime(mime string) (http *Http) {
	h.params.Mime = mime
	http = h

	return
}

func (h *Http) Header(key string, value string) (http *Http) {
	h.params.Headers[key] = value
	http = h

	return
}

func (h *Http) Body(body any) (http *Http) {
	h.params.Body = body
	http = h

	return
}

func (h *Http) Build() (job *Job) {
	h.req.Data = h.params
	job = h.job

	return
}

func (h *Http) method(method http.Method) (http *Http) {
	h.params.Method = method
	http = h

	return
}
