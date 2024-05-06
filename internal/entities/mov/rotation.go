package mov

import (
	"errors"
)

type Rotation struct {
	Obj ObjInterface
}

func (m *Rotation) GetDirection() (int, error) {

	d, ok := m.Obj.GetParameter("direction").(int)

	if !ok {
		return 0, errors.New("variable direction not found")
	}

	return d, nil
}

func (m *Rotation) GetAngularVelocity() (int, error) {

	av, ok := m.Obj.GetParameter("angularVelocity").(int)

	if !ok {
		return 0, errors.New("variable angularVelocity not found")
	}

	return av, nil
}

func (m *Rotation) GetDirectionNumber() (int, error) {

	dN, ok := m.Obj.GetParameter("directionNumber").(int)

	if !ok {
		return 0, errors.New("variable directionNumber not found")
	}

	return dN, nil
}

func (m *Rotation) SetDirection(d int) error {

	return m.Obj.SetParameter("direction", d)
}

func (m *Rotation) Execute() error {

	d, err := m.GetDirection()

	if err != nil {
		return err
	}

	vA, err := m.GetAngularVelocity()

	if err != nil {
		return err
	}

	dN, err := m.GetDirectionNumber()

	if err != nil {
		return err
	}

	err = m.SetDirection((d + vA) % dN)

	if err != nil {
		return err

	}
	return nil
}
