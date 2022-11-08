package main

import (
	"fmt"
	_ "fmt"
	_ "golang_shane_test/car"
)

// GetWeekDay returns the week day name of a week day index.
func main() {
	fmt.Println(GetWeekDay(2))
}
func GetWeekDay(index int) string {
	if index < 0 || index > 6 {
		return "Unknown"
	}
	weekDays := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	return weekDays[index]
}

// func main() {
// 	// a := car.UseData{D: car.Data{}}
// 	// c, err := a.UseDivision(1, 2)
// 	//
// 	//	if err != nil {
// 	//		fmt.Println(err)
// 	//	}
// 	//
// 	// fmt.Println(c)
// }
