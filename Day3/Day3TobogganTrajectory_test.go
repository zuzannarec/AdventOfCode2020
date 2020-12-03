package day3

import "testing"

func TestTobogganTrajectory(t *testing.T) {
	result11 := tobogganTrajectory("input.txt", 1, 1)
	result31 := tobogganTrajectory("input.txt", 3, 1)
	result51 := tobogganTrajectory("input.txt", 5, 1)
	result71 := tobogganTrajectory("input.txt", 7, 1)
	result12 := tobogganTrajectory("input.txt", 1, 2)
	result := result11 * result31 * result51 * result71 * result12
	t.Logf("Result found %d", result)
	t.FailNow()
}

func TestTobogganTrajectorySample(t *testing.T) {
	result := tobogganTrajectory("input_sample.txt", 3, 1)
	if result != 7 {
		t.Fatalf("Incorrect value")
	}
}
