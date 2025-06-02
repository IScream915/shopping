package errs

import "fmt"

type ErrorCode struct {
	Code int
	Msg  string
}

func New(code int, msg string) *ErrorCode {
	return &ErrorCode{
		Code: code,
		Msg:  msg,
	}
}

// 实现了ERROR方法就满足error参数
func (e *ErrorCode) Error() string {
	return fmt.Sprintf("[ERROR] code: %d, msg: %s", e.Code, e.Msg)
}
