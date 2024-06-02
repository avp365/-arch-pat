package macro

import (
	"errors"

	"github.com/avp365/arch-pat/internal/command"
)

var ErrVariableFuelNotFoundBurnFuel = errors.New("variable fuel not found with burn")
var ErrVariableFuelConsuptionNotFoundBurnFuel = errors.New("variable fuel consuption  not found with burn")
var ErrFuelIsEmptyBurnFuel = errors.New("variable fuel is empty with burn")

type MacroCheckBurnMoveComamnd struct {
	CheckFuelComamnd command.Command
	BurnFuelComamnd  command.Command
	MoveComamnd      command.Command
}

func (m *MacroCheckBurnMoveComamnd) Execute() error {

	err := m.CheckFuelComamnd.Execute()

	if err != nil {
		return err
	}

	err = m.BurnFuelComamnd.Execute()

	if err != nil {
		return err
	}

	err = m.MoveComamnd.Execute()

	if err != nil {
		return err
	}

	return nil
}
