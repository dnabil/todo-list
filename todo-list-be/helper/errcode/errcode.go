package errcode

type ErrCode struct {
	message string
	code    int
}

type ErrCodeI interface {
	error
	Code() int
}

func New(msg string, code int) *ErrCode {
	return &ErrCode{
		message: msg,
		code:    code,
	}
}

func (e *ErrCode) Error() string {
	return e.message
}

func (e *ErrCode) Code() int {
	return e.code
}