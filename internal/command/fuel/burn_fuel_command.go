package fuel

import (
	"errors"

	"github.com/avp365/arch-pat/internal/entities/obj"
)

var ErrVariableFuelNotFoundBurnFuel = errors.New("variable fuel not found with burn")
var ErrVariableFuelConsuptionNotFoundBurnFuel = errors.New("variable fuel consuption  not found with burn")
var ErrFuelIsEmptyBurnFuel = errors.New("variable fuel is empty with burn")

type BurnFuelComamnd struct {
	Obj obj.ObjInterface
}

func (m *BurnFuelComamnd) Execute() error {

	f, ok := m.Obj.GetParameter("fuel").(int)

	if !ok {
		return ErrVariableFuelNotFoundBurnFuel

	}

	fc, ok := m.Obj.GetParameter("fuelСonsumption").(int)

	if !ok {
		return ErrVariableFuelConsuptionNotFoundBurnFuel

	}

	fn := f - fc

	err := m.Obj.SetParameter("fuel", fn)

	if err != nil {
		return err

	}

	return nil
}
