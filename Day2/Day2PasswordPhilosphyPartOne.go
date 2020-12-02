package day2

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseBounds(bounds []string) (int, int) {
	freqDown, err := strconv.Atoi(bounds[0])
	if err != nil {
		log.Fatal(err, bounds[0])
		freqDown = -1
	}
	freqUp, err := strconv.Atoi(bounds[1])
	if err != nil {
		log.Fatal(err, bounds[1])
		freqUp = -1
	}
	return freqDown, freqUp
}

func checkPartOne(pass string, letter string, boundDown int, boundUp int) bool {
	lettersCount := 0
	for _, l := range pass {
		if string(l) == letter {
			lettersCount++
			if lettersCount > boundUp {
				break
			}
		}
	}
	if lettersCount >= boundDown && lettersCount <= boundUp {
		return true
	}
	return false
}

func parseInput(scanner bufio.Scanner) (string, string, int, int) {
	split := strings.Split(scanner.Text(), " ")
	freq := strings.Split(split[0], "-")
	boundDown, boundUp := parseBounds(freq)
	if boundDown == -1 || boundUp == -1 {
		log.Fatal("Incorrect bounds")
	}
	letter := strings.Split(strings.TrimSpace(split[1]), ":")[0]
	pass := split[2]
	return pass, letter, boundDown, boundUp
}

func passwordPhilosphyPartOne(inputFilePath string) int {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		pass, letter, freqDown, freqUp := parseInput(*scanner)
		result := checkPartOne(pass, letter, freqDown, freqUp)
		if result {
			count++
		}
	}
	return count
}
