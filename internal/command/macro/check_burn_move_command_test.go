package macro

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// подготовка к новому тесту
type MockCheckFuelComamnd_CBMC1 struct {
}

func (m MockCheckFuelComamnd_CBMC1) Execute() error {
	return nil
}

type MockBurnFuelComamnd_CBMC1 struct {
}

func (m MockBurnFuelComamnd_CBMC1) Execute() error {
	return nil
}

type MockMoveComamnd_CBMC1 struct {
}

func (m MockMoveComamnd_CBMC1) Execute() error {
	return nil
}

// Тест, когда  у объекта достаточно горючего
func TestMoveSuccess(t *testing.T) {
	check := MacroCheckBurnMoveComamnd{MockCheckFuelComamnd_CBMC1{}, MockBurnFuelComamnd_CBMC1{}, MockMoveComamnd_CBMC1{}}

	assert.Equal(t, check.Execute(), nil)

}

type MockMoveComamnd_CBMC2 struct {
}

func (m MockMoveComamnd_CBMC2) Execute() error {
	return errors.New("error save")
}

// Тест, когда  сжигание произошло с ошибкой
func TestMoveFailed(t *testing.T) {
	check := MacroCheckBurnMoveComamnd{MockCheckFuelComamnd_CBMC1{}, MockBurnFuelComamnd_CBMC1{}, MockMoveComamnd_CBMC2{}}

	assert.EqualError(t, check.Execute(), errors.New("error save").Error())

}
