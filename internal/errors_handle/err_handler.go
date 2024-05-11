package errors_handle

import "github.com/avp365/arch-pat/internal/entities/mov"

var c = make(map[error]func())

func init() {
	c[mov.ErrVariablePositionNotFound] = mov.ErrVariablePositionNotFoundHandler
	c[mov.ErrVariableXPositionNotFound] = mov.ErrVariableXPositionNotFoundHandler
}
func errHandler(err error) {

	handler := c[err]
	handler()
}
