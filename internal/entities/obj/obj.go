package Obj

import (
	"errors"

	Movement "github.com/avp365/arch-pat/internal/entities/mov"
)

type Obj struct {
	params map[string]interface{}
}

func CreateObject() Movement.ObjInterface {
	return Obj{make(map[string]interface{})}
}

func (o Obj) GetParameter(name string) interface{} {
	return o.params[name]
}

func (o Obj) SetParameter(name string, val interface{}) error {

	o.params[name] = val
	_, ok := o.params[name]

	if ok {
		return nil
	}

	return errors.New("Variable not found")
}
