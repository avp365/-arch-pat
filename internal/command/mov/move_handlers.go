package mov

import (
	"fmt"

	"github.com/avp365/arch-pat/internal/command"
)

func ErrVariablePositionNotFoundHandler(q chan command.Command) {

	fmt.Println("ErrVariablePositionNotFoundHandler")
}

func ErrVariableXPositionNotFoundHandler(q chan command.Command) {

	fmt.Println("ErrVariableXPositionNotFoundHandler")
}
