package fuel

import (
	"errors"

	"github.com/avp365/arch-pat/internal/entities/obj"
)

type CheckFuelComamnd struct {
	Obj obj.ObjInterface
}

var ErrVariableFuelNotFound = errors.New("variable fuel not found")
var ErrVariableFuelConsuptionNotFound = errors.New("variable fuel consuption  not found")
var ErrFuelIsEmpty = errors.New("variable fuel is empty")

func (m *CheckFuelComamnd) GetFuel() (int, error) {

	f, ok := m.Obj.GetParameter("fuel").(int)

	if !ok {
		return 0, ErrVariableFuelNotFound

	}

	return f, nil
}

func (m *CheckFuelComamnd) GetFuelConsumption() (int, error) {
	fc, ok := m.Obj.GetParameter("fuelСonsumption").(int)

	if !ok {
		return 0, ErrVariableFuelConsuptionNotFound

	}

	return fc, nil
}

func (m *CheckFuelComamnd) Execute() error {
	fuel, err := m.GetFuel()

	if err != nil {
		return err

	}
	fuelCons, err := m.GetFuelConsumption()

	if err != nil {
		return err

	}

	if fuel-fuelCons < 0 {
		return ErrFuelIsEmpty
	}

	return nil
}
