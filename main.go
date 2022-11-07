package main

import (
	"fmt"
	"golang_shane_test/car"
)

func main() {
	a := car.UseData{D: car.Data{}}
	c, err := a.UseDivision(1, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)
}
