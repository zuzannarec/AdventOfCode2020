package day2

import (
	"bufio"
	"log"
	"os"
)

func checkPartTwo(pass string, letter string, idDown int, idUp int) bool {
	if idDown == -1 || idUp == -1 {
		log.Fatal("Incorrect bounds")
	}
	if string(pass[idDown]) == letter && string(pass[idUp]) != letter {
		return true
	}
	if string(pass[idDown]) != letter && string(pass[idUp]) == letter {
		return true
	}
	return false
}

func passwordPhilosphyPartTwo(inputFilePath string) int {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		pass, letter, idDown, idUp := parseInput(*scanner)
		idDown--
		idUp--
		if checkPartTwo(pass, letter, idDown, idUp) {
			count++
		}
	}
	return count
}
