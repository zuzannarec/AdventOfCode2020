package day5

import "testing"

func TestBinaryBoardingPartTwo(t *testing.T) {
	result := binaryBoardingPartTwo("input.txt")
	t.Logf("Result found %d", result)
	t.FailNow()
}
