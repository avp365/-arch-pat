package mov

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockObjA1 struct {
}

func (o MockObjA1) GetParameter(name string) interface{} {

	params := make(map[string]interface{})

	params["direction"] = 180
	params["angularVelocity"] = 20
	params["directionNumber"] = 60

	return params[name]
}

func (o MockObjA1) SetParameter(name string, val interface{}) error {

	if name == "direction" && val == 20 {
		return nil

	}

	return errors.New("Variable not found")
}

// /Для объекта, 360
func TestRotation(t *testing.T) {
	move := Rotation{Obj: MockObjA1{}}

	assert.Equal(t, move.Execute(), nil)

}

type MockObjA2 struct {
}

func (o MockObjA2) GetParameter(name string) interface{} {

	params := make(map[string]interface{})

	params["angularVelocity"] = 20
	params["directionNumber"] = 60

	return params[name]
}
func (o MockObjA2) SetParameter(name string, val interface{}) error {

	if name == "position" {
		p := val.(map[string]int)

		if p["x"] == 5 && p["y"] == 8 {
			return nil
		}

	}
	return errors.New("Variable not found")
}

// Не возможно прочитать direction
func TestDirectionnNotFound(t *testing.T) {
	move := Rotation{Obj: MockObjA2{}}

	err := move.Execute()

	assert.EqualError(t, err, err.Error())

}
