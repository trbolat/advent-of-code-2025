package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	from, to int
}

func main() {
	ranges := readRanges("02/input.txt")

	sum := 0
	for _, r := range ranges {
		for id := r.from; id <= r.to; id++ {
			idStr := strconv.Itoa(id)
			half := idStr[0:int(math.Ceil(float64(len(idStr))/2))]

			if idStr == half+half {
				sum += id
			}
		}
	}

	fmt.Println(sum)
}

func readRanges(file string) []Range {
	bytes, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	rawRanges := strings.Split(string(bytes), ",")

	out := make([]Range, 0, len(rawRanges))
	for _, rawRange := range rawRanges {
		split := strings.Split(strings.TrimSpace(rawRange), "-")
		from, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		to, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		out = append(
			out, Range{
				from: from,
				to:   to,
			},
		)
	}
	return out
}
