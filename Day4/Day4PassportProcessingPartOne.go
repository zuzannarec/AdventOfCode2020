package day4

import (
	"bufio"
	"os"
	"strings"
)

func passportProcessingPartOne(inputFilePath string) int {
	requiredFieldsMap := map[string]int{"byr": 0, "iyr": 1, "eyr": 2, "hgt": 3, "hcl": 4, "ecl": 5, "pid": 6}

	file, err := os.Open(inputFilePath)
	if err != nil {
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	valid := 0
	foundFields := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			if foundFields == 127 {
				valid++
			}
			foundFields = 0
			continue
		}
		lineSplit := strings.Split(line, " ")
		for _, field := range lineSplit {
			fieldSplit := strings.Split(field, ":")
			elem, ok := requiredFieldsMap[fieldSplit[0]]
			if ok {
				foundFields = foundFields | (1 << elem)
			}
		}
	}
	if foundFields == 127 {
		valid++
	}
	return valid
}
