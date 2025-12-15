package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	startIndex = 50
	minIndex   = 0
	maxIndex   = 99
)

func main() {
	lines := getLines("01/input.txt")
	cnt := 0
	index := startIndex

	for _, line := range lines {
		if line == "" {
			continue
		}
		turnBy := getTurnBy(line)
		for i := 0; i < turnBy; i++ {
			if isIncrease(line) {
				index++
			} else {
				index--
			}
			if index%100 == 0 {
				cnt++
			}
		}

		if index < minIndex || index > maxIndex {
			index %= 100
		}
	}
	fmt.Println(cnt)
}

func getTurnBy(line string) int {
	turnByStr := line[1:]
	turnBy, err := strconv.Atoi(turnByStr)
	if err != nil {
		panic(err)
	}
	return turnBy
}

func isIncrease(line string) bool {
	return strings.HasPrefix(line, "R")
}

func getLines(file string) []string {
	bytes, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(bytes), "\n")
}
