package car

import (
	"errors"
)

// 代表以下寫得Code都在Date的struct底下
type Data struct {
}

// 要開啟interface 來給外面訪問
type dataInterface interface {
	Division(a, b float64) (float64, error)
	DOT(a, b float64) (float64, error)
}

// 這是我原本寫的Code
func (d Data) Division(a, b float64) (float64, error) {

	if b == 0 {
		return 0, errors.New("除數不能為 0")
	}

	return a / b, nil
}

// 新寫得測試Code
func (d Data) DOT(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除數不能為 0")
	}

	return a * b, nil
}
