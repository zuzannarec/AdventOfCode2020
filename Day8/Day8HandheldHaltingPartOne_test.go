package day8

import "testing"

func TestHandheldHaltingSamplePartOne(t *testing.T) {
	result := handheldHaltingPartOne("input_sample.txt")
	if result != 5 {
		t.Fatalf("Incorrect value %d", result)
	}
}

func TestHandheldHaltingPartOne(t *testing.T) {
	result := handheldHaltingPartOne("input.txt")
	t.Logf("Result found %d", result)
	t.FailNow()
}
