package day8

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkSign(sign string) int {
	if sign == "+" {
		return 1
	}
	return -1
}

func parseLine(line string) (string, int) {
	lineSplit := strings.Split(line, " ")
	valSplit := strings.SplitN(lineSplit[1], "", 2)
	sign := checkSign(valSplit[0])
	val, err := strconv.Atoi(valSplit[1])
	if err != nil {
		log.Fatalf("Incorrect value")
	}
	return lineSplit[0], sign * val
}

func handheldHaltingPartOne(inputFilePath string) int {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return 0
	}
	defer file.Close()

	content, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		return 0
	}
	lines := strings.Split(string(content), "\n")

	acc := 0
	idx := 0
	instructionIdxMap := make([]int, len(lines))
	for idx < len(lines) && instructionIdxMap[idx] == 0 {
		instructionIdxMap[idx] = 1
		instruction, value := parseLine(lines[idx])
		switch instruction {
		case "acc":
			acc += value
			idx++
		case "jmp":
			idx += value
		case "nop":
			idx++
		default:
			log.Fatalf("Incorrent instruction %s", instruction)
		}
	}
	return acc
}
