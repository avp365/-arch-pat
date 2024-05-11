package errors_handle

import (
	"github.com/avp365/arch-pat/internal/command"
	"github.com/avp365/arch-pat/internal/command/mov"
)

var store = make(map[error]func(q chan command.Command))

func init() {
	store[mov.ErrVariablePositionNotFound] = mov.ErrVariablePositionNotFoundHandler
	store[mov.ErrVariableXPositionNotFound] = mov.ErrVariableXPositionNotFoundHandler
}

type ErrorHandler struct {
	cmd   command.Command
	store map[error]func(q chan command.Command)
	e     error
	queue chan command.Command
}

func NewErrorHandler(cmd command.Command, err error, store map[error]func(chan command.Command), q chan command.Command) ErrorHandler {

	return ErrorHandler{cmd: cmd, e: err, store: store, queue: q}
}

func (eh *ErrorHandler) Handle() {

	handler, ok := eh.store[eh.e]

	if !ok {
		panic("not error handler")
	}

	handler(eh.queue)

}
