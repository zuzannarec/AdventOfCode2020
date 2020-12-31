package day9

import (
	utils "AdventOfCode/Utils"
	"bufio"
	"os"
)

func encodingErrorPartOne(inputFilePath string, preambleLen int) int {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	preamble := make([]int, 0)
	preambleSet := make(map[int]bool)

	for i := 0; i < preambleLen; i++ {
		scanner.Scan()
		line := scanner.Text()
		val := utils.ReadNumber(line)
		preamble = append(preamble, val)
		preambleSet[val] = true
	}
	for scanner.Scan() {
		correct := false
		line := scanner.Text()
		val := utils.ReadNumber(line)
		for _, el := range preamble {
			diff := val - el
			if diff == el {
				continue
			}
			_, ok := preambleSet[diff]
			if ok {
				preamble = append(preamble, val)
				preambleSet[val] = true
				delete(preambleSet, preamble[0])
				preamble = preamble[1:]
				correct = true
				break
			}
		}
		if !correct {
			return val
		}
	}
	return -1
}
