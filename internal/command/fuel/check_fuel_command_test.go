package fuel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// подготовка к новому тесту
type MockObj_CFC1 struct {
}

func (o MockObj_CFC1) GetParameter(name string) interface{} {

	params := make(map[string]interface{})

	params["fuel"] = 100
	params["fuelСonsumption"] = 10

	return params[name]
}
func (o MockObj_CFC1) SetParameter(string, interface{}) error {
	return nil
}

// Тест, когда  у объекта достаточно горючего
func TestFuelNotEmpty(t *testing.T) {
	check := CheckFuelComamnd{Obj: MockObj_CFC1{}}

	assert.Equal(t, check.Execute(), nil)

}

// подготовка к новому тесту
type MockObj_CFC2 struct {
}

func (o MockObj_CFC2) GetParameter(name string) interface{} {

	params := make(map[string]interface{})

	params["fuel"] = 0
	params["fuelСonsumption"] = 10

	return params[name]
}
func (o MockObj_CFC2) SetParameter(string, interface{}) error {
	return nil
}

// Тест, когда  у объекта недостаточно горючего
func TestFuelIsEmpty(t *testing.T) {
	check := CheckFuelComamnd{Obj: MockObj_CFC2{}}

	assert.EqualError(t, check.Execute(), ErrFuelIsEmpty.Error())

}
