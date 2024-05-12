package errors_handle

import (
	"github.com/avp365/arch-pat/internal/command"
	"github.com/avp365/arch-pat/internal/command/mov"
)

var s = make(map[error]func(command.Command, map[string]interface{}))

func init() {
	s[mov.ErrVariablePositionNotFound] = mov.ErrVariablePositionNotFoundHandler
	s[mov.ErrVariableXPositionNotFound] = mov.ErrVariableXPositionNotFoundHandler
}

type ErrorHandler struct {
	cmd   command.Command
	store map[error]func(command.Command, map[string]interface{})
	e     error
	data  map[string]interface{}
}

func NewErrorHandler(cmd command.Command, err error, s map[error]func(command.Command, map[string]interface{}), data map[string]interface{}) ErrorHandler {

	return ErrorHandler{cmd: cmd, e: err, store: s, data: data}
}

func (eh *ErrorHandler) Handle() {

	handler, ok := eh.store[eh.e]

	if !ok {
		panic("not error handler")
	}

	handler(eh.cmd, eh.data)

}
