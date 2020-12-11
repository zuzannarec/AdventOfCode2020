package day6

import "testing"

func TestCustomCustomsSamplePartOne(t *testing.T) {
	result := customCustomsPartOne("input_sample.txt")
	if result != 11 {
		t.Fatalf("Incorrect value %d", result)
	}
}

func TestCustomCustomsPartOne(t *testing.T) {
	result := customCustomsPartOne("input.txt")
	t.Logf("Result found %d", result)
	t.FailNow()
}
