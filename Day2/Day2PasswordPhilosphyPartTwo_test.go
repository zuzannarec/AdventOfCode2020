package day2

import "testing"

func TestPasswordPhilosphyPartTwo(t *testing.T) {
	result := passwordPhilosphyPartTwo("input.txt")
	if result != 0 {
		t.Logf("Result found %d", result)
	}
	t.FailNow()
}
