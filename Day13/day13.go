package main

import (
	"adventofcode2022/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := utils.ReadFile("../Datafiles/day13.txt")

	Assignment1(input)
	Assignment2(input)
}

func Assignment1(input []string) {
	score := 0
	for i, line := range input {
		if i%3 != 0 {
			continue
		}
		result := compare(line, input[i+1])
		if result == 0 {
			fmt.Println("error")
		}
		if result == 1 {
			score += (3 + i) / 3
		}
	}
	fmt.Println(score)
}

func Assignment2(input []string) {
	key1 := "[[2]]"
	key2 := "[[6]]"

	above1 := 1
	above2 := 2

	for _, line := range input {
		if line == "" {
			continue
		}
		if compare(line, key1) == 1 {
			above1++
		}
		if compare(line, key2) == 1 {
			above2++
		}
	}
	fmt.Println(above1 * above2)
}

func compare(left, right string) int {
	if left == "" && right == "" {
		return 0
	}
	if left == "" {
		return 1
	}
	if right == "" {
		return -1
	}
	if isList(left) && isList(right) {
		return compare(left[1:len(left)-1], right[1:len(right)-1])
	} else if isList(left) || isList(right) {
		if isList(left) {
			r := splitlist(right)
			newr := fmt.Sprintf("[%v]", r[0])
			return compare(left, newr)
		} else {
			l := splitlist(left)
			var newl string
			if !isList(l[0]) {
				newl = fmt.Sprintf("[%v]", l[0])
			} else {
				newl = l[0]
			}
			return compare(newl, right)
		}
	}
	l := splitlist(left)
	r := splitlist(right)
	shortest := int(math.Min(float64(len(l)), float64(len(r))))
	for i := 0; i < shortest; i++ {
		if isList(l[i]) || isList(r[i]) {
			res := compare(l[i], r[i])
			if res != 0 {
				return res
			}
		} else {
			lnum, _ := strconv.ParseInt(l[i], 0, 64)
			rnum, _ := strconv.ParseInt(r[i], 0, 64)
			if lnum < rnum {
				return 1
			} else if rnum < lnum {
				return -1
			}
		}
	}
	if len(l) < len(r) {
		return 1
	} else if len(l) > len(r) {
		return -1
	}
	return 0
}

func isList(line string) bool {
	if line[0] == '[' && line[len(line)-1] == ']' {
		if len(splitlist(line)) == 1 {
			return true
		}
	}
	return false
}

func splitlist(line string) []string {
	if line == "" {
		return strings.Split(line, ",")
	}
	result := make([]string, 0)
	depth := 0
	openindex := 0
	for i, char := range line {
		if char == '[' {
			depth++
		} else if char == ']' {
			depth--
		} else if depth == 0 && char == ',' {
			result = append(result, line[openindex:i])
			openindex = i + 1
		}
	}
	result = append(result, line[openindex:])
	return result
}
