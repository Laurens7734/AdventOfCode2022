package main

import (
	"adventofcode2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := utils.ReadFile("../Datafiles/infi.txt")
	Assignment1(input)
	Assignment2(input)
}

func Assignment1(input []string) {
	direction := 0
	x := 0
	y := 0
	for _, line := range input {
		parts := strings.Fields(line)
		num64, _ := strconv.ParseInt(parts[1], 0, 0)
		num := int(num64)
		if parts[0] == "draai" {

			direction += num
			if direction < 0 {
				direction += 360
			}
			direction %= 360
		} else {
			if direction < 46 || direction > 314 {
				y += num
			}
			if direction < 136 && direction > 44 {
				x += num
			}
			if direction < 226 && direction > 134 {
				y -= num
			}
			if direction < 316 && direction > 224 {
				x -= num
			}
		}
	}
	fmt.Println(x + y)
}

func Assignment2(input []string) {
	direction := 0
	x := 0
	y := 0
	grid := make([][]byte, 10)
	for i, _ := range grid {
		grid[i] = make([]byte, 100)
		for j := 0; j < 100; j++ {
			grid[i][j] = '.'
		}
	}
	for _, line := range input {
		parts := strings.Fields(line)
		num64, _ := strconv.ParseInt(parts[1], 0, 64)
		num := int(num64)
		if parts[0] == "draai" {
			direction += num
			if direction < 0 {
				direction += 360
			}
			direction %= 360
		} else if parts[0] == "spring" {
			if direction < 46 || direction > 314 {
				y += num
			}
			if direction < 136 && direction > 44 {
				x += num
			}
			if direction < 226 && direction > 134 {
				y -= num
			}
			if direction < 316 && direction > 224 {
				x -= num
			}
			grid[y][x] = '#'
		} else {
			absnum := utils.Abs(num)
			stepdir := num / absnum
			for i := 0; i < absnum; i++ {
				if direction < 46 || direction > 314 {
					y += stepdir
				}
				if direction < 136 && direction > 44 {
					x += stepdir
				}
				if direction < 226 && direction > 134 {
					y -= stepdir
				}
				if direction < 316 && direction > 224 {
					x -= stepdir
				}
				grid[y][x] = '#'
			}
		}
	}
	for i := len(grid) - 1; i >= 0; i-- {
		fmt.Println(string(grid[i]))
	}
}
