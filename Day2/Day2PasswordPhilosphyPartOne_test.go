package day2

import "testing"

func TestPasswordPhilosphy(t *testing.T) {
	result := passwordPhilosphyPartOne("input.txt")
	if result != 0 {
		t.Logf("Result found %d", result)
	}
	t.FailNow()
}
