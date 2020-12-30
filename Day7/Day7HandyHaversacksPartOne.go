package day7

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func readRules(inputFilePath string) map[string][]string {
	adjMap := make(map[string][]string)

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
			innerBag := strings.Split(bagSplit[1], " bag")[0]
			adjMap[innerBag] = append(adjMap[innerBag], outerBag)
		}
	}
	return adjMap
}

func dfs(adjMap map[string][]string, startingBag string, countedBags map[string]bool) int {
	count := 0
	for _, outerBag := range adjMap[startingBag] {
		_, ok := countedBags[outerBag]
		if ok {
			continue
		}
		count++
		countedBags[outerBag] = true
		count += dfs(adjMap, outerBag, countedBags)
	}
	return count
}

func handyHaversacksPartOne(inputFilePath string, startingBag string) int {
	rules := readRules(inputFilePath)
	countedBags := make(map[string]bool)
	return dfs(rules, startingBag, countedBags)
}
