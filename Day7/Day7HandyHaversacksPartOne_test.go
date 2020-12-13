package day7

import "testing"

func TestHandyHaversacksSamplePartOne(t *testing.T) {
	result := handyHaversacksPartOne("input_sample.txt", "shiny gold")
	if result != 4 {
		t.Fatalf("Incorrect value %d", result)
	}
}

func TestHandyHaversacksPartOne(t *testing.T) {
	result := handyHaversacksPartOne("input.txt", "shiny gold")
	t.Logf("Result found %d", result)
	t.FailNow()
}
