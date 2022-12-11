package main

import (
	"adventofcode2022/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	id          int
	test        func(int) int
	operation   func(int) int
	items       []int
	inspections int
}

var alltests = 0

func main() {
	input := utils.ReadFile("../Datafiles/day11.txt")
	Assignment1(input)
	Assignment2(input)
}

func Assignment1(input []string) {
	monkeys := createMonkeys(input)
	for i := 0; i < 20; i++ {
		for j, monkey := range monkeys {
			for _, item := range monkey.items {
				newval := monkey.operation(item)
				monkeys[j].inspections++
				newval /= 3
				monkeys[monkey.test(newval)].items = utils.AppendSlice(monkeys[monkey.test(newval)].items, newval)
			}
			monkeys[j].items = make([]int, 0)
		}
	}
	inspections := make([]int, 0)
	for _, monkey := range monkeys {
		inspections = utils.AppendSlice(inspections, monkey.inspections)
	}
	sort.Ints(inspections)
	fmt.Println(inspections[len(inspections)-1] * inspections[len(inspections)-2])
}

func Assignment2(input []string) {
	alltests = 1
	monkeys := createMonkeys(input)
	for i := 0; i < 10000; i++ {
		for j, monkey := range monkeys {
			for _, item := range monkey.items {
				newval := monkey.operation(item)
				newval %= alltests
				monkeys[j].inspections++
				monkeys[monkey.test(newval)].items = utils.AppendSlice(monkeys[monkey.test(newval)].items, newval)
			}
			monkeys[j].items = make([]int, 0)
		}
	}
	inspections := make([]int, 0)
	for _, monkey := range monkeys {
		inspections = utils.AppendSlice(inspections, monkey.inspections)
	}
	sort.Ints(inspections)
	fmt.Println(inspections[len(inspections)-1] * inspections[len(inspections)-2])
}

func createMonkeys(input []string) []Monkey {
	monkeys := make([]Monkey, 0)
	var id int
	var operation func(int) int
	var test func(int) int
	var items []int
	for i, line := range input {
		if line == "" {
			monkey := Monkey{id: id, test: test, operation: operation, items: items, inspections: 0}
			id = 0
			operation = nil
			test = nil
			items = nil
			monkeys = utils.AppendSlice(monkeys, monkey)
		} else if line[0] == 'M' {
			num, _ := strconv.ParseInt(line[7:8], 0, 64)
			id = int(num)
		} else {
			parts := strings.Fields(line)
			if parts[0] == "Starting" {
				iitems := make([]int, 0)
				sitems := parts[2:]
				for _, item := range sitems {
					num, _ := strconv.ParseInt(TrimSuffix(item, ","), 0, 64)
					iitems = utils.AppendSlice(iitems, int(num))
				}
				items = iitems
			} else if parts[0] == "Operation:" {
				operation = buildOperation(parts[3:])
			} else if parts[0] == "Test:" {
				num, _ := strconv.ParseInt(TrimSuffix(parts[3], ","), 0, 64)
				alltests *= int(num)
				truemonkey, _ := strconv.ParseInt(input[i+1][29:], 0, 64)
				falsemonkey, _ := strconv.ParseInt(input[i+2][30:], 0, 64)
				test = func(a int) int {
					if a%int(num) == 0 {
						return int(truemonkey)
					} else {
						return int(falsemonkey)
					}
				}
			}
		}
	}
	monkey := Monkey{id: id, test: test, operation: operation, items: items}
	monkeys = utils.AppendSlice(monkeys, monkey)
	return monkeys
}

func buildOperation(parts []string) func(int) int {
	result := func(a int) int {
		var1 := 0
		var2 := 0
		if parts[0] == "old" {
			var1 = a
		} else {
			num, _ := strconv.ParseInt(TrimSuffix(parts[0], ","), 0, 64)
			var1 = int(num)
		}
		if parts[2] == "old" {
			var2 = a
		} else {
			num, _ := strconv.ParseInt(TrimSuffix(parts[2], ","), 0, 64)
			var2 = int(num)
		}
		switch parts[1] {
		case "+":
			return var1 + var2
		case "-":
			return var1 - var2
		case "*":
			return var1 * var2
		}
		return -1
	}
	return result
}

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}
