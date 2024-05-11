package mov

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockObj1 struct {
}

func (o MockObj1) GetParameter(name string) interface{} {

	params := make(map[string]interface{})

	pos := map[string]int{}
	pos["x"] = 12
	pos["y"] = 5

	params["position"] = pos

	vel := map[string]int{}
	vel["x"] = -7
	vel["y"] = 3

	params["velocity"] = vel

	return params[name]
}

func (o MockObj1) SetParameter(name string, val interface{}) error {

	if name == "position" {
		p := val.(map[string]int)

		if p["x"] == 5 && p["y"] == 8 {
			return nil
		}

	}
	return errors.New("Variable not found")
}

// /Для объекта, находящегося в точке (12, 5) и движущегося со скоростью (-7, 3) движение меняет положение объекта на (5, 8)
func TestMoveVector(t *testing.T) {
	move := Move{Obj: MockObj1{}}

	assert.Equal(t, move.Execute(), nil)

}

type MockObj2 struct {
}

func (o MockObj2) GetParameter(name string) interface{} {

	return make(map[string]interface{})
}

func (o MockObj2) SetParameter(name string, val interface{}) error {

	if name == "position" {
		p := val.(map[string]int)

		if p["x"] == 5 && p["y"] == 8 {
			return nil
		}

	}
	return errors.New("Variable not found")
}

// Попытка сдвинуть объект, у которого невозможно прочитать положение в пространстве, приводит к ошибке
func TestPositionNotFound(t *testing.T) {
	move := Move{Obj: MockObj2{}}

	err := move.Execute()

	assert.EqualError(t, err, err.Error())

}

type MockObj3 struct {
}

func (o MockObj3) GetParameter(name string) interface{} {

	params := make(map[string]interface{})

	pos := map[string]int{}
	pos["x"] = 12
	pos["y"] = 5

	params["position"] = pos

	return params[name]
}

func (o MockObj3) SetParameter(name string, val interface{}) error {

	if name == "position" {
		p := val.(map[string]int)

		if p["x"] == 5 && p["y"] == 8 {
			return nil
		}

	}
	return errors.New("Variable not found")
}

// Попытка сдвинуть объект, у которого невозможно прочитать значение мгновенной скорости, приводит к ошибке
func TestVelocityNotFound(t *testing.T) {
	move := Move{Obj: MockObj3{}}

	err := move.Execute()

	assert.EqualError(t, err, err.Error())

}

type MockObj4 struct {
}

func (o MockObj4) GetParameter(name string) interface{} {

	params := make(map[string]interface{})

	pos := map[string]int{}
	pos["x"] = 12
	pos["y"] = 5

	params["position"] = pos

	vel := map[string]int{}
	vel["x"] = -7
	vel["y"] = 3

	params["velocity"] = vel

	return params[name]
}

func (o MockObj4) SetParameter(name string, val interface{}) error {

	return errors.New("Variable not found")
}

// Попытка сдвинуть объект, у которого невозможно изменить положение в пространстве, приводит к ошибке
func TestChangePositionObject(t *testing.T) {
	move := Move{Obj: MockObj4{}}

	err := move.Execute()

	assert.EqualError(t, err, err.Error())

}
