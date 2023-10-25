package status

type Err struct {
	Code    int
	Message string
}

func (e *Err) Error() string {
	return e.Message
}

func NewError(code int, message string) *Err {
	return &Err{
		Code:    code,
		Message: message,
	}
}
