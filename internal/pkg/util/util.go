package util

import (
	"errors"
	"math"
)

type Sqrteq struct {
}

func (sqrteq *Sqrteq) GetE() float64 {

	return 0.0001
}

func (sqrteq *Sqrteq) Solve(a float64, b float64, c float64) ([]float64, error) {

	var x1, x2 float64

	if isFinite(a) || isFinite(b) || isFinite(c) {
		return []float64{}, nil
	}

	if math.Abs(a) <= sqrteq.GetE() {
		return []float64{}, errors.New("something didn't work")
	}

	D := b*b - 4*a*c

	if D < -sqrteq.GetE() {
		return []float64{}, nil
	}

	if D <= sqrteq.GetE() {
		x1 = (-1 * b) / (2 * a)

		return []float64{x1, x1}, nil
	}

	if D > sqrteq.GetE() {
		x1 = (-1*b + math.Sqrt(D)) / (2 * a)

		x2 = (-1*b - math.Sqrt(D)) / (2 * a)

		return []float64{x1, x2}, nil
	}

	return []float64{}, nil
}

func isFinite(num float64) bool {
	return math.IsInf(num, 0) || math.IsNaN(num)
}
