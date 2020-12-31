package day8

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func handheldHaltingPartTwo(inputFilePath string) int {
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
	lastIdx := 0
	lastAcc := 0
	changed := false
	instructionIdxMap := make([]int, len(lines))
	for {
		if idx == len(lines)-1 {
			instruction, value := parseLine(lines[idx])
			switch instruction {
			case "acc":
				acc += value
			case "jmp":
				if value <= 0 && changed {
					log.Fatalf("Incorrent input. Cannot complete")
				}
			case "nop":
			default:
				log.Fatalf("Incorrent instruction %s", instruction)
			}
			return acc
		}
		if idx >= len(lines) {
			log.Fatalf("Incorrent input. Cannot complete")
		}
		if instructionIdxMap[idx] == 1 {
			acc = lastAcc
			idx = lastIdx
			changed = false
			instructionIdxMap = make([]int, len(lines))
			continue
		}
		instructionIdxMap[idx] = 1
		instruction, value := parseLine(lines[idx])
		switch instruction {
		case "acc":
			acc += value
			idx++
		case "jmp":
			if changed {
				idx += value
			} else {
				lastIdx = idx + value
				lastAcc = acc
				idx++
				changed = true
			}
		case "nop":
			if changed {
				idx++
			} else {
				lastIdx = idx + 1
				lastAcc = acc
				idx += value
				changed = true
			}
		default:
			log.Fatalf("Incorrent instruction %s", instruction)
		}
	}
}
