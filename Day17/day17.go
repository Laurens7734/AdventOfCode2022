package main

import (
	"adventofcode2022/utils"
	"fmt"
)

func main() {
	input := utils.ReadFile("../Datafiles/day17.txt")
	Assignment1(input)
	Assignment2(input)
}

var moveleft = [5]func(int, int, [][]bool) int{lefthline, leftcross, leftcorner, leftvline, leftcube}
var moveright = [5]func(int, int, [][]bool) int{righthline, rightcross, rightcorner, rightvline, rightcube}
var movedown = [5]func(int, int, [][]bool) int{downhline, downcross, downcorner, downvline, downcube}
var updateGrid = [5]func(int, int, [][]bool) [][]bool{updatehline, updatecross, updatecorner, updatevline, updatecube}

func Assignment1(input []string) {
	orders := input[0]
	grid := make([][]bool, 0)
	grid = append(grid, make([]bool, 7))
	maxhight := 0
	movecounter := 0
	for i := 0; i < 2022; i++ {
		grid = expandGrid(grid, maxhight)
		x := 2
		y := maxhight + 4
		falling := true
		for falling {
			if orders[movecounter] == '<' {
				x = moveleft[i%len(moveleft)](x, y, grid)
			} else {
				x = moveright[i%len(moveright)](x, y, grid)
			}
			newy := movedown[i%len(movedown)](x, y, grid)
			if newy == y {
				falling = false
				grid = updateGrid[i%len(updateGrid)](x, y, grid)
			} else {
				y = newy
			}
			movecounter++
			movecounter = movecounter % len(orders)
		}
		for i := len(grid) - 1; i > maxhight; i-- {
			for _, b := range grid[i] {
				if b {
					maxhight = i
				}
			}
		}
	}
	fmt.Println(maxhight)
}

func Assignment2(input []string) {
	orders := input[0]
	grid := make([][]bool, 0)
	grid = append(grid, make([]bool, 7))
	maxhight := 0
	movecounter := 0
	firstencounter := true
	storedMovecounter := -1
	totalmoves := 1000000000000
	usedstartcounts := make(map[int][2]int)
	skipedhight := -1
	for i := 0; i < totalmoves; i++ {
		if i%len(movedown) == 0 && skipedhight == -1 {
			if usedstartcounts[movecounter][0] != 0 {
				if firstencounter {
					firstencounter = false
					storedMovecounter = movecounter
					usedstartcounts[movecounter] = [2]int{i, maxhight}
				} else if movecounter == storedMovecounter {
					deltai := i - usedstartcounts[movecounter][0]
					deltahight := maxhight - usedstartcounts[movecounter][1]
					times := int((totalmoves - usedstartcounts[movecounter][0]) / deltai)
					rem := (totalmoves - usedstartcounts[movecounter][0]) % deltai
					i = totalmoves - (rem + 1)
					skipedhight = deltahight * (times - 1)
					continue
				}
			} else {
				usedstartcounts[movecounter] = [2]int{i, maxhight}
			}
		}
		grid = expandGrid(grid, maxhight)
		x := 2
		y := maxhight + 4
		falling := true
		for falling {
			if orders[movecounter] == '<' {
				x = moveleft[i%len(moveleft)](x, y, grid)
			} else {
				x = moveright[i%len(moveright)](x, y, grid)
			}
			newy := movedown[i%len(movedown)](x, y, grid)
			if newy == y {
				falling = false
				grid = updateGrid[i%len(updateGrid)](x, y, grid)
			} else {
				y = newy
			}
			movecounter++
			movecounter = movecounter % len(orders)
		}
		for i := len(grid) - 1; i > maxhight; i-- {
			for _, b := range grid[i] {
				if b {
					maxhight = i
				}
			}
		}
	}
	fmt.Println(maxhight + skipedhight)
}

func expandGrid(grid [][]bool, max int) [][]bool {
	for len(grid) < max+8 {
		line := make([]bool, len(grid[0]))
		grid = append(grid, line)
	}
	return grid
}

func lefthline(x, y int, grid [][]bool) int {
	if x == 0 {
		return x
	}
	if grid[y][x-1] {
		return x
	}
	return x - 1
}
func leftvline(x, y int, grid [][]bool) int {
	if x == 0 {
		return x
	}
	if grid[y][x-1] || grid[y+1][x-1] || grid[y+2][x-1] || grid[y+3][x-1] {
		return x
	}
	return x - 1
}
func leftcorner(x, y int, grid [][]bool) int {
	if x == 0 {
		return x
	}
	if grid[y][x-1] || grid[y+1][x+1] || grid[y+2][x+1] {
		return x
	}
	return x - 1
}
func leftcross(x, y int, grid [][]bool) int {
	if x == 0 {
		return x
	}
	if grid[y][x] || grid[y+1][x-1] || grid[y+2][x] {
		return x
	}
	return x - 1
}
func leftcube(x, y int, grid [][]bool) int {
	if x == 0 {
		return x
	}
	if grid[y][x-1] || grid[y+1][x-1] {
		return x
	}
	return x - 1
}
func righthline(x, y int, grid [][]bool) int {
	if x == len(grid[0])-4 {
		return x
	}
	if grid[y][x+4] {
		return x
	}
	return x + 1
}
func rightvline(x, y int, grid [][]bool) int {
	if x == len(grid[0])-1 {
		return x
	}
	if grid[y][x+1] || grid[y+1][x+1] || grid[y+2][x+1] || grid[y+3][x+1] {
		return x
	}
	return x + 1
}
func rightcorner(x, y int, grid [][]bool) int {
	if x == len(grid[0])-3 {
		return x
	}
	if grid[y][x+3] || grid[y+1][x+3] || grid[y+2][x+3] {
		return x
	}
	return x + 1
}
func rightcross(x, y int, grid [][]bool) int {
	if x == len(grid[0])-3 {
		return x
	}
	if grid[y][x+2] || grid[y+1][x+3] || grid[y+2][x+2] {
		return x
	}
	return x + 1
}
func rightcube(x, y int, grid [][]bool) int {
	if x == len(grid[0])-2 {
		return x
	}
	if grid[y][x+2] || grid[y+1][x+2] {
		return x
	}
	return x + 1
}
func downhline(x, y int, grid [][]bool) int {
	if y == 1 {
		return y
	}
	if grid[y-1][x] || grid[y-1][x+1] || grid[y-1][x+2] || grid[y-1][x+3] {
		return y
	}
	return y - 1
}
func downvline(x, y int, grid [][]bool) int {
	if y == 1 {
		return y
	}
	if grid[y-1][x] {
		return y
	}
	return y - 1
}
func downcorner(x, y int, grid [][]bool) int {
	if y == 1 {
		return y
	}
	if grid[y-1][x] || grid[y-1][x+1] || grid[y-1][x+2] {
		return y
	}
	return y - 1
}
func downcross(x, y int, grid [][]bool) int {
	if y == 1 {
		return y
	}
	if grid[y][x] || grid[y-1][x+1] || grid[y][x+2] {
		return y
	}
	return y - 1
}
func downcube(x, y int, grid [][]bool) int {
	if y == 1 {
		return y
	}
	if grid[y-1][x] || grid[y-1][x+1] {
		return y
	}
	return y - 1
}
func updatehline(x, y int, grid [][]bool) [][]bool {
	grid[y][x] = true
	grid[y][x+1] = true
	grid[y][x+2] = true
	grid[y][x+3] = true
	return grid
}
func updatevline(x, y int, grid [][]bool) [][]bool {
	grid[y][x] = true
	grid[y+1][x] = true
	grid[y+2][x] = true
	grid[y+3][x] = true
	return grid
}
func updatecorner(x, y int, grid [][]bool) [][]bool {
	grid[y][x] = true
	grid[y][x+1] = true
	grid[y][x+2] = true
	grid[y+1][x+2] = true
	grid[y+2][x+2] = true
	return grid
}
func updatecross(x, y int, grid [][]bool) [][]bool {
	grid[y+1][x] = true
	grid[y][x+1] = true
	grid[y+1][x+1] = true
	grid[y+2][x+1] = true
	grid[y+1][x+2] = true
	return grid
}
func updatecube(x, y int, grid [][]bool) [][]bool {
	grid[y][x] = true
	grid[y+1][x] = true
	grid[y][x+1] = true
	grid[y+1][x+1] = true
	return grid
}
