package day9

import (
	utils "AdventOfCode/Utils"
	"bufio"
	"math"
	"os"
)

func encodingErrorPartTwo(inputFilePath string, preambleLen int) int {
	invalidNumber := encodingErrorPartOne(inputFilePath, preambleLen)

	file, err := os.Open(inputFilePath)
	if err != nil {
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numbers := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		val := readNumber(line)
		numbers = append(numbers, val)
	}
	diff := invalidNumber
	minVal := utils.Pair{A: math.MaxInt32, B: 0}
	maxVal := utils.Pair{A: math.MinInt32, B: 0}
	startingIdx := 0
	idx := 0
	for idx < len(numbers) {
		diff = diff - numbers[idx]
		minVal = utils.MinPair(minVal, utils.Pair{A: numbers[idx], B: idx})
		maxVal = utils.MaxPair(maxVal, utils.Pair{A: numbers[idx], B: idx})
		if diff == 0 && minVal.B != maxVal.B {
			return minVal.A + maxVal.A
		}
		if diff < 0 {
			diff = invalidNumber
			minVal = utils.Pair{A: math.MaxInt32, B: 0}
			maxVal = utils.Pair{A: math.MinInt32, B: 0}
			startingIdx++
			idx = startingIdx
		} else {
			idx++
		}
	}
	return -1
}
