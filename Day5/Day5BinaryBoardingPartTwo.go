package day5

import (
	"bufio"
	"os"
)

const maxSeat = 825

func binaryBoardingPartTwo(inputFilePath string) int {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return 0
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	seatMap := make(map[int]bool)

	for scanner.Scan() {
		line := scanner.Text()
		minRowIdx, minColIdx := calcInterval(line)
		seatID := (minRowIdx-1)*8 + minColIdx - 1
		seatMap[seatID] = true
	}
	for seat := 8; seat <= maxSeat; seat++ {
		_, ok := seatMap[seat]
		if !ok {
			return seat
		}
	}
	return 0
}
