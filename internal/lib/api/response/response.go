package response

type Respone struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func Error(msg string) Respone {
	return Respone{
		Status: StatusError,
		Error:  msg,
	}
}

func OK() Respone {
	return Respone{
		Status: StatusOK,
	}
}
