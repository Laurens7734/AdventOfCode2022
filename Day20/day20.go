package main

import (
	"adventofcode2022/utils"
	"fmt"
	"strconv"
)

type Number struct {
	value        int
	isfirst      bool
	previus      *Number
	next         *Number
	originalnext *Number
}

func main() {
	input := utils.ReadFile("../Datafiles/day20.txt")
	Assignment1(input)
	Assignment2(input)
}

func Assignment1(input []string) {
	var first Number
	var zero *Number
	var previus *Number
	numberofitems := 0
	for i, line := range input {
		numberofitems++
		num, _ := strconv.ParseInt(line, 0, 64)
		number := Number{value: int(num), previus: previus, isfirst: false}
		if i != 0 {
			(*previus).next = &number
			(*previus).originalnext = &number
		}
		if number.value == 0 {
			zero = &number
		}
		if i == 1 {
			previus.isfirst = true
			first = *previus
		}
		previus = &number
	}
	previus.next = &first
	previus.originalnext = &first
	first.previus = previus

	move(&first, numberofitems)
	tomove := first.originalnext
	for !tomove.isfirst {
		move(tomove, numberofitems)
		tomove = tomove.originalnext
	}
	result := 0
	point := zero
	for i := 0; i < 3000; i++ {
		point = point.next
		if i%1000 == 999 {
			result += point.value
		}
	}
	fmt.Println(result)
}
func Assignment2(input []string) {
	var first Number
	var zero *Number
	var previus *Number
	numberofitems := 0
	for i, line := range input {
		numberofitems++
		num, _ := strconv.ParseInt(line, 0, 64)
		number := Number{value: int(num) * 811589153, previus: previus, isfirst: false}
		if i != 0 {
			(*previus).next = &number
			(*previus).originalnext = &number
		}
		if number.value == 0 {
			zero = &number
		}
		if i == 1 {
			previus.isfirst = true
			first = *previus
		}
		previus = &number
	}
	previus.next = &first
	previus.originalnext = &first
	first.previus = previus

	for i := 0; i < 10; i++ {
		move(&first, numberofitems)
		tomove := first.originalnext
		for !tomove.isfirst {
			move(tomove, numberofitems)
			tomove = tomove.originalnext
		}
	}

	result := 0
	point := zero
	for i := 0; i < 3000; i++ {
		point = point.next
		if i%1000 == 999 {
			result += point.value
		}
	}
	fmt.Println(result)
}

func move(item *Number, numberofitems int) {
	(*item.previus).next = item.next
	(*item.next).previus = item.previus
	insertNextTo := item.previus
	for i := 0; i < utils.Abs(item.value)%(numberofitems-1); i++ {
		if item.value < 0 {
			insertNextTo = insertNextTo.previus
		} else {
			insertNextTo = insertNextTo.next
		}
	}
	newNeighbour := (*insertNextTo).next
	(*insertNextTo).next = item
	(*newNeighbour).previus = item
	(*item).previus = insertNextTo
	(*item).next = newNeighbour
}
