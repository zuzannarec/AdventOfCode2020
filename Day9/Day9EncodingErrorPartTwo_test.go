package day9

import "testing"

func TestEncodingErrorSamplePartTwo(t *testing.T) {
	result := encodingErrorPartTwo("input_sample.txt", 5)
	if result != 62 {
		t.Fatalf("Incorrect value %d", result)
	}
}

func TestEncodingErrorPartTwo(t *testing.T) {
	result := encodingErrorPartTwo("input.txt", 25)
	t.Logf("Result found %d", result)
	t.FailNow()
}
