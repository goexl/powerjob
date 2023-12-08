package powerjob

type Request struct {
	App int64 `json:"appId,omitempty"`
}

func (r *Request) Appid(appid int64) {
	r.App = appid
}
