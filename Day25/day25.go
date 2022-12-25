package main

import (
	"adventofcode2022/utils"
	"fmt"
)

func main() {
	input := utils.ReadFile("../Datafiles/day25.txt")
	Assignment1(input)
}

func Assignment1(input []string) {
	total := 0
	for _, line := range input {
		total += readSNAFU(line)
	}
	fmt.Println(writeSNAFU(total))
}

func readSNAFU(snafu string) int {
	value := 0
	pos := 1
	for i := len(snafu) - 1; i >= 0; i-- {
		char := snafu[i]
		if char == '1' {
			value += pos
		}
		if char == '2' {
			value += pos * 2
		}
		if char == '-' {
			value -= pos
		}
		if char == '=' {
			value -= pos * 2
		}
		pos *= 5
	}
	return value
}

func writeSNAFU(number int) string {
	output := ""
	pos := 1
	add1 := false
	for pos < number {
		amount := (number%(pos*5) - (number % pos)) / pos
		if add1 {
			amount++
		}
		add1 = false
		if amount == 0 {
			output = "0" + output
		}
		if amount == 1 {
			output = "1" + output
		}
		if amount == 2 {
			output = "2" + output
		}
		if amount == 3 {
			output = "=" + output
			add1 = true
		}
		if amount == 4 {
			output = "-" + output
			add1 = true
		}
		if amount == 5 {
			output = "0" + output
			add1 = true
		}
		pos *= 5
	}
	return output
}
