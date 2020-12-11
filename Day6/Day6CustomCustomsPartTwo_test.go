package day6

import "testing"

func TestCustomCustomsSamplePartTwo(t *testing.T) {
	result := customCustomsPartTwo("input_sample.txt")
	if result != 6 {
		t.Fatalf("Incorrect value %d", result)
	}
}

func TestCustomCustomsPartTwo(t *testing.T) {
	result := customCustomsPartTwo("input.txt")
	t.Logf("Result found %d", result)
	t.FailNow()
}
