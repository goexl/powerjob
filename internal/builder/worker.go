package builder

import (
	"github.com/goexl/powerjob/internal/param"
)

type Worker struct {
	client *Client
	params *param.Worker
}

func NewWorker(endpoint string, app string, client *Client) *Worker {
	return &Worker{
		client: client,
		params: param.NewWorker(endpoint, app),
	}
}

func (w *Worker) Label(label string) (worker *Worker) {
	w.params.Label = label
	worker = w

	return
}

func (w *Worker) Password(password string) (worker *Worker) {
	w.params.Password = password
	worker = w

	return
}

func (w *Worker) Build() (client *Client) {
	w.client.params.Workers[w.params.Label] = w.params
	client = w.client

	return
}
