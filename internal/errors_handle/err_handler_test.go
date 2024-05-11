package errors_handle

import (
	"errors"
	"testing"

	"github.com/avp365/arch-pat/internal/command"
	"github.com/stretchr/testify/assert"
)

var mockStore = make(map[error]func(chan command.Command))

// Реализовать Команду, которая записывает информацию о выброшенном исключении в лог.

var ErrSimple = errors.New("error simple")

var GlobalLog string

type MockWriteLogCommand struct {
}

func (m MockWriteLogCommand) Execute() error {

	//выбрасываем ошибку
	err := ErrToRepeatWriteLogCommand

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

	e <- MockWriteLogCommand{}
}

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
	MockCommand command.Command
}

func (m MockRepeatWriteLogCommand) Execute() error {
	err := m.MockCommand.Execute()

	return err

}
func (m MockRepeatWriteLogCommand) Log(err error) {

	GlobalLog = ErrSimple.Error()
}

// Тест Реализовать Команду, которая повторяет Команду, выбросившую исключение.
func TestMockRepeatWriteLogCommand(t *testing.T) {

	msc := MockRepeatWriteLogCommand{MockWriteLogCommand{}}
	_ = msc.Execute()

	assert.EqualValues(t, GlobalLog, "error simple")

}

// Реализовать обработчик исключения, который ставит Команду, пишущую в лог в очередь Команд.
var ErrToRepeatWriteLogCommand = errors.New("err to repeat write log command")

// Реализовать обработчик исключения, который ставит в очередь Команду - повторитель команды, выбросившей исключение.
func ErrToRepeatWriteLogCommandHandle(e chan command.Command) {

	e <- MockRepeatWriteLogCommand{}

}

// Тест Реализовать обработчик исключения, который ставит в очередь Команду - повторитель команды, выбросившей исключение.
func TestErrToRepeatWriteLogCommand(t *testing.T) {

	mockStore[ErrToRepeatWriteLogCommand] = ErrToRepeatWriteLogCommandHandle

	q := make(chan command.Command, 100)

	cmd := MockRepeatWriteLogCommand{}

	eh := NewErrorHandler(cmd, ErrToRepeatWriteLogCommand, mockStore, q)
	eh.Handle()

	assert.EqualValues(t, GlobalLog, "error simple")

}

// С помощью Команд из пункта 4 и пункта 6 реализовать следующую обработку исключений:
func TestErrToRepeatIfErrorWriteLogCommand(t *testing.T) {
	mockStore[ErrToRepeatWriteLogCommand] = ErrToRepeatWriteLogCommandHandle

	q := make(chan command.Command, 100)
	q <- MockWriteLogCommand{}

	for i := 0; i < 3; i++ {

		cmd := <-q
		err := cmd.Execute()

		if err != nil {
			eh := NewErrorHandler(cmd, err, mockStore, q)
			eh.Handle()
		}

	}

	assert.EqualValues(t, GlobalLog, "error simple")

}
