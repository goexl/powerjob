package powerjob

type Response struct {
	Success bool   `json:"success,omitempty"`
	Message string `json:"message,omitempty"`
}

func (r *Response) Successful() bool {
	return r.Success
}
