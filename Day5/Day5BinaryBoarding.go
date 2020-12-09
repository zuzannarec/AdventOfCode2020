package day5

import (
	"bufio"
	"os"
)

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

const rowCount = 128
const colCount = 8

func binaryBoarding(inputFilePath string) int {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return 0
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	maxSeat := 0
	for scanner.Scan() {
		line := scanner.Text()
		maxRowIdx := rowCount
		minRowIdx := 1
		maxColIdx := colCount
		minColIdx := 1
		for idx, char := range line {
			if idx < 7 {
				rowHalf := minRowIdx + ((maxRowIdx - minRowIdx + 1) / 2)
				if string(char) == "F" {
					maxRowIdx = rowHalf
				} else if string(char) == "B" {
					minRowIdx = rowHalf
				}
			} else {
				colHalf := minColIdx + ((maxColIdx - minColIdx + 1) / 2)
				if string(char) == "L" {
					maxColIdx = colHalf
				} else if string(char) == "R" {
					minColIdx = colHalf
				}
			}

		}
		maxSeat = max(maxSeat, (minRowIdx-1)*8+minColIdx-1)
	}
	return maxSeat
}
