package errors_handle

import (
	"github.com/avp365/arch-pat/internal/com/mov"
)

var c = make(map[error]func())

func init() {
	c[mov.ErrVariablePositionNotFound] = mov.ErrVariablePositionNotFoundHandler
	c[mov.ErrVariableXPositionNotFound] = mov.ErrVariableXPositionNotFoundHandler
}

type ErrorHandler struct {
	c map[error]func()
	e error
	q string
}

func NewErrorHandler(err error, c map[error]func()) ErrorHandler {

	return ErrorHandler{e: err, c: c}
}

func (eh *ErrorHandler) Handle() {

	handler, ok := eh.c[eh.e]

	if !ok {
		panic("not error handler")
	}

	handler()

}
