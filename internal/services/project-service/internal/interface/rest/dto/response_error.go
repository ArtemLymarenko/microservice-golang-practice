package dto

type ResponseErr struct {
	Message string `json:"message"`
}

func NewResponseErr(err error) *ResponseErr {
	return &ResponseErr{
		err.Error(),
	}
}
