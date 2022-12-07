package main

import (
	"adventofcode2022/utils"
	"fmt"
	"strconv"
	"strings"
)

type CratePosition struct {
	position int
	char     byte
}

func main() {
	input := utils.ReadFile("../Datafiles/day05.txt")
	Assignment1(input)
	Assignment2(input)
}

func Assignment1(input []string) {
	var stacks []string
	for _, line := range input {
		if strings.Contains(line, "[") {
			positions := getCrates(line)
			for _, crate := range positions {
				for len(stacks) <= crate.position {
					stacks = AppendString(stacks, "")
				}
				stacks[crate.position] = string(crate.char) + stacks[crate.position]
			}
		} else if strings.Contains(line, "move") {
			amount, from, to := getNumbers(line)
			for i := 0; i < amount; i++ {
				stacks[to] += string(stacks[from][len(stacks[from])-1])
				stacks[from] = stacks[from][:len(stacks[from])-1]
			}
		}
	}
	result := ""
	for _, s := range stacks {
		if s != "" {
			result += string(s[len(s)-1])
		}
	}
	fmt.Println(result)
}

func Assignment2(input []string) {
	var stacks []string
	for _, line := range input {
		if strings.Contains(line, "[") {
			positions := getCrates(line)
			for _, crate := range positions {
				for len(stacks) <= crate.position {
					stacks = AppendString(stacks, "")
				}
				stacks[crate.position] = string(crate.char) + stacks[crate.position]
			}
		} else if strings.Contains(line, "move") {
			amount, from, to := getNumbers(line)
			stacks[to] += stacks[from][len(stacks[from])-amount:]
			stacks[from] = stacks[from][:len(stacks[from])-amount]
		}
	}
	result := ""
	for _, s := range stacks {
		if s != "" {
			result += string(s[len(s)-1])
		}
	}
	fmt.Println(result)
}

func getCrates(input string) []CratePosition {
	var result []CratePosition
	parts := strings.Split(input, " ")
	emptycount := 0
	currentposition := 0
	for _, crate := range parts {
		if crate == "" {
			emptycount += 1
			continue
		}
		currentposition += (emptycount / 4)
		emptycount = 0
		newPosition := CratePosition{
			position: currentposition,
			char:     crate[1]}
		result = AppendCratePosition(result, newPosition)
		currentposition += 1
	}
	return result
}

func getNumbers(input string) (int, int, int) {
	parts := strings.Split(input, " ")
	amount, _ := strconv.ParseInt(parts[1], 0, 64)
	from, _ := strconv.ParseInt(parts[3], 0, 64)
	to, _ := strconv.ParseInt(parts[5], 0, 64)
	return int(amount), int(from) - 1, int(to) - 1
}

func AppendCratePosition(slice []CratePosition, data CratePosition) []CratePosition {
	m := len(slice)
	n := m + 1
	if n > cap(slice) {
		newSlice := make([]CratePosition, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	slice[m] = data
	return slice
}

func AppendString(slice []string, data string) []string {
	m := len(slice)
	n := m + 1
	if n > cap(slice) {
		newSlice := make([]string, n)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	slice[m] = data
	return slice
}
