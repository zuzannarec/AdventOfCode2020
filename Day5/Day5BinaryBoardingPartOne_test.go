package day5

import "testing"

func TestBinaryBoardingSamplePartOne(t *testing.T) {
	result := binaryBoardingPartOne("input_sample.txt")
	if result != 357 {
		t.Fatalf("Incorrect value %d", result)
	}
}

func TestBinaryBoardingPartOne(t *testing.T) {
	result := binaryBoardingPartOne("input.txt")
	t.Logf("Result found %d", result)
	t.FailNow()
}
