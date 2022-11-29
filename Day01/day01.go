package day01

import (
    "fmt" 
	"adventofcode2022/filereader"
)

func Assignment1() {
	strings := filereader.ReadFile("../Datafiles/day01.txt")
	for _, s := range strings{
		fmt.Println(s)
	}
}
