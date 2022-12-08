package main

import "testing"

func TestGetWeekDay(t *testing.T) {
	// a subtest named "index=0"
	t.Run("index=0", func(t *testing.T) {
		index := 0
		want := "Sunday"
		if got := GetWeekDay(index); got != want {
			t.Errorf("GetWeekDay() = %v, want %v", got, want)
		}
	})

	// a subtest named "index=1"
	t.Run("index=1", func(t *testing.T) {
		index := 1
		want := "Monday"
		if got := GetWeekDay(index); got != want {
			t.Errorf("GetWeekDay() = %v, want %v", got, want)
		}
	})
}
