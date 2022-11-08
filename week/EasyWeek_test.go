package main

import "testing"

// test case for index=0
func TestGetWeekDay_Sunday(t *testing.T) {
	index := 0
	want := "Sunday"
	if got := GetWeekDay(index); got != want {
		t.Errorf("GetWeekDay() = %v, want %v", got, want)
	}
}

// test case for index=1
func TestGetWeekDay_Monday(t *testing.T) {
	index := 1
	want := "Monday"
	if got := GetWeekDay(index); got != want {
		t.Errorf("GetWeekDay() = %v, want %v", got, want)
	}
}
