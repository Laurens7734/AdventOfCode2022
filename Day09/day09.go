package main

import (
	"adventofcode2022/utils"
	"fmt"
	"strconv"
)

type knot struct {
	x int
	y int
}

func main() {
	input := utils.ReadFile("../Datafiles/day09.txt")
	Assignment1(input)
	Assignment2(input)
}

func Assignment1(input []string) {
	head := knot{x: 0, y: 0}
	tail := knot{x: 0, y: 0}
	var visited []string
	for _, line := range input {
		amount, _ := strconv.ParseInt(string(line[2:]), 0, 64)
		dir := getDir(line[0])
		for i := 0; i < int(amount); i++ {
			head = dir(head)
			taildir := nextDir(head, tail)
			tail = taildir(tail)
			newcord := makeCordString(tail.x, tail.y)
			if !utils.Contains(visited, newcord) {
				visited = utils.AppendSlice(visited, newcord)
			}
		}
	}
	fmt.Println(len(visited))
}

func Assignment2(input []string) {
	knots := make([]knot, 10)
	fmt.Println(knots)
	var visited []string
	for _, line := range input {
		amount, _ := strconv.ParseInt(string(line[2:]), 0, 64)
		for i := 0; i < int(amount); i++ {
			dir := getDir(line[0])
			for j, knot := range knots {
				if j != 0 {
					dir = nextDir(knots[j-1], knot)
				}
				knots[j] = dir(knot)
			}
			newcord := makeCordString(knots[9].x, knots[9].y)
			if !utils.Contains(visited, newcord) {
				visited = utils.AppendSlice(visited, newcord)
			}
		}
	}
	fmt.Println(len(visited))
}

func nextDir(h, t knot) func(knot) knot {
	dx := h.x - t.x
	dy := h.y - t.y
	if (utils.Abs(dx) + utils.Abs(dy)) > 2 {
		if dx > 0 {
			if dy > 0 {
				return getDir('E')
			} else {
				return getDir('C')
			}
		} else {
			if dy > 0 {
				return getDir('Q')
			} else {
				return getDir('Z')
			}
		}
	} else if utils.Abs(dx) == 2 {
		if dx > 0 {
			return getDir('R')
		} else {
			return getDir('L')
		}
	} else if utils.Abs(dy) == 2 {
		if dy > 0 {
			return getDir('U')
		} else {
			return getDir('D')
		}
	} else {
		return getDir('S')
	}
}

func makeCordString(x, y int) string {
	return fmt.Sprintf("%v-%v", x, y)
}

func getDir(dir byte) func(knot) knot {
	switch dir {
	case 'R':
		return Right
	case 'L':
		return Left
	case 'U':
		return Up
	case 'D':
		return Down
	case 'Q':
		return UpLeft
	case 'E':
		return UpRight
	case 'Z':
		return DownLeft
	case 'C':
		return DownRight
	case 'S':
		return Stationairy
	}
	return nil
}

func Right(k knot) knot {
	k.x += 1
	return k
}

func Left(k knot) knot {
	k.x -= 1
	return k
}

func Up(k knot) knot {
	k.y += 1
	return k
}

func Down(k knot) knot {
	k.y -= 1
	return k
}

func UpLeft(k knot) knot {
	k.x -= 1
	k.y += 1
	return k
}

func UpRight(k knot) knot {
	k.x += 1
	k.y += 1
	return k
}

func DownLeft(k knot) knot {
	k.x -= 1
	k.y -= 1
	return k
}

func DownRight(k knot) knot {
	k.x += 1
	k.y -= 1
	return k
}

func Stationairy(k knot) knot {
	return k
}
