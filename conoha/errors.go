package conoha

type errorResponseBody struct {
	Error        *errorMessage        `json:"error"`
	Unauthorized *unauthorizedMessage `json:"unauthorized"`
}

type errorMessage struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type unauthorizedMessage struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// ErrUnauthorized will used when unauthorized response returned.
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
