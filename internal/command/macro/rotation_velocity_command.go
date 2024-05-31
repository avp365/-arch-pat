package macro

import (
	"github.com/avp365/arch-pat/internal/command"
)

type MacroRorationVelocityCommand struct {
	ChangeVelocityComamnd command.Command
	RotationСommand       command.Command
}

func (m *MacroRorationVelocityCommand) Execute() error {

	err := m.ChangeVelocityComamnd.Execute()

	if err != nil {
		return err
	}

	err = m.RotationСommand.Execute()

	if err != nil {
		return err
	}

	return nil
}
