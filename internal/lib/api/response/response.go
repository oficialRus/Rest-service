package response

type Response struct {
	Status string
	Error  string
}

const (
	StatusOk    = "OK"
	StatusError = "Error"
)
