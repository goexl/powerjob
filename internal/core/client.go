package core

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/goexl/exception"
	"github.com/goexl/gox/field"
	builder2 "github.com/goexl/mengpo/internal/builder"
	"github.com/goexl/powerjob/internal/internal"
	"github.com/goexl/powerjob/internal/internal/builder"
	"github.com/goexl/powerjob/internal/internal/core"
	"github.com/goexl/powerjob/internal/internal/powerjob/request"
	"github.com/goexl/powerjob/internal/internal/powerjob/response"
	"github.com/goexl/powerjob/internal/param"
	"github.com/goexl/structer"
	"github.com/goexl/xiren"
)

type Client struct {
	params *param.Client
	logins map[string]int64
}

func NewClient(params *param.Client) *Client {
	return &Client{
		params: params,
		logins: make(map[string]int64),
	}
}

func (c *Client) Job() *builder.Job {
	return builder.NewJob(c.post)
}

func (c *Client) post(
	ctx context.Context,
	core *param.Core, path string,
	req core.Application, rsp core.SuccessChecker,
) (err error) {
	label := core.Label
	if id, ok := c.logins[label]; !ok || 0 == id {
		err = c.login(ctx, core)
	}
	if nil == err {
		req.Appid(c.logins[label])
		err = c.do(ctx, core, path, req, rsp, c.json(req))
	}

	return
}

func (c *Client) do(
	ctx context.Context,
	core *param.Core, path string,
	req core.Application, rsp core.SuccessChecker, body internal.Body,
) (err error) {
	http := c.params.Http.R().SetContext(ctx).SetResult(rsp).SetBody(req)
	label := core.Label
	if mse := builder2.New().Build().Set(req); nil != mse {
		err = mse
	} else if xse := xiren.New().Struct(req); nil != xse {
		err = xse
	} else if worker, ok := c.params.Workers[label]; !ok {
		err = exception.New().Message("未找到执行器").Field(field.New("label", label)).Build()
	} else if be := body(http, worker); nil != be {
		err = be
	} else if pr, pe := http.Post(worker.Url(path)); nil != pe {
		err = pe
	} else if pr.IsError() {
		err = exception.New().Message("服务器返回错误").Field(field.New("response", pr.String())).Build()
	} else if !rsp.Successful() {
		err = exception.New().Message("操作失败").Field(field.New("response", rsp)).Build()
	}

	return
}

func (c *Client) login(ctx context.Context, core *param.Core) (err error) {
	req := new(request.Login)
	rsp := new(response.Login)
	if de := c.do(ctx, core, "assert", req, rsp, c.setLogin(req, req)); nil != de {
		err = de
	} else {
		c.logins[core.Label] = rsp.Appid
	}

	return
}

func (c *Client) json(req core.Application) internal.Body {
	return func(request *resty.Request, _ *param.Worker) (err error) {
		request.SetBody(req)

		return
	}
}

func (c *Client) form(req core.Application) internal.Body {
	return func(request *resty.Request, _ *param.Worker) (err error) {
		form := make(map[string]string)
		err = structer.Converter().From(req).To(form).Build().Convert()
		if nil == err {
			request.SetFormData(form)
		}

		return
	}
}

func (c *Client) setLogin(req core.Application, login *request.Login) internal.Body {
	return func(request *resty.Request, worker *param.Worker) (err error) {
		login.Name = worker.App
		login.Password = worker.Password
		form := make(map[string]string)
		err = structer.Converter().From(req).To(&form).Build().Convert()
		if nil == err {
			request.SetFormData(form)
		}

		return
	}
}
