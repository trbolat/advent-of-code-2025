package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	banks := readBanks("03/input.txt")

	sum := 0
	for _, bank := range banks {
		bankMax := 0
		for i := 0; i < len(bank)-1; i++ {
			for j := i + 1; j < len(bank); j++ {
				tmpStr := fmt.Sprintf("%s%s", bank[i:i+1], bank[j:j+1])
				tmp, err := strconv.Atoi(tmpStr)
				if err != nil {
					panic(err)
				}
				if tmp > bankMax {
					bankMax = tmp
				}
			}
		}
		fmt.Printf("found max %d for bank %s\n", bankMax, bank)
		sum += bankMax
	}
	fmt.Println(sum)
}

func readBanks(fileName string) []string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(file), "\n")
}
