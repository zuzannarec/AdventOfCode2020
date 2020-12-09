package day5

import "testing"

func TestBinaryBoardingSample(t *testing.T) {
	result := binaryBoarding("input_sample.txt")
	if result != 357 {
		t.Fatalf("Incorrect value %d", result)
	}
}

func TestBinaryBoarding(t *testing.T) {
	result := binaryBoarding("input.txt")
	t.Logf("Result found %d", result)
	t.FailNow()
}
