package day4

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func evalByr(val string) bool {
	v, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	return v >= 1920 && v <= 2002
}

func evalIyr(val string) bool {
	v, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	return v >= 2010 && v <= 2020
}

func evalEyr(val string) bool {
	v, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	return v >= 2020 && v <= 2030
}

func evalHgt(val string) bool {
	unitCmStart := strings.Index(val, "cm")
	unitInStart := strings.Index(val, "in")
	if (unitCmStart != -1 && unitInStart != -1) ||
		(unitCmStart == -1 && unitInStart == -1) {
		return false
	}
	height := ""
	for _, char := range val {
		_, err := strconv.Atoi(string(char))
		if err != nil {
			break
		}
		height += string(char)
	}
	v, err := strconv.Atoi(height)
	if err != nil {
		return false
	}
	if unitCmStart != -1 {
		return v >= 150 && v <= 193
	}
	return v >= 59 && v <= 76
}

func evalHcl(val string) bool {
	stringSplit := strings.SplitAfter(val, "#")
	if cap(stringSplit) != 2 || stringSplit[0] != "#" {
		return false
	}
	re := regexp.MustCompile(`[a-z0-9]*`)
	if stringSplit[1] == re.FindString(stringSplit[1]) {
		return true
	}
	return false
}

func evalEcl(val string) bool {
	eyeColors := map[string]bool{
		"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}
	_, ok := eyeColors[val]
	return ok
}

func evalPid(val string) bool {
	a := len(val)
	if a != 9 {
		return false
	}
	re := regexp.MustCompile(`[0-9]*`)
	if val == re.FindString(val) {
		return true
	}
	return false
}

func evalSwitch(field string, val string) bool {
	switch field {
	case "byr":
		return evalByr(val)
	case "iyr":
		return evalIyr(val)
	case "eyr":
		return evalEyr(val)
	case "hgt":
		return evalHgt(val)
	case "hcl":
		return evalHcl(val)
	case "ecl":
		return evalEcl(val)
	case "pid":
		return evalPid(val)
	default:
		return false
	}
}

func passportProcessingPartTwo(inputFilePath string) int {
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
				eval := evalSwitch(fieldSplit[0], fieldSplit[1])
				if eval {
					foundFields = foundFields | (1 << elem)
				}
			}
		}
	}
	if foundFields == 127 {
		valid++
	}
	return valid
}
