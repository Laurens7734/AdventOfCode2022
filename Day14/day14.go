package main

import (
	"adventofcode2022/utils"
	"fmt"
	"strconv"
	"strings"
)

var lowestpoint = -1

func main() {
	input := utils.ReadFile("../Datafiles/day14.txt")

	rocks := make(map[string]bool)
	for _, line := range input {
		parts := strings.Split(line, " -> ")
		lastX := -1
		lastY := -1
		for _, cord := range parts {
			if lastX == -1 {
				rocks[cord] = true
			}
			x, _ := strconv.ParseInt(strings.Split(cord, ",")[0], 0, 64)
			y, _ := strconv.ParseInt(strings.Split(cord, ",")[1], 0, 64)
			xcord := int(x)
			ycord := int(y)
			if lastX != -1 && lastX != xcord {
				for i := 1; i <= utils.Abs(lastX-xcord); i++ {
					if lastX-xcord > 0 {
						cordstring := fmt.Sprintf("%v,%v", lastX-i, ycord)
						rocks[cordstring] = true
					}
					if lastX-xcord < 0 {
						cordstring := fmt.Sprintf("%v,%v", lastX+i, ycord)
						rocks[cordstring] = true
					}
				}
			}
			if lastY != -1 && lastY != ycord {
				for i := 1; i <= utils.Abs(lastY-ycord); i++ {
					if lastY-ycord > 0 {
						cordstring := fmt.Sprintf("%v,%v", xcord, lastY-i)
						rocks[cordstring] = true
					}
					if lastY-ycord < 0 {
						cordstring := fmt.Sprintf("%v,%v", xcord, lastY+i)
						rocks[cordstring] = true
					}
				}
			}
			lastX = xcord
			lastY = ycord
			if ycord > lowestpoint {
				lowestpoint = ycord
			}
		}
	}

	Assignment1(rocks)
	Assignment2(rocks)
}

func Assignment1(input map[string]bool) {
	sand := make(map[string]bool)
	nextCord := getNextCord(input, sand)
	for nextCord != "done" {
		sand[nextCord] = true
		nextCord = getNextCord(input, sand)
	}
	fmt.Println(len(sand))
}

func Assignment2(input map[string]bool) {
	sand := make(map[string]bool)
	nextCord := getNextCord2(input, sand)
	for nextCord != "done" {
		sand[nextCord] = true
		nextCord = getNextCord2(input, sand)
	}
	fmt.Println(len(sand))
}

func getNextCord(rocks, sand map[string]bool) string {
	x := 500
	y := 0
	for y < lowestpoint {
		corddown := fmt.Sprintf("%v,%v", x, y+1)
		cordleft := fmt.Sprintf("%v,%v", x-1, y+1)
		cordright := fmt.Sprintf("%v,%v", x+1, y+1)
		if !(rocks[corddown] || sand[corddown]) {
			y++
		} else if !(rocks[cordleft] || sand[cordleft]) {
			y++
			x--
		} else if !(rocks[cordright] || sand[cordright]) {
			y++
			x++
		} else {
			return fmt.Sprintf("%v,%v", x, y)
		}
	}
	return "done"
}

func getNextCord2(rocks, sand map[string]bool) string {
	if sand["500,0"] {
		return "done"
	}
	x := 500
	y := 0
	for true {
		if y == lowestpoint+1 {
			return fmt.Sprintf("%v,%v", x, y)
		}
		corddown := fmt.Sprintf("%v,%v", x, y+1)
		cordleft := fmt.Sprintf("%v,%v", x-1, y+1)
		cordright := fmt.Sprintf("%v,%v", x+1, y+1)
		if !(rocks[corddown] || sand[corddown]) {
			y++
		} else if !(rocks[cordleft] || sand[cordleft]) {
			y++
			x--
		} else if !(rocks[cordright] || sand[cordright]) {
			y++
			x++
		} else {
			return fmt.Sprintf("%v,%v", x, y)
		}
	}
	return "error"
}
