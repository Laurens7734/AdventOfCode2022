package main

import (
	"adventofcode2022/utils"
	"fmt"
	"strconv"
	"strings"
)

type Storm struct {
	x   int
	y   int
	dir rune
}

var borderl = 0
var borderr = 0
var borderu = 0
var borderd = 0

func main() {
	input := utils.ReadFile("../Datafiles/day24.txt")
	borderr = len(input[0]) - 1
	borderd = len(input) - 1
	start, end, storms := parseInput(input)
	Assignment1(start, end, storms)
	Assignment2(start, end, storms)
}

func Assignment1(start string, end string, storms []Storm) {
	possiblePositions := make(map[string]bool)
	borders := make(map[string]bool)
	for i := 0; i <= borderd; i++ {
		if i == borderu {
			for j := 0; j <= borderr; j++ {
				cord := fmt.Sprintf("%v,%v", j, i)
				if cord != start {
					borders[cord] = true
				}
			}
		} else if i == borderd {
			for j := 0; j <= borderr; j++ {
				cord := fmt.Sprintf("%v,%v", j, i)
				if cord != end {
					borders[cord] = true
				}
			}
		} else {
			left := fmt.Sprintf("%v,%v", borderl, i)
			right := fmt.Sprintf("%v,%v", borderr, i)
			borders[left] = true
			borders[right] = true
		}
	}
	possiblePositions[start] = true
	count := 0
	searching := true
	for searching {
		count++
		storms = moveStorms(storms)
		stormcords := getStormCords(storms)
		possiblePositions = allMoves(possiblePositions, stormcords, borders)
		if possiblePositions[end] {
			searching = false
		}
	}
	fmt.Println(count)
}

func Assignment2(start string, end string, storms []Storm) {
	possiblePositions := make(map[string]bool)
	borders := make(map[string]bool)
	for i := 0; i <= borderd; i++ {
		if i == borderu {
			for j := 0; j <= borderr; j++ {
				cord := fmt.Sprintf("%v,%v", j, i)
				if cord != start {
					borders[cord] = true
				}
			}
		} else if i == borderd {
			for j := 0; j <= borderr; j++ {
				cord := fmt.Sprintf("%v,%v", j, i)
				if cord != end {
					borders[cord] = true
				}
			}
		} else {
			left := fmt.Sprintf("%v,%v", borderl, i)
			right := fmt.Sprintf("%v,%v", borderr, i)
			borders[left] = true
			borders[right] = true
		}
	}
	possiblePositions[start] = true
	count := 0
	searching := true
	for searching {
		count++
		storms = moveStorms(storms)
		stormcords := getStormCords(storms)
		possiblePositions = allMoves(possiblePositions, stormcords, borders)
		if possiblePositions[end] {
			searching = false
		}
	}
	possiblePositions = make(map[string]bool)
	possiblePositions[end] = true
	searching = true
	for searching {
		count++
		storms = moveStorms(storms)
		stormcords := getStormCords(storms)
		possiblePositions = allMoves(possiblePositions, stormcords, borders)
		if possiblePositions[start] {
			searching = false
		}
	}
	possiblePositions = make(map[string]bool)
	possiblePositions[start] = true
	searching = true
	for searching {
		count++
		storms = moveStorms(storms)
		stormcords := getStormCords(storms)
		possiblePositions = allMoves(possiblePositions, stormcords, borders)
		if possiblePositions[end] {
			searching = false
		}
	}
	fmt.Println(count)
}

func allMoves(currentMoves map[string]bool, stormlocations map[string]bool, border map[string]bool) map[string]bool {
	positions := make(map[string]bool)
	for pos, _ := range currentMoves {
		if !stormlocations[pos] {
			positions[pos] = true
		}
		parts := strings.Split(pos, ",")
		num, _ := strconv.ParseInt(parts[0], 0, 64)
		x := int(num)
		num, _ = strconv.ParseInt(parts[1], 0, 64)
		y := int(num)
		if x-1 >= borderl {
			newpos := fmt.Sprintf("%v,%v", x-1, y)
			if !stormlocations[newpos] && !border[newpos] {
				positions[newpos] = true
			}
		}
		if x+1 <= borderr {
			newpos := fmt.Sprintf("%v,%v", x+1, y)
			if !stormlocations[newpos] && !border[newpos] {
				positions[newpos] = true
			}
		}
		if y-1 >= borderu {
			newpos := fmt.Sprintf("%v,%v", x, y-1)
			if !stormlocations[newpos] && !border[newpos] {
				positions[newpos] = true
			}
		}
		if y+1 <= borderd {
			newpos := fmt.Sprintf("%v,%v", x, y+1)
			if !stormlocations[newpos] && !border[newpos] {
				positions[newpos] = true
			}
		}
	}
	return positions
}

func moveStorms(storms []Storm) []Storm {
	newStorms := make([]Storm, 0)
	for _, storm := range storms {
		if storm.dir == 'U' {
			storm.y--
			if storm.y == borderu {
				storm.y = borderd - 1
			}
		} else if storm.dir == 'D' {
			storm.y++
			if storm.y == borderd {
				storm.y = borderu + 1
			}
		} else if storm.dir == 'L' {
			storm.x--
			if storm.x == borderl {
				storm.x = borderr - 1
			}
		} else if storm.dir == 'R' {
			storm.x++
			if storm.x == borderr {
				storm.x = borderl + 1
			}
		}
		newStorms = append(newStorms, storm)
	}
	return newStorms
}

func getStormCords(storms []Storm) map[string]bool {
	cords := make(map[string]bool)
	for _, storm := range storms {
		cord := fmt.Sprintf("%v,%v", storm.x, storm.y)
		cords[cord] = true
	}
	return cords
}

func parseInput(input []string) (string, string, []Storm) {
	start := ""
	end := ""
	storms := make([]Storm, 0)

	for i, line := range input {
		if i == borderu {
			for j, c := range line {
				if c == '.' {
					start = fmt.Sprintf("%v,%v", j, i)
					break
				}
			}
		} else if i == borderd {
			for j, c := range line {
				if c == '.' {
					end = fmt.Sprintf("%v,%v", j, i)
					break
				}
			}
		} else {
			for j, c := range line {
				if c != '.' && c != '#' {
					storm := Storm{x: j, y: i}
					if c == 'v' {
						storm.dir = 'D'
					} else if c == '>' {
						storm.dir = 'R'
					} else if c == '<' {
						storm.dir = 'L'
					} else if c == '^' {
						storm.dir = 'U'
					}
					storms = append(storms, storm)
				}
			}
		}
	}
	return start, end, storms
}
