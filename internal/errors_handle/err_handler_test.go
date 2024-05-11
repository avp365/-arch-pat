package errors_handle

import (
	"errors"
	"testing"

	"github.com/avp365/arch-pat/internal/command"
	"github.com/stretchr/testify/assert"
)

var mockStore = make(map[error]func(command.Command, chan command.Command))

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

// 4 Реализовать Команду, которая записывает информацию о выброшенном исключении в лог.

func (m MockWriteLogCommand) Log(err error) {

	GlobalLog = ErrToRepeatWriteLogCommand.Error()
}

func ErrHandler(e command.Command, q chan command.Command) {

	q <- MockWriteLogCommand{}
}

// Тест. "Реализовать Команду, которая записывает информацию о выброшенном исключении в лог.""
func TestCommandRecordLog(t *testing.T) {

	mockStore[ErrSimple] = ErrHandler

	msc := MockWriteLogCommand{}
	err := msc.Execute()

	assert.EqualError(t, err, err.Error())
	assert.EqualValues(t, GlobalLog, "err to repeat write log command")

}

// 5 Реализовать обработчик исключения, который ставит Команду, пишущую в лог в очередь Команд.
var ErrToQueueCommand = errors.New("err to queue command")

func ErrToQueueCommandRecordLog(e command.Command, q chan command.Command) {

	q <- MockWriteLogCommand{}

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

// 6 Реализовать Команду, которая повторяет Команду, выбросившую исключение.
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

	assert.EqualValues(t, GlobalLog, "err to repeat write log command")

}

// 7 Реализовать обработчик исключения, который ставит Команду, пишущую в лог в очередь Команд.
var ErrToRepeatWriteLogCommand = errors.New("err to repeat write log command")

// Реализовать обработчик исключения, который ставит в очередь Команду - повторитель команды, выбросившей исключение.
func ErrToRepeatWriteLogCommandHandle(e command.Command, q chan command.Command) {

	q <- MockRepeatWriteLogCommand{e}

}

// Тест Реализовать обработчик исключения, который ставит в очередь Команду - повторитель команды, выбросившей исключение.
func TestErrToRepeatWriteLogCommand(t *testing.T) {

	mockStore[ErrToRepeatWriteLogCommand] = ErrToRepeatWriteLogCommandHandle

	q := make(chan command.Command, 100)

	cmd := MockRepeatWriteLogCommand{}

	eh := NewErrorHandler(cmd, ErrToRepeatWriteLogCommand, mockStore, q)
	eh.Handle()

	assert.EqualValues(t, GlobalLog, "err to repeat write log command")

}

// 8 С помощью Команд из пункта 4 и пункта 6 реализовать следующую обработку исключений:

func TestErrToRepeatIfErrorWriteLogCommand(t *testing.T) {
	mockStore[ErrToRepeatWriteLogCommand] = ErrToRepeatWriteLogCommandHandle

	q := make(chan command.Command, 100)
	q <- MockWriteLogCommand{}

	for i := 0; i < 3; i++ {

		cmd := <-q

		if cmd != nil {
			err := cmd.Execute()

			if err != nil {
				eh := NewErrorHandler(cmd, err, mockStore, q)
				eh.Handle()
			}
		}

	}

	assert.EqualValues(t, GlobalLog, "err to repeat write log command")

}

// 8 С помощью Команд из пункта 4 и пункта 6 реализовать следующую обработку исключений:
type MockTwoWriteLogCommand struct {
}

func (m MockTwoWriteLogCommand) Execute() error {

	//выбрасываем ошибку
	err := ErrToTwoRepeatWriteLogCommand

	//пишем в лог
	m.Log(ErrToTwoRepeatWriteLogCommand)

	if err != nil {

		return err

	}

	return nil
}
func (m MockTwoWriteLogCommand) Log(err error) {

	GlobalLog = err.Error()
}

type MockTwoRepeatIfErrorWriteLogCommand struct {
	MockCommand command.Command
}

func (m MockTwoRepeatIfErrorWriteLogCommand) Execute() error {

	err := m.MockCommand.Execute()
	if err != nil {
		err = m.MockCommand.Execute()
	}

	if err != nil {
		err = m.MockCommand.Execute()
	}

	m.Log(ErrToRepeatWriteLogCommand)

	return err

}
func (m MockTwoRepeatIfErrorWriteLogCommand) Log(err error) {

	GlobalLog = err.Error()
}

var ErrToTwoRepeatWriteLogCommand = errors.New("err to two repeat write log command")

// 9 Реализовать обработчик исключения, который ставит Команду, пишущую в лог в очередь Команд.
func ErrToTwoRepeatWriteLogCommandHandle(e command.Command, q chan command.Command) {

	q <- MockRepeatWriteLogCommand{e}

}

// Реализовать стратегию обработки исключения - повторить два раза, потом записать в лог.
func TestErrToTwoRepeatIfErrorWriteLogCommand(t *testing.T) {
	mockStore[ErrToTwoRepeatWriteLogCommand] = ErrToTwoRepeatWriteLogCommandHandle

	q := make(chan command.Command, 100)
	q <- MockTwoWriteLogCommand{}

	for i := 0; i < 3; i++ {

		cmd := <-q

		if cmd != nil {
			err := cmd.Execute()

			if err != nil {
				eh := NewErrorHandler(cmd, err, mockStore, q)
				eh.Handle()
			}
		}

	}

	assert.EqualValues(t, GlobalLog, "err to two repeat write log command")

}
