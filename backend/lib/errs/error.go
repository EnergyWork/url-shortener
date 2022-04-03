package errs

import (
	"fmt"
)

type Error struct {
	Msg  string `json:"message"`
	Code int    `json:"code"`
}

func NewError() *Error {
	return &Error{}
}

func (obj *Error) SetCode(code int) *Error {
	obj.Code = code
	return obj
}

func (obj *Error) SetMsg(msg string) *Error {
	obj.Msg = msg
	return obj
}

func (obj *Error) String() string {
	return fmt.Sprintf("error: %d, %s", obj.Code, obj.Msg)
}

func (obj *Error) Error() error {
	return fmt.Errorf("error: %d, %s", obj.Code, obj.Msg)
}
