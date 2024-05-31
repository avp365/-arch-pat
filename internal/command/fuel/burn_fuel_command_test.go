package fuel

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// подготовка к новому тесту
type MockObj_BFC1 struct {
}

func (o MockObj_BFC1) GetParameter(name string) interface{} {

	params := make(map[string]interface{})

	params["fuel"] = 100
	params["fuelСonsumption"] = 10

	return params[name]
}
func (o MockObj_BFC1) SetParameter(string, interface{}) error {
	return nil
}

// Тест, когда  у объекта достаточно горючего
func TestFuelBurn(t *testing.T) {
	check := BurnFuelComamnd{Obj: MockObj_BFC1{}}

	assert.Equal(t, check.Execute(), nil)

}

// подготовка к новому тесту
type MockObj_BFC2 struct {
}

func (o MockObj_BFC2) GetParameter(name string) interface{} {

	params := make(map[string]interface{})

	params["fuel"] = 0
	params["fuelСonsumption"] = 10

	return params[name]
}
func (o MockObj_BFC2) SetParameter(string, interface{}) error {
	return errors.New("error save")
}

// Тест, когда  сжигание произошло с ошибкой
func TestFuelIsBurnError(t *testing.T) {
	check := BurnFuelComamnd{Obj: MockObj_BFC2{}}

	assert.EqualError(t, check.Execute(), errors.New("error save").Error())

}
