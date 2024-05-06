package entities

// Для объекта, находящегося в точке (12, 5) и движущегося со скоростью (-7, 3) движение меняет положение объекта на (5, 8)
// func TestNoSolutionOfAnEquation(t *testing.T) {
// 	sqrteq := Sqrteq{}
// 	res, _ := sqrteq.Solve(1, 0, 1)
// 	assert.Equal(t, len(res), 0)

// }

// Попытка сдвинуть объект, у которого невозможно прочитать положение в пространстве, приводит к ошибке
// func TestSolution1(t *testing.T) {
// 	sqrteq := Sqrteq{}
// 	res, _ := sqrteq.Solve(1, 0, -1)
// 	assert.Equal(t, res[0], float64(1))
// 	assert.Equal(t, res[1], float64(-1))

// }

// Попытка сдвинуть объект, у которого невозможно прочитать значение мгновенной скорости, приводит к ошибке
// func TestSolution2(t *testing.T) {
// 	sqrteq := Sqrteq{}
// 	res, _ := sqrteq.Solve(1, 2, 1)
// 	assert.Equal(t, res[0], float64(-1))
// 	assert.Equal(t, res[1], float64(-1))

// }

// Попытка сдвинуть объект, у которого невозможно изменить положение в пространстве, приводит к ошибке
// func TestANotEqualZero(t *testing.T) {
// 	sqrteq := Sqrteq{}
// 	_, err := sqrteq.Solve(0.0000000456, 2, 1)
// 	assert.EqualError(t, err, err.Error())

// }
