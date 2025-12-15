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
		if isIncrease(line) {
			index += getTurnBy(line)
		} else {
			index -= getTurnBy(line)
		}
		if index < minIndex || index > maxIndex {
			index %= 100
		}
		if index == 0 {
			cnt++
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
