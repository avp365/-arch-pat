package macro

import (
	"errors"
	"testing"

	"github.com/avp365/arch-pat/internal/command/mov"
	"github.com/stretchr/testify/assert"
)

// подготовка к новому тесту
type MockChangeVelocityComamnd_RVC1 struct {
}

func (m MockChangeVelocityComamnd_RVC1) Execute() error {
	return nil
}

type MockRotationСommand_RVC1 struct {
}

func (m MockRotationСommand_RVC1) Execute() error {
	return nil
}

// Тест, когда  у объекта есть мгновенная скорость
func TestRorationVelocity(t *testing.T) {
	check := MacroRorationVelocityCommand{MockChangeVelocityComamnd_RVC1{}, MockRotationСommand_RVC1{}}

	assert.Equal(t, check.Execute(), nil)

}

type MockRotationСommand_RVC2 struct {
}

func (m MockRotationСommand_RVC2) Execute() error {
	return errors.New("error save")
}

// Тест, когда  произошла ошибка при выполнении макрокоманды
func TestRorationVelocityFailed(t *testing.T) {
	check := MacroRorationVelocityCommand{MockChangeVelocityComamnd_RVC1{}, MockRotationСommand_RVC2{}}

	assert.EqualError(t, check.Execute(), errors.New("error save").Error())

}

type MockObj_RVC1 struct {
}

func (o MockObj_RVC1) GetParameter(name string) interface{} {

	return make(map[string]interface{})
}

func (o MockObj_RVC1) SetParameter(name string, val interface{}) error {

	if name == "position" {
		p := val.(map[string]int)

		if p["x"] == 5 && p["y"] == 8 {
			return nil
		}

	}
	return errors.New("Variable not found")
}

// Тест, когда  у объекта есть мгновенная скорость равна 0
func TestRorationVelocityIsZero(t *testing.T) {
	check := MacroRorationVelocityCommand{ChangeVelocityComamnd: mov.ChangeVelocityCommand{Obj: MockObj_RVC1{}, DX: 0, DY: 0}, RotationСommand: MockRotationСommand_RVC2{}}

	assert.EqualError(t, check.Execute(), errors.New("Variable not found").Error())

}
