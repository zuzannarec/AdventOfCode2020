package day10

import (
	utils "AdventOfCode/Utils"
	"bufio"
	"log"
	"math"
	"os"
)

func adapterArrayPartOne(inputFilePath string) int {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	adapters := make(map[int]bool)
	maxAdapter := math.MinInt32
	for scanner.Scan() {
		adapter := utils.ReadNumber(scanner.Text())
		adapters[adapter] = true
		maxAdapter = utils.MaxInt(maxAdapter, adapter)
	}
	diffOne := 0
	diffThree := 0
	startingJ := 0
	for startingJ != maxAdapter {
		for i := 1; i <= 3; i++ {
			startingJ++
			if startingJ > maxAdapter {
				log.Fatalf("Cannot use all adapter. No such chain")
			}
			_, ok := adapters[startingJ]
			if ok && i == 1 {
				diffOne++
				break
			}
			if ok && i == 2 {
				break
			}
			if ok && i == 3 {
				diffThree++
				break
			}
		}
	}
	return diffOne * (diffThree + 1)
}
