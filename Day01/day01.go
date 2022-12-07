package main

import (
	"fmt"
	"strconv"
)

func main() {
	Assignment1()
	Assignment2()
}

func Assignment1() {
	max := int64(0)
	cur := int64(0)
	for _, s := range utils.ReadFile("../Datafiles/day01.txt") {
		if s != "" {
			num, _ := strconv.ParseInt(s, 0, 64)
			cur += num
		} else {
			if cur > max {
				max = cur
			}
			cur = int64(0)
		}
	}
	fmt.Println(max)
}

func Assignment2() {
	max1 := int64(0)
	max2 := int64(0)
	max3 := int64(0)
	cur := int64(0)
	for _, s := range utils.ReadFile("../Datafiles/day01.txt") {
		if s != "" {
			num, _ := strconv.ParseInt(s, 0, 64)
			cur += num
		} else {
			if cur > max1 {
				max3 = max2
				max2 = max1
				max1 = cur
			} else if cur > max2 {
				max3 = max2
				max2 = cur
			} else if cur > max3 {
				max3 = cur
			}
			cur = int64(0)
		}
	}
	fmt.Println(max1 + max2 + max3)
}
