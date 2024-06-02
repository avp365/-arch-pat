package mov

import (
	"github.com/avp365/arch-pat/internal/entities/obj"
)

type ChangeVelocityCommand struct {
	Obj obj.ObjInterface
	DX  int
	DY  int
}

func (m *ChangeVelocityCommand) SetVelocity(x int, y int) error {

	p := map[string]int{}

	p["x"] = x
	p["y"] = y

	return m.Obj.SetParameter("velocity", p)
}

func (m ChangeVelocityCommand) Execute() error {

	err := m.SetVelocity(m.DX, m.DY)

	if err != nil {
		return err

	}
	return nil
}
