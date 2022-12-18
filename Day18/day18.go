package main

import (
	"adventofcode2022/utils"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

func main() {
	input := utils.ReadFile("../Datafiles/day18.txt")
	Assignment1(input)
	Assignment2(input)
}

func Assignment1(input []string) {
	faceCount := 0
	space := make(map[string]bool)
	for _, line := range input {
		parts := strings.Split(line, ",")
		cord, _ := strconv.ParseInt(parts[0], 0, 64)
		x := int(cord)
		cord, _ = strconv.ParseInt(parts[1], 0, 64)
		y := int(cord)
		cord, _ = strconv.ParseInt(parts[2], 0, 64)
		z := int(cord)
		space[line] = true
		faceCount += 6
		if space[fmt.Sprintf("%v,%v,%v", x+1, y, z)] {
			faceCount -= 2
		}
		if space[fmt.Sprintf("%v,%v,%v", x-1, y, z)] {
			faceCount -= 2
		}
		if space[fmt.Sprintf("%v,%v,%v", x, y+1, z)] {
			faceCount -= 2
		}
		if space[fmt.Sprintf("%v,%v,%v", x, y-1, z)] {
			faceCount -= 2
		}
		if space[fmt.Sprintf("%v,%v,%v", x, y, z+1)] {
			faceCount -= 2
		}
		if space[fmt.Sprintf("%v,%v,%v", x, y, z-1)] {
			faceCount -= 2
		}
	}
	fmt.Println(faceCount)
}

func Assignment2(input []string) {
	faceCount := 0
	space := make(map[string]bool)
	posInternal := make(map[string]int)
	bigx, bigy, bigz := 0, 0, 0
	for _, line := range input {
		parts := strings.Split(line, ",")
		cord, _ := strconv.ParseInt(parts[0], 0, 64)
		x := int(cord)
		if x > bigx {
			bigx = x
		}
		cord, _ = strconv.ParseInt(parts[1], 0, 64)
		y := int(cord)
		if y > bigy {
			bigy = y
		}
		cord, _ = strconv.ParseInt(parts[2], 0, 64)
		z := int(cord)
		if z > bigz {
			bigz = z
		}

		space[line] = true
		faceCount += 6
		if posInternal[line] != 0 {
			delete(posInternal, line)
		}

		cordstr := fmt.Sprintf("%v,%v,%v", x+1, y, z)
		if space[cordstr] {
			faceCount -= 2
		} else {
			posInternal[cordstr]++
		}
		cordstr = fmt.Sprintf("%v,%v,%v", x-1, y, z)
		if space[cordstr] {
			faceCount -= 2
		} else {
			posInternal[cordstr]++
		}
		cordstr = fmt.Sprintf("%v,%v,%v", x, y+1, z)
		if space[cordstr] {
			faceCount -= 2
		} else {
			posInternal[cordstr]++
		}
		cordstr = fmt.Sprintf("%v,%v,%v", x, y-1, z)
		if space[cordstr] {
			faceCount -= 2
		} else {
			posInternal[cordstr]++
		}
		cordstr = fmt.Sprintf("%v,%v,%v", x, y, z+1)
		if space[cordstr] {
			faceCount -= 2
		} else {
			posInternal[cordstr]++
		}
		cordstr = fmt.Sprintf("%v,%v,%v", x, y, z-1)
		if space[cordstr] {
			faceCount -= 2
		} else {
			posInternal[cordstr]++
		}
	}

	knownInternal := make(map[string]bool)
	knownExtrenal := make(map[string]bool)
	for cord, amount := range posInternal {
		parts := strings.Split(cord, ",")
		c, _ := strconv.ParseInt(parts[0], 0, 64)
		x := int(c)
		c, _ = strconv.ParseInt(parts[1], 0, 64)
		y := int(c)
		c, _ = strconv.ParseInt(parts[2], 0, 64)
		z := int(c)
		if knownExtrenal[cord] {
			continue
		} else if knownInternal[cord] {
			faceCount -= amount
		} else {
			res, visited := isInternal(space, x, y, z, bigx, bigy, bigz, make(map[string]bool))
			if res {
				maps.Copy(knownInternal, visited)
				faceCount -= amount
			} else {
				maps.Copy(knownExtrenal, visited)
			}
		}
	}

	fmt.Println(faceCount)
}

func isInternal(space map[string]bool, x, y, z, maxx, maxy, maxz int, visited map[string]bool) (bool, map[string]bool) {
	newVisited := visited
	newVisited[fmt.Sprintf("%v,%v,%v", x, y, z)] = true
	if x >= maxx || x <= 0 || y >= maxy || y <= 0 || z >= maxz || z <= 0 {
		return false, newVisited
	}
	cord := fmt.Sprintf("%v,%v,%v", x+1, y, z)
	if !space[cord] && !newVisited[cord] {
		res, newVisited := isInternal(space, x+1, y, z, maxx, maxy, maxz, newVisited)
		if !res {
			return false, newVisited
		}
	}
	cord = fmt.Sprintf("%v,%v,%v", x-1, y, z)
	if !space[cord] && !newVisited[cord] {
		res, newVisited := isInternal(space, x-1, y, z, maxx, maxy, maxz, newVisited)
		if !res {
			return false, newVisited
		}
	}
	cord = fmt.Sprintf("%v,%v,%v", x, y+1, z)
	if !space[cord] && !newVisited[cord] {
		res, newVisited := isInternal(space, x, y+1, z, maxx, maxy, maxz, newVisited)
		if !res {
			return false, newVisited
		}
	}
	cord = fmt.Sprintf("%v,%v,%v", x, y-1, z)
	if !space[cord] && !newVisited[cord] {
		res, newVisited := isInternal(space, x, y-1, z, maxx, maxy, maxz, newVisited)
		if !res {
			return false, newVisited
		}
	}
	cord = fmt.Sprintf("%v,%v,%v", x, y, z+1)
	if !space[cord] && !newVisited[cord] {
		res, newVisited := isInternal(space, x, y, z+1, maxx, maxy, maxz, newVisited)
		if !res {
			return false, newVisited
		}
	}
	cord = fmt.Sprintf("%v,%v,%v", x, y, z-1)
	if !space[cord] && !newVisited[cord] {
		res, newVisited := isInternal(space, x, y, z-1, maxx, maxy, maxz, newVisited)
		if !res {
			return false, newVisited
		}
	}
	return true, newVisited
}
