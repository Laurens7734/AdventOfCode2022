package main

import (
	"adventofcode2022/filereader"
	"fmt"
	"strings"
)

func main() {
	input := filereader.ReadFile("../Datafiles/day06.txt")
	Assignment1(input[0])
	Assignment2(input[0])
}

func Assignment1(input string) {
	chars := input[:4]
	for position, newchar := range input {
		if strings.Contains(chars, string(newchar)) {
			chars = chars[1:] + string(newchar)
		} else {
			chars = chars[1:] + string(newchar)
			allunique := true
			for i, c := range chars {
				if strings.Contains(chars[i+1:], string(c)) {
					allunique = false
				}
			}
			if allunique {
				fmt.Println(position + 1)
				break
			}

		}
	}
}

func Assignment2(input string) {
	chars := input[:14]
	for position, newchar := range input {
		if strings.Contains(chars, string(newchar)) {
			chars = chars[1:] + string(newchar)
		} else {
			chars = chars[1:] + string(newchar)
			allunique := true
			for i, c := range chars {
				if strings.Contains(chars[i+1:], string(c)) {
					allunique = false
				}
			}
			if allunique {
				fmt.Println(position + 1)
				break
			}

		}
	}
}
