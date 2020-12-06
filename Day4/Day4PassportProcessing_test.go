package day3

import "testing"

func TestTobogganTrajectory(t *testing.T) {
	result := passportProcessing("input.txt")
	t.Logf("Result found %d", result)
	t.FailNow()
}

func TestTobogganTrajectorySample(t *testing.T) {
	result := passportProcessing("input_sample.txt")
	if result != 2 {
		t.Fatalf("Incorrect value %d", result)
	}
}
