package errors_handle

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Реализовать Команду, которая записывает информацию о выброшенном исключении в лог.

var ErrSimple = errors.New("error simple")

var GlobalLog string

type MockSimpleCommand struct {
}

func (m *MockSimpleCommand) Execute() error {

	//выбрасвыем ошибку
	err := ErrSimple

	//пишем в лог
	m.Log(ErrSimple)

	if err != nil {

		return err

	}

	return nil
}

func (m *MockSimpleCommand) Log(err error) {

	GlobalLog = ErrSimple.Error()
}

func ErrSimpleHandler() {

	fmt.Println("ErrSimpleHandler")
}

var cm = make(map[error]func())

func TestSimpleCommandRecordLog(t *testing.T) {

	cm[ErrSimple] = ErrSimpleHandler

	msc := MockSimpleCommand{}
	err := msc.Execute()

	assert.EqualError(t, err, err.Error())
	assert.EqualValues(t, GlobalLog, "error simple")

}

// func TestSimpleCommandRecordLog(t *testing.T) {

// 	cm[ErrSimple] = ErrSimpleHandler

// 	msc := MockSimpleCommand{}
// 	err := msc.Execute()

// 	if err != nil {
// 		eh := NewErrorHandler(ErrSimple, cm)
// 		eh.Handle()
// 	}

// 	assert.EqualError(t, err, err.Error())

// }
