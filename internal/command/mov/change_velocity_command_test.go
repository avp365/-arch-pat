package mov

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// подготовка к новому тесту
type MockObj_CVC1 struct {
}

func (o MockObj_CVC1) GetParameter(name string) interface{} {
	return nil
}

func (o MockObj_CVC1) SetParameter(name string, p interface{}) error {
	pos := p.(map[string]int)

	if pos["x"] == 5 && pos["y"] == 7 {
		return nil
	}

	return errors.New("Variable error comparison")
}

// Тест, смена мгновенной скорости
func TestChangeVelocity(t *testing.T) {
	check := ChangeVelocityCommand{Obj: MockObj_CVC1{}, DX: 5, DY: 7}

	assert.Equal(t, check.Execute(), nil)

}

// подготовка к новому тесту
type MockObj_CVC2 struct {
}

func (o MockObj_CVC2) GetParameter(name string) interface{} {
	return nil
}

func (o MockObj_CVC2) SetParameter(string, interface{}) error {
	return errors.New("Variable not found")
}

// Тест, когда смена происходит с ошибкой
func TestChangeVelocityError(t *testing.T) {
	check := ChangeVelocityCommand{Obj: MockObj_CVC2{}}

	assert.EqualError(t, check.Execute(), errors.New("Variable not found").Error())

}
