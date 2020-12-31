package day8

import "testing"

func TestHandheldHaltingSamplePartTwo(t *testing.T) {
	result := handheldHaltingPartTwo("input_sample.txt")
	if result != 8 {
		t.Fatalf("Incorrect value %d", result)
	}
}

func TestHandheldHaltingPartTwo(t *testing.T) {
	result := handheldHaltingPartTwo("input.txt")
	t.Logf("Result found %d", result)
	t.FailNow()
}
