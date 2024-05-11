package errors_handle

import (
	"github.com/avp365/arch-pat/internal/command"
	"github.com/avp365/arch-pat/internal/command/mov"
)

var c = make(map[error]func(q chan command.Command))

func init() {
	c[mov.ErrVariablePositionNotFound] = mov.ErrVariablePositionNotFoundHandler
	c[mov.ErrVariableXPositionNotFound] = mov.ErrVariableXPositionNotFoundHandler
}

type ErrorHandler struct {
	c map[error]func(q chan command.Command)
	e error
	q chan command.Command
}

func NewErrorHandler(q chan command.Command, err error, c map[error]func(chan command.Command)) ErrorHandler {

	return ErrorHandler{q: q, e: err, c: c}
}

func (eh *ErrorHandler) Handle() {

	handler, ok := eh.c[eh.e]

	if !ok {
		panic("not error handler")
	}

	handler(eh.q)

}
