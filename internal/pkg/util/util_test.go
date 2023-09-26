package util

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Написать тест, который проверяет, что для уравнения x^2+1 = 0 корней нет (возвращается пустой массив)
func TestNoSolutionOfAnEquation(t *testing.T) {
	sqrteq := Sqrteq{}
	res, _ := sqrteq.Solve(1, 0, 1)
	assert.Equal(t, len(res), 0)

}

// Написать тест, который проверяет, что для уравнения x^2-1 = 0 есть два корня кратности 1 (x1=1, x2=-1)
func TestSolution1(t *testing.T) {
	sqrteq := Sqrteq{}
	res, _ := sqrteq.Solve(1, 0, -1)
	assert.Equal(t, res[0], float64(1))
	assert.Equal(t, res[1], float64(-1))

}

// Написать тест, который проверяет, что для уравнения x^2+2x+1 = 0 есть один корень кратности 2 (x1= x2 = -1).
func TestSolution2(t *testing.T) {
	sqrteq := Sqrteq{}
	res, _ := sqrteq.Solve(1, 2, 1)
	assert.Equal(t, res[0], float64(-1))
	assert.Equal(t, res[1], float64(-1))

}

// Написать тест, который проверяет, что коэффициент a не может быть равен 0. В этом случае solve выбрасывает исключение.
func TestANotEqualZero(t *testing.T) {
	sqrteq := Sqrteq{}
	_, err := sqrteq.Solve(0.0000000456, 2, 1)
	assert.EqualError(t, err, err.Error())

}

// С учетом того, что дискриминант тоже нельзя сравнивать с 0 через знак равенства, подобрать
// такие коэффициенты квадратного уравнения для случая одного корня кратности два, чтобы
// дискриминант был отличный от нуля, но меньше заданного эпсилон.
func TestDlessE(t *testing.T) {
	sqrteq := Sqrteq{}
	_, err := sqrteq.Solve(0.0000000456, 2, 1)
	assert.EqualError(t, err, err.Error())

}

// Проверка на бесконечность a
func TestIninityABC(t *testing.T) {
	sqrteq := Sqrteq{}

	res, _ := sqrteq.Solve(math.Inf(-1), 2, 1)
	assert.Equal(t, len(res), 0)

	res, _ = sqrteq.Solve(math.Inf(1), 2, 1)
	assert.Equal(t, len(res), 0)

	res, _ = sqrteq.Solve(math.Inf(-1), 2, 1)
	assert.Equal(t, len(res), 0)

	res, _ = sqrteq.Solve(1, math.Inf(1), 1)
	assert.Equal(t, len(res), 0)

	res, _ = sqrteq.Solve(1, math.Inf(-1), 1)
	assert.Equal(t, len(res), 0)

	res, _ = sqrteq.Solve(1, 2, math.Inf(1))
	assert.Equal(t, len(res), 0)

	res, _ = sqrteq.Solve(1, 2, math.Inf(-1))
	assert.Equal(t, len(res), 0)

	res, _ = sqrteq.Solve(math.NaN(), 2, 1)
	assert.Equal(t, len(res), 0)

	res, _ = sqrteq.Solve(math.NaN(), 2, 1)
	assert.Equal(t, len(res), 0)

	res, _ = sqrteq.Solve(1, math.NaN(), 1)
	assert.Equal(t, len(res), 0)

	res, _ = sqrteq.Solve(1, math.NaN(), 1)
	assert.Equal(t, len(res), 0)

	res, _ = sqrteq.Solve(1, 2, math.NaN())
	assert.Equal(t, len(res), 0)

	res, _ = sqrteq.Solve(1, 2, math.NaN())
	assert.Equal(t, len(res), 0)

}
