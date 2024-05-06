package mov

import (
	"errors"
)

type Move struct {
	Obj ObjInterface
}

func (m *Move) GetPosition() (Pos, error) {

	p, ok := m.Obj.GetParameter("position").(map[string]int)

	if !ok {
		return Pos{}, errors.New("variable position not found")
	}

	_, ok = p["x"]

	if !ok {
		return Pos{}, errors.New("variable x position not found")
	}

	_, ok = p["y"]

	if !ok {
		return Pos{}, errors.New("variable y position not found")
	}

	return Pos{X: p["x"], Y: p["y"]}, nil
}

func (m *Move) GetVelocity() (Pos, error) {

	p, ok := m.Obj.GetParameter("velocity").(map[string]int)

	if !ok {
		return Pos{}, errors.New("variable velocity not found")
	}

	_, ok = p["x"]

	if !ok {
		return Pos{}, errors.New("variable x velocity not found")
	}

	_, ok = p["y"]

	if !ok {
		return Pos{}, errors.New("variable y velocity not found")
	}

	return Pos{X: p["x"], Y: p["y"]}, nil
}

func (m *Move) SetPosition(x int, y int) error {

	p := map[string]int{}

	p["x"] = x
	p["y"] = y

	return m.Obj.SetParameter("position", p)
}

func (m *Move) Execute() error {

	pos, err := m.GetPosition()

	if err != nil {
		return err
	}

	dpos, err := m.GetVelocity()

	if err != nil {
		return err
	}

	err = m.SetPosition(pos.X+dpos.X, pos.Y+dpos.Y)

	if err != nil {
		return err

	}
	return nil
}
