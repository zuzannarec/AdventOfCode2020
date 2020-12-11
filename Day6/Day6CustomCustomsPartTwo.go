package day6

import (
	"bufio"
	"os"
	"strings"
)

func customCustomsPartTwo(inputFilePath string) int {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	personCount := 0
	ansMap := make(map[rune]int)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			for _, el := range ansMap {
				if el == personCount {
					count++
				}
			}
			personCount = 0
			ansMap = make(map[rune]int)
			continue
		}
		for _, char := range line {
			_, ok := ansMap[char]
			if ok {
				ansMap[char]++
			} else {
				ansMap[char] = 1
			}
		}
		personCount++
	}
	for _, el := range ansMap {
		if el == personCount {
			count++
		}
	}
	return count
}
