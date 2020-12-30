package day7

import "testing"

func TestHandyHaversacksSamplePartTwo(t *testing.T) {
	result := handyHaversacksPartTwo("input_sample2.txt", "shiny gold")
	if result != 126 {
		t.Fatalf("Incorrect value %d", result)
	}
}

func TestHandyHaversacksPartTwo(t *testing.T) {
	result := handyHaversacksPartTwo("input.txt", "shiny gold")
	t.Logf("Result found %d", result)
	t.FailNow()
}
