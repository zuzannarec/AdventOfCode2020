package day10

import "testing"

func TestAdapterArraySampleShortPartOne(t *testing.T) {
	result := adapterArrayPartOne("input_sample_short.txt")
	if result != 35 {
		t.Fatalf("Incorrect value %d", result)
	}
}
func TestAdapterArraySamplePartOne(t *testing.T) {
	result := adapterArrayPartOne("input_sample.txt")
	if result != 220 {
		t.Fatalf("Incorrect value %d", result)
	}
}

func TestAdapterArrayPartOne(t *testing.T) {
	result := adapterArrayPartOne("input.txt")
	t.Logf("Result found %d", result)
	t.FailNow()
}
