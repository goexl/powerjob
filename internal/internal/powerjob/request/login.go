package request

type Login struct {
	Name     string `json:"appName,omitempty"`
	Password string `json:"password,omitempty"`
}

func (l *Login) Appid(_ int64) {}
