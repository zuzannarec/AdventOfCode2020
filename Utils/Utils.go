package utils

import (
	"log"
	"strconv"
)

func MaxInt(one int, two int) int {
	if one >= two {
		return one
	}
	return two
}

func MinInt(one int, two int) int {
	if one < two {
		return one
	}
	return two
}

type Pair struct {
	A, B int
}

func MaxPair(one Pair, two Pair) Pair {
	if one.A >= two.A {
		return one
	}
	return two
}

func MinPair(one Pair, two Pair) Pair {
	if one.A < two.A {
		return one
	}
	return two
}

func ReadNumber(line string) int {
	val, err := strconv.Atoi(line)
	if err != nil {
		log.Fatalf("Incorrect input %s", line)
	}
	return val
}
