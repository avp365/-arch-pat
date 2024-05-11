package errors_handle

import (
	"github.com/avp365/arch-pat/internal/command"
	"github.com/avp365/arch-pat/internal/command/mov"
)

var s = make(map[error]func(command.Command, chan command.Command))

func init() {
	s[mov.ErrVariablePositionNotFound] = mov.ErrVariablePositionNotFoundHandler
	s[mov.ErrVariableXPositionNotFound] = mov.ErrVariableXPositionNotFoundHandler
}

type ErrorHandler struct {
	cmd   command.Command
	store map[error]func(command.Command, chan command.Command)
	e     error
	queue chan command.Command
}

func NewErrorHandler(cmd command.Command, err error, s map[error]func(command.Command, chan command.Command), q chan command.Command) ErrorHandler {

	return ErrorHandler{cmd: cmd, e: err, store: s, queue: q}
}

func (eh *ErrorHandler) Handle() {

	handler, ok := eh.store[eh.e]

	if !ok {
		panic("not error handler")
	}

	handler(eh.cmd, eh.queue)

}
