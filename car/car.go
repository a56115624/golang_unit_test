package car

import (
	"errors"
)

type Data struct {
}
type dataInterface interface {
	Division(a, b float64) (float64, error)
	DOT(a, b float64) (float64, error)
}

func (d Data) Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除數不能為 0")
	}

	return a / b, nil
}
func (d Data) DOT(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除數不能為 0")
	}

	return a * b, nil
}
