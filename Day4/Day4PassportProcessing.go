package day3

import (
	"bufio"
	"os"
	"strings"
)

func passportProcessing(inputFilePath string) int {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	requiredFieldsMap := make(map[string]int)
	for idx, field := range requiredFields {
		requiredFieldsMap[field] = idx
	}
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
		current := ""
		for _, char := range line {
			if string(char) == " " {
				current = ""
				continue
			}
			if string(char) == ":" {
				elem, ok := requiredFieldsMap[current]
				if ok {
					foundFields = foundFields | (1 << elem)
				}
				current = ""
				continue
			}
			current += string(char)
		}
	}
	if foundFields == 127 {
		valid++
	}
	return valid
}
