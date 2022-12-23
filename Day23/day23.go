package main

import (
	"adventofcode2022/utils"
	"fmt"
)

type Elf struct {
	x    int
	y    int
	desx int
	desy int
}

func main() {
	input := utils.ReadFile("../Datafiles/day23.txt")

	Assignment1(input)
	Assignment2(input)
}

func Assignment1(input []string) {
	field := make(map[string]bool)
	elves := make([]Elf, 0)
	for i, line := range input {
		for j, char := range line {
			if char == '#' {
				field[fmt.Sprintf("%v,%v", j, i)] = true
				elves = append(elves, Elf{x: j, y: i})
			}
		}
	}
	order := make([]func(int, int, map[string]bool) (int, int), 4)
	order[0] = checkNorth
	order[1] = checkSouth
	order[2] = checkWest
	order[3] = checkEast
	for i := 0; i < 10; i++ {
		moves := make(map[string]bool, 0)
		forbidden := make(map[string]bool, 0)
		for c, elf := range elves {
			if checkAll(elf.x, elf.y, field) {
				elves[c].desx = elf.x
				elves[c].desy = elf.y
			} else {
				for j := 0; j < 4; j++ {
					newx, newy := order[(i+j)%4](elf.x, elf.y, field)
					if elf.x != newx || elf.y != newy {
						elves[c].desx = newx
						elves[c].desy = newy
						break
					}
					elves[c].desx = newx
					elves[c].desy = newy
				}
			}
			cords := fmt.Sprintf("%v,%v", elves[c].desx, elves[c].desy)
			if moves[cords] {
				forbidden[cords] = true
			} else {
				moves[cords] = true
			}
		}

		for c, elf := range elves {
			cords := fmt.Sprintf("%v,%v", elf.desx, elf.desy)
			if !forbidden[cords] {
				elves[c].x = elf.desx
				elves[c].y = elf.desy
				elves[c].desx = 0
				elves[c].desy = 0
			} else {
				elves[c].desx = 0
				elves[c].desy = 0
			}
		}
		field = buildMap(elves)
	}
	minx := 5
	maxx := 0
	miny := 5
	maxy := 0
	for _, elf := range elves {
		if elf.x < minx {
			minx = elf.x
		}
		if elf.x > maxx {
			maxx = elf.x
		}
		if elf.y < miny {
			miny = elf.y
		}
		if elf.y > maxy {
			maxy = elf.y
		}
	}
	deltax := maxx - minx + 1
	deltay := maxy - miny + 1
	fmt.Println(deltax*deltay - len(elves))
}

func Assignment2(input []string) {
	field := make(map[string]bool)
	elves := make([]Elf, 0)
	for i, line := range input {
		for j, char := range line {
			if char == '#' {
				field[fmt.Sprintf("%v,%v", j, i)] = true
				elves = append(elves, Elf{x: j, y: i})
			}
		}
	}
	order := make([]func(int, int, map[string]bool) (int, int), 4)
	order[0] = checkNorth
	order[1] = checkSouth
	order[2] = checkWest
	order[3] = checkEast
	running := true
	round := 0
	for i := 0; running; i++ {
		round++
		moves := make(map[string]bool, 0)
		forbidden := make(map[string]bool, 0)
		for c, elf := range elves {
			if checkAll(elf.x, elf.y, field) {
				elves[c].desx = elf.x
				elves[c].desy = elf.y
			} else {
				for j := 0; j < 4; j++ {
					newx, newy := order[(i+j)%4](elf.x, elf.y, field)
					if elf.x != newx || elf.y != newy {
						cords := fmt.Sprintf("%v,%v", newx, newy)
						if moves[cords] {
							forbidden[cords] = true
						} else {
							moves[cords] = true
						}
						elves[c].desx = newx
						elves[c].desy = newy
						break
					}
					elves[c].desx = newx
					elves[c].desy = newy
				}
			}
		}

		if len(moves)-len(forbidden) == 0 {
			running = false
		}

		for c, elf := range elves {
			cords := fmt.Sprintf("%v,%v", elf.desx, elf.desy)
			if !forbidden[cords] {
				elves[c].x = elf.desx
				elves[c].y = elf.desy
				elves[c].desx = 0
				elves[c].desy = 0
			} else {
				elves[c].desx = 0
				elves[c].desy = 0
			}
		}
		field = buildMap(elves)
	}
	fmt.Println(round)
}

func buildMap(elves []Elf) map[string]bool {
	field := make(map[string]bool)
	for _, elf := range elves {
		cords := fmt.Sprintf("%v,%v", elf.x, elf.y)
		field[cords] = true
	}
	return field
}

func checkNorth(x, y int, field map[string]bool) (int, int) {
	nw := fmt.Sprintf("%v,%v", x-1, y-1)
	n := fmt.Sprintf("%v,%v", x, y-1)
	ne := fmt.Sprintf("%v,%v", x+1, y-1)
	if field[ne] || field[n] || field[nw] {
		return x, y
	}
	return x, y - 1
}
func checkSouth(x, y int, field map[string]bool) (int, int) {
	sw := fmt.Sprintf("%v,%v", x-1, y+1)
	s := fmt.Sprintf("%v,%v", x, y+1)
	se := fmt.Sprintf("%v,%v", x+1, y+1)
	if field[se] || field[s] || field[sw] {
		return x, y
	}
	return x, y + 1
}
func checkWest(x, y int, field map[string]bool) (int, int) {
	nw := fmt.Sprintf("%v,%v", x-1, y-1)
	w := fmt.Sprintf("%v,%v", x-1, y)
	sw := fmt.Sprintf("%v,%v", x-1, y+1)
	if field[sw] || field[w] || field[nw] {
		return x, y
	}
	return x - 1, y
}
func checkEast(x, y int, field map[string]bool) (int, int) {
	ne := fmt.Sprintf("%v,%v", x+1, y-1)
	e := fmt.Sprintf("%v,%v", x+1, y)
	se := fmt.Sprintf("%v,%v", x+1, y+1)
	if field[ne] || field[e] || field[se] {
		return x, y
	}
	return x + 1, y
}

func checkAll(x, y int, field map[string]bool) bool {
	nw := fmt.Sprintf("%v,%v", x-1, y-1)
	n := fmt.Sprintf("%v,%v", x, y-1)
	ne := fmt.Sprintf("%v,%v", x+1, y-1)
	sw := fmt.Sprintf("%v,%v", x-1, y+1)
	s := fmt.Sprintf("%v,%v", x, y+1)
	se := fmt.Sprintf("%v,%v", x+1, y+1)
	w := fmt.Sprintf("%v,%v", x-1, y)
	e := fmt.Sprintf("%v,%v", x+1, y)
	if field[ne] || field[e] || field[se] || field[nw] || field[w] || field[sw] || field[n] || field[s] {
		return false
	}
	return true
}
