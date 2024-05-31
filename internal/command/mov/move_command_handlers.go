package mov

import (
	"fmt"

	"github.com/avp365/arch-pat/internal/command"
)

func ErrVariablePositionNotFoundHandler(command.Command, map[string]interface{}) {

	fmt.Println("ErrVariablePositionNotFoundHandler")
}

func ErrVariableXPositionNotFoundHandler(command.Command, map[string]interface{}) {

	fmt.Println("ErrVariableXPositionNotFoundHandler")
}
