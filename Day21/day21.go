package main

import (
	"adventofcode2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := utils.ReadFile("../Datafiles/day21.txt")
	Assignment1(input)
	Assignment2(input)
}

func Assignment1(input []string) {
	results := make(map[string]int)
	waiting := make([]string, 0)
	for _, line := range input {
		parts := strings.Split(line, " ")
		if len(parts) == 2 {
			num, _ := strconv.ParseInt(parts[1], 0, 64)
			results[parts[0][:4]] = int(num)
		} else {
			waiting = append(waiting, line)
		}
	}
	for results["root"] == 0 {
		results, waiting = addResults(results, waiting)
	}
	fmt.Println(results["root"])
}

func Assignment2(input []string) {
	results := make(map[string]int)
	waiting := make([]string, 0)
	opperations := make([]string, 0)
	next := "humn"
	for _, line := range input {
		parts := strings.Split(line, " ")
		if len(parts) == 2 {
			num, _ := strconv.ParseInt(parts[1], 0, 64)
			results[parts[0][:4]] = int(num)
		} else {
			if parts[1] == next || parts[3] == next {
				opperations = append(opperations, line)
				next = parts[0][:4]
			}
			waiting = append(waiting, line)
		}
	}
	for next != "root" {
		for _, line := range input {
			parts := strings.Split(line, " ")
			if len(parts) != 2 {
				if parts[1] == next || parts[3] == next {
					opperations = append(opperations, line)
					next = parts[0][:4]
				}
			}
		}
	}
	for results["root"] == 0 {
		results, waiting = addResults(results, waiting)
	}
	for _, line := range opperations {
		results[line[:4]] = 0
	}
	results["humn"] = 0
	toyell := backcalc(results, opperations)
	fmt.Println(toyell)
}

func addResults(results map[string]int, tocalc []string) (map[string]int, []string) {
	newresults := results
	notplaced := make([]string, 0)
	for _, line := range tocalc {
		parts := strings.Split(line, " ")
		if results[parts[1]] != 0 && results[parts[3]] != 0 {
			if parts[2] == "+" {
				newresults[parts[0][:4]] = results[parts[1]] + results[parts[3]]
			}
			if parts[2] == "-" {
				newresults[parts[0][:4]] = results[parts[1]] - results[parts[3]]
			}
			if parts[2] == "*" {
				newresults[parts[0][:4]] = results[parts[1]] * results[parts[3]]
			}
			if parts[2] == "/" {
				newresults[parts[0][:4]] = results[parts[1]] / results[parts[3]]
			}
		} else {
			notplaced = append(notplaced, line)
		}
	}
	return newresults, notplaced
}

func backcalc(results map[string]int, opperations []string) int {
	parts := strings.Split(opperations[len(opperations)-1], " ")
	result := 0
	if results[parts[1]] != 0 {
		result = results[parts[1]]
	} else {
		result = results[parts[3]]
	}
	for i := len(opperations) - 2; i >= 0; i-- {
		parts = strings.Split(opperations[i], " ")
		knownval := 0
		firstknow := false
		if results[parts[1]] == 0 {
			knownval = results[parts[3]]
		} else {
			knownval = results[parts[1]]
			firstknow = true
		}
		if parts[2] == "+" {
			result = result - knownval
		}
		if parts[2] == "-" {
			if firstknow {
				result = knownval - result
			} else {
				result = result + knownval
			}

		}
		if parts[2] == "*" {
			result = result / knownval
		}
		if parts[2] == "/" {
			if firstknow {
				result = knownval / result
			} else {
				result = result * knownval
			}
		}
	}
	return result
}
