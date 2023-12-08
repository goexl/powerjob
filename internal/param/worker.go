package param

import (
	"github.com/goexl/gox"
	"github.com/goexl/powerjob/internal/internal/constant"
)

type Worker struct {
	Label    string `json:"label,omitempty" validate:"required"`
	Endpoint string `json:"endpoint,omitempty" validate:"required,url"`
	App      string `json:"app,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}

func NewWorker(endpoint string, app string) *Worker {
	return &Worker{
		Label:    constant.LabelSystem,
		Endpoint: endpoint,
		App:      app,
	}
}

func (w *Worker) Url(path string) string {
	return gox.StringBuilder(w.Endpoint, constant.Slash, "openApi", constant.Slash, path).String()
}
