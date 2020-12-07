package day4

import (
	"testing"
)

func TestPassportProcessingPartTwo(t *testing.T) {
	result := passportProcessingPartTwo("input.txt")
	t.Logf("Result found %d", result)
	t.FailNow()
}

func TestPassportProcessingPartTwoAllValid(t *testing.T) {
	result := passportProcessingPartTwo("input_all_valid.txt")
	if result != 4 {
		t.Fatalf("Incorrect value %d", result)
	}
}

func TestPassportProcessingPartTwoAllInvalid(t *testing.T) {
	result := passportProcessingPartTwo("input_all_invalid.txt")
	if result != 0 {
		t.Fatalf("Incorrect value %d", result)
	}
}
