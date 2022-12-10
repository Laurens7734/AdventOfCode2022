package main

import (
	"adventofcode2022/utils"
	"fmt"
	"strconv"
)

func main() {
	input := utils.ReadFile("../Datafiles/day10.txt")
	Assignment1(input)
	Assignment2(input)
}

func Assignment1(input []string) {
	cycle := 0
	x := 1
	signalsum := 0
	for _, line := range input {
		if line[:4] == "addx" {
			num, _ := strconv.ParseInt(line[5:], 0, 64)
			cycle++
			signalsum += check(cycle, x)
			cycle++
			signalsum += check(cycle, x)
			x += int(num)
		} else {
			cycle++
			signalsum += check(cycle, x)
		}
	}
	fmt.Println(signalsum)
}

func Assignment2(input []string) {
	cycle := 0
	x := 1
	image := make([]string, 8)
	for _, line := range input {
		if line[:4] == "addx" {
			num, _ := strconv.ParseInt(line[5:], 0, 64)
			image[int(cycle/40)] += draw(cycle, x)
			cycle++
			image[int(cycle/40)] += draw(cycle, x)
			cycle++
			x += int(num)
		} else {
			image[int(cycle/40)] += draw(cycle, x)
			cycle++

		}
	}
	for _, line := range image {
		fmt.Println(line)
	}

}

func check(cycle, x int) int {
	if cycle%40 == 20 {
		return cycle * x
	}
	return 0
}

func draw(cycle, x int) string {
	drawing := cycle % 40
	if drawing >= x-1 && drawing <= x+1 {
		return "#"
	} else {
		return "."
	}
}
