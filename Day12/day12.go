package main

import (
	"adventofcode2022/utils"
	"fmt"
)

type cord struct {
	x int
	y int
}

var start cord
var end cord
var grid [][]int

func main() {
	input := utils.ReadFile("../Datafiles/day12.txt")
	grid = make([][]int, len(input))
	for i, line := range input {
		grid[i] = make([]int, len(line))
		for j, char := range line {
			if char < 'Z' {
				if char == 'E' {
					end = cord{x: j, y: i}
					grid[i][j] = 26
				} else {
					start = cord{x: j, y: i}
					grid[i][j] = 1
				}
			} else {
				grid[i][j] = int(char) - int('a') + 1
			}
		}
	}
	Assignment1(grid)
	Assignment2(grid)
}

func Assignment1(input [][]int) {
	pastFromStart := make([]cord, 0)
	pastFromEnd := make([]cord, 0)
	currentStart := make([]cord, 0)
	currentEnd := make([]cord, 0)
	currentStart = append(currentStart, start)
	currentEnd = append(currentEnd, end)
	steps := 0
	lookingforpath := true

	for lookingforpath {
		pastFromStart = append(pastFromStart, currentStart...)
		pastFromEnd = append(pastFromEnd, currentEnd...)
		currentStart = nextFromStart(currentStart, pastFromStart)
		currentEnd = nextFromEnd(currentEnd, pastFromEnd)

		steps += 2

		for _, c := range currentStart {
			if containsCord(pastFromEnd, c) {
				steps--
				lookingforpath = false
				break
			}
			if containsCord(currentEnd, c) {
				lookingforpath = false
				break
			}

		}
	}
	fmt.Println(steps)
}

func Assignment2(input [][]int) {
	pastFromEnd := make([]cord, 0)
	currentEnd := make([]cord, 0)
	currentEnd = append(currentEnd, end)
	steps := 0
	lookingforpath := true

	for lookingforpath {
		pastFromEnd = append(pastFromEnd, currentEnd...)
		currentEnd = nextFromEnd(currentEnd, pastFromEnd)

		steps++

		for _, c := range currentEnd {
			if grid[c.y][c.x] == 1 {
				lookingforpath = false
				break
			}
		}
	}
	fmt.Println(steps)
}

func nextFromEnd(currentPositions, traveledPositions []cord) []cord {
	newPositions := make([]cord, 0)
	for _, c := range currentPositions {
		hight := grid[c.y][c.x]
		if c.y > 0 && grid[c.y-1][c.x] >= (hight-1) {
			newCord := cord{x: c.x, y: c.y - 1}
			if !containsCord(traveledPositions, newCord) && !containsCord(newPositions, newCord) {
				newPositions = append(newPositions, newCord)
			}
		}
		if c.x > 0 && grid[c.y][c.x-1] >= (hight-1) {
			newCord := cord{x: c.x - 1, y: c.y}
			if !containsCord(traveledPositions, newCord) && !containsCord(newPositions, newCord) {
				newPositions = append(newPositions, newCord)
			}
		}
		if c.y < len(grid)-1 && grid[c.y+1][c.x] >= (hight-1) {
			newCord := cord{x: c.x, y: c.y + 1}
			if !containsCord(traveledPositions, newCord) && !containsCord(newPositions, newCord) {
				newPositions = append(newPositions, newCord)
			}
		}
		if c.x < len(grid[0])-1 && grid[c.y][c.x+1] >= (hight-1) {
			newCord := cord{x: c.x + 1, y: c.y}
			if !containsCord(traveledPositions, newCord) && !containsCord(newPositions, newCord) {
				newPositions = append(newPositions, newCord)
			}
		}
	}
	return newPositions
}

func nextFromStart(currentPositions, traveledPositions []cord) []cord {
	newPositions := make([]cord, 0)
	for _, c := range currentPositions {
		hight := grid[c.y][c.x]
		if c.y > 0 && grid[c.y-1][c.x] <= (hight+1) {
			newCord := cord{x: c.x, y: c.y - 1}
			if !containsCord(traveledPositions, newCord) && !containsCord(newPositions, newCord) {
				newPositions = append(newPositions, newCord)
			}
		}
		if c.x > 0 && grid[c.y][c.x-1] <= (hight+1) {
			newCord := cord{x: c.x - 1, y: c.y}
			if !containsCord(traveledPositions, newCord) && !containsCord(newPositions, newCord) {
				newPositions = append(newPositions, newCord)
			}
		}
		if c.y < len(grid)-1 && grid[c.y+1][c.x] <= (hight+1) {
			newCord := cord{x: c.x, y: c.y + 1}
			if !containsCord(traveledPositions, newCord) && !containsCord(newPositions, newCord) {
				newPositions = append(newPositions, newCord)
			}
		}
		if c.x < len(grid[0])-1 && grid[c.y][c.x+1] <= (hight+1) {
			newCord := cord{x: c.x + 1, y: c.y}
			if !containsCord(traveledPositions, newCord) && !containsCord(newPositions, newCord) {
				newPositions = append(newPositions, newCord)
			}
		}
	}
	return newPositions
}

func containsCord(list []cord, search cord) bool {
	for _, c := range list {
		if c.x == search.x && c.y == search.y {
			return true
		}
	}
	return false
}
