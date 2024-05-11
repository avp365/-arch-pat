package Obj

import (
	"testing"

	Movement "github.com/avp365/arch-pat/internal/entities/mov"
	"github.com/stretchr/testify/assert"
)

func TestCreateObject(t *testing.T) {
	obj := CreateObject()

	assert.Implements(t, (*Movement.ObjInterface)(nil), obj)

}

func TestGetSetParameterInt(t *testing.T) {
	obj := CreateObject()

	obj.SetParameter("val", 1)

	assert.Equal(t, obj.GetParameter("val"), int(1))

}
func TestGetSetParameterString(t *testing.T) {
	obj := CreateObject()

	obj.SetParameter("val", "s")

	assert.Equal(t, obj.GetParameter("val"), "s")

}
