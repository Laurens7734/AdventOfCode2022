package main

import (
	"adventofcode2022/filereader"
	"fmt"
)

func main() {
	Assignment1()
}

func Assignment1() {
	strings := filereader.ReadFile("../Datafiles/day02.txt")
	for _, s := range strings {
		fmt.Println(s)
	}
}
