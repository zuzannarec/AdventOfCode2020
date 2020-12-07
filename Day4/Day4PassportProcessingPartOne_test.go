package day4

import "testing"

func TestPassportProcessingPartOne(t *testing.T) {
	result := passportProcessingPartOne("input.txt")
	t.Logf("Result found %d", result)
	t.FailNow()
}

func TestPassportProcessingPartOneSample(t *testing.T) {
	result := passportProcessingPartOne("input_sample.txt")
	if result != 2 {
		t.Fatalf("Incorrect value %d", result)
	}
}
