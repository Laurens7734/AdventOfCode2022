package main

import (
	"adventofcode2022/filereader"
	"fmt"
	"strings"
)

func main() {
	input := filereader.ReadFile("../Datafiles/day02.txt")
	Assignment1(input)
	Assignment2(input)
}

func Assignment1(input []string) {
	score := 0
	for _, s := range input {
		parts := strings.Fields(s)
		if parts[1] == "X" {
			score += 1
			if parts[0] == "C" {
				score += 6
			}
			if parts[0] == "A" {
				score += 3
			}
		}
		if parts[1] == "Y" {
			score += 2
			if parts[0] == "A" {
				score += 6
			}
			if parts[0] == "B" {
				score += 3
			}
		}
		if parts[1] == "Z" {
			score += 3
			if parts[0] == "B" {
				score += 6
			}
			if parts[0] == "C" {
				score += 3
			}
		}
	}
	fmt.Println(score)
}

func Assignment2(input []string) {
	score := 0
	for _, s := range input {
		parts := strings.Fields(s)
		if parts[1] == "X" {
			switch parts[0] {
			case "A":
				score += 3
			case "B":
				score += 1
			case "C":
				score += 2
			}
		}
		if parts[1] == "Y" {
			score += 3
			switch parts[0] {
			case "A":
				score += 1
			case "B":
				score += 2
			case "C":
				score += 3
			}
		}
		if parts[1] == "Z" {
			score += 6
			switch parts[0] {
			case "A":
				score += 2
			case "B":
				score += 3
			case "C":
				score += 1
			}
		}
	}
	fmt.Println(score)
}
