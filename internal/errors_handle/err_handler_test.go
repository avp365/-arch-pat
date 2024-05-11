package errors_handle

import (
	"testing"

	"github.com/avp365/arch-pat/internal/entities/mov"
)

type MockMoveE1 struct {
}

func TestMoveError(t *testing.T) {

	errHandler(mov.ErrVariablePositionNotFound)

	//assert.EqualError(t, err, err.Error())

}
