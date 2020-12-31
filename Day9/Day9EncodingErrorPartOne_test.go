package day9

import "testing"

func TestEncodingErrorSamplePartOne(t *testing.T) {
	result := encodingErrorPartOne("input_sample.txt", 5)
	if result != 127 {
		t.Fatalf("Incorrect value %d", result)
	}
}

func TestEncodingErrorPartOne(t *testing.T) {
	result := encodingErrorPartOne("input.txt", 25)
	t.Logf("Result found %d", result)
	t.FailNow()
}
