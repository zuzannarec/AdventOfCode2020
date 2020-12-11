package day6

import (
	"bufio"
	"os"
	"strings"
)

func customCustomsPartOne(inputFilePath string) int {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	ansMap := make(map[rune]bool)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			ansMap = make(map[rune]bool)
			continue
		}
		for _, char := range line {
			_, ok := ansMap[char]
			if !ok {
				ansMap[char] = true
				count++
			}
		}
	}
	return count
}
