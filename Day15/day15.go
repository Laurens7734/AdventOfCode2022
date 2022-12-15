package main

import (
	"adventofcode2022/utils"
	"fmt"
	"strconv"
	"strings"
)

type Sensor struct {
	x        int
	y        int
	distance int
}

func main() {
	input := utils.ReadFile("../Datafiles/day15.txt")
	Assignment1(input)
	fmt.Println(Assignment2(input))
}

func Assignment1(input []string) {
	goal := 2000000
	positions := make(map[string]bool)
	beaconlist := make(map[string]bool)
	for _, line := range input {
		parts := strings.Split(line, " ")
		x, _ := strconv.ParseInt(parts[2][2:len(parts[2])-1], 0, 64)
		y, _ := strconv.ParseInt(parts[3][2:len(parts[3])-1], 0, 64)
		sx := int(x)
		sy := int(y)
		x, _ = strconv.ParseInt(parts[8][2:len(parts[8])-1], 0, 64)
		y, _ = strconv.ParseInt(parts[9][2:], 0, 64)
		bx := int(x)
		by := int(y)
		distance := utils.Abs(int(sx-bx)) + utils.Abs(int(sy-by))
		if sy < goal && sy+distance >= goal {
			maxXDistance := distance - (goal - sy)
			startingX := sx - maxXDistance
			lastX := sx + maxXDistance
			for i := startingX; i <= lastX; i++ {
				cordString := fmt.Sprintf("%v,%v", i, goal)
				if !positions[cordString] {
					positions[cordString] = true
				}
			}
		}
		if sy > goal && sy-distance <= goal {
			maxXDistance := distance - (sy - goal)
			startingX := sx - maxXDistance
			lastX := sx + maxXDistance
			for i := startingX; i <= lastX; i++ {
				cordString := fmt.Sprintf("%v,%v", i, goal)
				if !positions[cordString] {
					positions[cordString] = true
				}
			}
		}
		if sy == goal {
			startingX := sx - distance
			lastX := sx + distance
			for i := startingX; i <= lastX; i++ {
				cordString := fmt.Sprintf("%v,%v", i, goal)
				if !positions[cordString] {
					positions[cordString] = true
				}
			}
		}
		if by == goal {
			cord := fmt.Sprintf("%v,%v", bx, goal)
			beaconlist[cord] = true
		}
	}
	beaconcount := 0
	for beacon, _ := range beaconlist {
		if positions[beacon] {
			beaconcount++
		}
	}

	fmt.Println(len(positions) - beaconcount)
}

func Assignment2(input []string) int {
	border := 4000000
	sensors := make([]Sensor, 0)
	calculations := make([]func(int, int) bool, 0)
	for _, line := range input {
		parts := strings.Split(line, " ")
		x, _ := strconv.ParseInt(parts[2][2:len(parts[2])-1], 0, 64)
		y, _ := strconv.ParseInt(parts[3][2:len(parts[3])-1], 0, 64)
		sx := int(x)
		sy := int(y)
		x, _ = strconv.ParseInt(parts[8][2:len(parts[8])-1], 0, 64)
		y, _ = strconv.ParseInt(parts[9][2:], 0, 64)
		bx := int(x)
		by := int(y)
		distance := utils.Abs(sx-bx) + utils.Abs(sy-by)
		s := Sensor{x: sx, y: sy, distance: distance}
		sensors = append(sensors, s)
		calculations = append(calculations, getCalc(s))
	}
	for _, s := range sensors {
		x := s.x - (s.distance + 1)
		y := s.y
		for x < s.x {
			x++
			y++
			if x < 0 || x > border || y < 0 || y > border {
				continue
			}
			failed := false
			for _, calc := range calculations {
				if calc(x, y) {
					failed = true
				}
			}
			if !failed {
				return (x*4000000 + y)
			}
		}
		for y > s.y {
			x++
			y--
			if x < 0 || x > border || y < 0 || y > border {
				continue
			}
			failed := false
			for _, calc := range calculations {
				if calc(x, y) {
					failed = true
				}
			}
			if !failed {
				return (x*4000000 + y)
			}
		}
	}
	return -1
}

func getCalc(s Sensor) func(int, int) bool {
	return func(x, y int) bool {
		return (utils.Abs(s.x-x) + utils.Abs(s.y-y)) <= s.distance
	}
}
