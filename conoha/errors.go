package conoha

type ErrorResponseBody struct {
	Error        *ErrorMessage        `json:"error"`
	Unauthorized *UnauthorizedMessage `json:"unauthorized"`
}

type ErrorMessage struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type UnauthorizedMessage struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type ErrUnauthorized struct {
	message string
}

func NewErrorUnauthorized(message string) *ErrUnauthorized {
	return &ErrUnauthorized{
		message: message,
	}
}

func (e *ErrUnauthorized) Error() string {
	return e.message
}
