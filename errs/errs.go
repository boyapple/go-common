package errs

import (
	"fmt"
)

const (
	RetOk      = 0
	RetUnknown = 999

	Success = "success"
)

type Error struct {
	Code int
	Msg  string
}

func (e *Error) Error() string {
	if e == nil {
		return Success
	}
	return fmt.Sprintf("code:%d msg:%s", e.Code, e.Msg)
}

func New(code int, msg string) error {
	err := &Error{
		Code: code,
		Msg:  msg,
	}
	return err
}

func Newf(code int, format string, args ...interface{}) error {
	err := &Error{
		Code: code,
		Msg:  fmt.Sprintf(format, args...),
	}
	return err
}

func Code(e error) int {
	if e == nil {
		return RetOk
	}
	err, ok := e.(*Error)
	if !ok {
		return RetUnknown
	}
	if err == nil {
		return RetOk
	}
	return err.Code
}

func Msg(e error) string {
	if e == nil {
		return Success
	}
	err, ok := e.(*Error)
	if !ok {
		return e.Error()
	}
	if err == (*Error)(nil) {
		return Success
	}
	return err.Msg
}
