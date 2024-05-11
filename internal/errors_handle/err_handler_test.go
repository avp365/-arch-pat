package errors_handle

import (
	"errors"
	"fmt"
	"testing"

	"github.com/avp365/arch-pat/internal/command"
	"github.com/stretchr/testify/assert"
)

// Реализовать Команду, которая записывает информацию о выброшенном исключении в лог.

var ErrSimple = errors.New("error simple")

var GlobalLog string

type MockWriteLogCommand struct {
}

func (m MockWriteLogCommand) Execute() error {

	//выбрасвыем ошибку
	err := ErrSimple

	//пишем в лог
	m.Log(ErrSimple)

	if err != nil {

		return err

	}

	return nil
}

// Реализовать Команду, которая записывает информацию о выброшенном исключении в лог.
func (m MockWriteLogCommand) Log(err error) {

	GlobalLog = ErrSimple.Error()
}

func ErrHandler(e chan command.Command) {

	fmt.Println("ErrHandler")
}

var mockStore = make(map[error]func(chan command.Command))

// Тест. "Реализовать Команду, которая записывает информацию о выброшенном исключении в лог.""
func TestCommandRecordLog(t *testing.T) {

	mockStore[ErrSimple] = ErrHandler

	msc := MockWriteLogCommand{}
	err := msc.Execute()

	assert.EqualError(t, err, err.Error())
	assert.EqualValues(t, GlobalLog, "error simple")

}

// Реализовать обработчик исключения, который ставит Команду, пишущую в лог в очередь Команд.
var ErrToQueueCommand = errors.New("err to queue command")

func ErrToQueueCommandRecordLog(e chan command.Command) {

	e <- MockWriteLogCommand{}

}

// Тест. "Реализовать обработчик исключения, который ставит Команду, пишущую в лог в очередь Команд""
func TestToQueueCommandRecordLog(t *testing.T) {

	//очередь (канал)
	q := make(chan command.Command, 100)

	mockStore[ErrToQueueCommand] = ErrToQueueCommandRecordLog

	cmd := MockWriteLogCommand{}

	eh := NewErrorHandler(cmd, ErrToQueueCommand, mockStore, q)
	eh.Handle()

	ncmd := <-q
	assert.EqualValues(t, ncmd, cmd)

}

// Реализовать Команду, которая повторяет Команду, выбросившую исключение.
type MockRepeatWriteLogCommand struct {
	MockWriteLog MockWriteLogCommand
}

func (m MockRepeatWriteLogCommand) Execute() error {

	err := m.MockWriteLog.Execute()

	if err != nil {

		return err

	}

	return nil
}

// Тест Реализовать Команду, которая повторяет Команду, выбросившую исключение.
func TestMockRepeatWriteLogCommand(t *testing.T) {

	msc := MockRepeatWriteLogCommand{MockWriteLogCommand{}}
	_ = msc.Execute()

	assert.EqualValues(t, GlobalLog, "error simple")

}

// Реализовать обработчик исключения, который ставит Команду, пишущую в лог в очередь Команд.
var ErrToRepeatWriteLogCommandd = errors.New("err to repeat write log command")

// Реализовать обработчик исключения, который ставит в очередь Команду - повторитель команды, выбросившей исключение.
func ErrToRepeatWriteLogCommand(e chan command.Command) {

	e <- MockRepeatWriteLogCommand{}

}

// Тест Реализовать обработчик исключения, который ставит в очередь Команду - повторитель команды, выбросившей исключение.
func TestErrToRepeatWriteLogCommandd(t *testing.T) {

	q := make(chan command.Command, 100)

	cmd := MockRepeatWriteLogCommand{}

	mockStore[ErrToRepeatWriteLogCommandd] = ErrToRepeatWriteLogCommand

	eh := NewErrorHandler(cmd, ErrToRepeatWriteLogCommandd, mockStore, q)
	eh.Handle()

	assert.EqualValues(t, GlobalLog, "error simple")

}
