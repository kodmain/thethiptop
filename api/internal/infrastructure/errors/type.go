package errors

import "github.com/kodmain/thetiptop/api/internal/infrastructure/observability/logger"

type ErrorInterface interface {
	Error() string
	Code() int
	Log(error) *Error
}

type Error struct {
	code    int
	message string
	logged  bool
}

func (e *Error) Log(err error) *Error {
	if !e.logged {
		e.logged = logger.Error(err)
	}

	return e
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Error() string {
	return e.message
}

func New(code int, message string) *Error {
	err := &Error{
		code:    code,
		message: message,
	}

	registredErrors[message] = err

	return err
}

func ListErrors() map[string]*Error {
	return registredErrors
}