package main

import (
	"adventofcode2022/filereader"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := filereader.ReadFile("../Datafiles/day04.txt")
	Assignment1(input)
	Assignment2(input)
}

func Assignment1(input []string) {
	count := 0
	for _, line := range input {
		ranges := strings.Split(line, ",")
		min1, max1 := readRange(ranges[0])
		min2, max2 := readRange(ranges[1])
		if min1 >= min2 && max1 <= max2 {
			count++
		} else if min1 <= min2 && max1 >= max2 {
			count++
		}
	}

	fmt.Println(count)
}

func Assignment2(input []string) {
	count := 0
	for _, line := range input {
		ranges := strings.Split(line, ",")
		min1, max1 := readRange(ranges[0])
		min2, max2 := readRange(ranges[1])
		if min1 >= min2 && min1 <= max2 {
			count++
		} else if max1 >= min2 && max1 <= max2 {
			count++
		} else if min1 <= min2 && max1 >= max2 {
			count++
		}
	}

	fmt.Println(count)
}

func readRange(input string) (int, int) {
	parts := strings.Split(input, "-")
	min, _ := strconv.ParseInt(parts[0], 0, 64)
	max, _ := strconv.ParseInt(parts[1], 0, 64)
	return int(min), int(max)
}
