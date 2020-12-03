package day3

import (
	"bufio"
	"os"
	"strings"
)

func tobogganTrajectory(inputFilePath string, rightCount int, downCount int) int {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	count := 0
	idxX := 0
	for true {
		for it := 0; it < downCount; it++ {
			if !scanner.Scan() {
				return count
			}
		}
		line := strings.Split(scanner.Text(), "")
		idxX = idxX + rightCount
		if idxX >= len(line) {
			idxX = idxX - len(line)
		}
		sign := line[idxX]
		if sign == "#" {
			count++
		}
	}
	return count
}
