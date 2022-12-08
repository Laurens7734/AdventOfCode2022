package main

import (
	"adventofcode2022/utils"
	"fmt"
	"strconv"
)

func main() {
	input := utils.ReadFile("../Datafiles/day08.txt")
	grid := createGrid(input)
	Assignment1(grid)
	Assignment2(grid)
}

func Assignment1(grid [][]int) {
	visible := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			hight := grid[i][j]
			if checkPaths(grid, i, j, hight) {
				visible++
			}
		}
	}
	fmt.Println(visible)
}

func Assignment2(grid [][]int) {
	bestscore := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			score := calcScore(grid, i, j, grid[i][j])
			if score > bestscore {
				bestscore = score
			}
		}
	}
	fmt.Println(bestscore)
}

func createGrid(input []string) [][]int {
	response := make([][]int, len(input))
	for i := 0; i < len(input); i++ {
		response[i] = make([]int, len(input[0]))
	}
	for i, s := range input {
		for j := 0; j < len(s); j++ {
			val, _ := strconv.ParseInt(string(s[j]), 0, 64)
			response[i][j] += int(val)
		}
	}
	return response
}

func checkPaths(grid [][]int, x, y, hight int) bool {
	blockedcount := 0
	tempX := x - 1
	for tempX >= 0 {
		if grid[tempX][y] >= hight {
			blockedcount++
			break
		}
		tempX--
	}
	tempX = x + 1
	for tempX < len(grid) {
		if grid[tempX][y] >= hight {
			blockedcount++
			break
		}
		tempX++
	}
	tempY := y - 1
	for tempY >= 0 {
		if grid[x][tempY] >= hight {
			blockedcount++
			break
		}
		tempY--
	}
	tempY = y + 1
	for tempY < len(grid[0]) {
		if grid[x][tempY] >= hight {
			blockedcount++
			break
		}
		tempY++
	}
	return blockedcount < 4
}
func calcScore(grid [][]int, x, y, hight int) int {
	score := 1
	tempX := x - 1
	count := 0
	for tempX >= 0 {
		count++
		if grid[tempX][y] >= hight {
			break
		}
		tempX--
	}
	score *= count
	count = 0
	tempX = x + 1
	for tempX < len(grid) {
		count++
		if grid[tempX][y] >= hight {
			break
		}
		tempX++
	}
	score *= count
	count = 0
	tempY := y - 1
	for tempY >= 0 {
		count++
		if grid[x][tempY] >= hight {
			break
		}
		tempY--
	}
	score *= count
	count = 0
	tempY = y + 1
	for tempY < len(grid[0]) {
		count++
		if grid[x][tempY] >= hight {
			break
		}
		tempY++
	}
	score *= count
	return score
}
