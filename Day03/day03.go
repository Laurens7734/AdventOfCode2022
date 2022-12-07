package main

import (
	"adventofcode2022/utils"
	"fmt"
	"strings"
)

func main() {
	input := utils.ReadFile("../Datafiles/day03.txt")
	Assignment1(input)
	Assignment2(input)
}

func Assignment1(input []string) {
	totalscore := 0
	for _, line := range input {
		compartment1 := line[0 : len(line)/2]
		compartment2 := line[(len(line) / 2):len(line)]
		for _, c := range compartment1 {
			if strings.Contains(compartment2, string(c)) {
				totalscore += score(c)
				break
			}
		}
	}
	fmt.Println(totalscore)
}

func Assignment2(input []string) {
	totalscore := 0
	for i, line := range input {
		if i%3 == 0 {
			for _, c := range line {
				if strings.Contains(input[i+1], string(c)) {
					if strings.Contains(input[i+2], string(c)) {
						totalscore += score(c)
						break
					}
				}
			}
		}
	}
	fmt.Println(totalscore)
}

func score(item rune) int {
	if item >= 'a' {
		return int(item) - int(('a' - 1))
	} else {
		return int(item) - int(('A' - 1)) + 26
	}
}
