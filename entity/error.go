package entity

type ErrorResponse struct {
	HttpCode int
	Err      *ErrorPayload `json:"error"`
}

type ErrorPayload struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Param   any    `json:"param"`
	Code    string `json:"code"`
}

func (e *ErrorResponse) Error() string {
	return e.Err.Message
}
