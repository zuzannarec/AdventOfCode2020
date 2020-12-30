package day7

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type pair struct {
	name  string
	count int
}

func readRulesPartTwo(inputFilePath string) map[string][]pair {
	adjMap := make(map[string][]pair)

	file, err := os.Open(inputFilePath)
	if err != nil {
		return adjMap
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(` bag(s\s|\s)contain(s\s|\s)`)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := re.Split(line, 2)
		outerBag := lineSplit[0]
		restBags := strings.Split(lineSplit[1], ", ")
		for _, bag := range restBags {
			bagSplit := strings.SplitN(bag, " ", 2)
			if bagSplit[0] == "no" {
				continue
			}
			count, err := strconv.Atoi(bagSplit[0])
			if err != nil {
				log.Fatalf("Cannot convert string to int")
			}
			innerBag := strings.Split(bagSplit[1], " bag")[0]
			adjMap[outerBag] = append(adjMap[outerBag], pair{innerBag, count})
		}
	}
	return adjMap
}

func dfsPartTwo(adjMap map[string][]pair, startingBag string, multi int) int {
	count := 0
	for _, innerBagpair := range adjMap[startingBag] {
		count += multi * innerBagpair.count
		count += dfsPartTwo(adjMap, innerBagpair.name, innerBagpair.count*multi)
	}
	return count
}

func handyHaversacksPartTwo(inputFilePath string, startingBag string) int {
	rules := readRulesPartTwo(inputFilePath)
	return dfsPartTwo(rules, startingBag, 1)
}
