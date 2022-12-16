package main

import (
	"adventofcode2022/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Valve struct {
	flow       int
	neighbours []string
	paths      map[string]int
}

type Path struct {
	nodes []string
	score int
}

func main() {
	input := utils.ReadFile("../Datafiles/day16.txt")
	area := make(map[string]Valve)
	for _, line := range input {
		parts := strings.Split(line, " ")
		name := parts[1]
		f, _ := strconv.ParseInt(parts[4][5:len(parts[4])-1], 0, 64)
		flow := int(f)
		paths := make([]string, 0)
		for _, path := range parts[9:] {
			paths = append(paths, path[0:2])
		}
		area[name] = Valve{flow: flow, neighbours: paths, paths: make(map[string]int)}
	}
	area = calcPaths(area)
	Assignment1(area)
	Assignment2(area)
}

func Assignment1(input map[string]Valve) {
	opened := make([]string, 0)
	fmt.Println(bestPath(input, "AA", 30, opened))
}

func Assignment2(input map[string]Valve) {
	paths := getPaths(input, 26, "AA", make([]string, 0))
	sort.Slice(paths, func(i, j int) bool {
		return paths[i].score > paths[j].score
	})
	maxscore := 0
	for _, p1 := range paths {
		if maxscore > 2*p1.score {
			break
		}
		for _, p2 := range paths {
			valid := true
			for _, part := range p1.nodes[:len(p1.nodes)-1] {
				if utils.Contains(p2.nodes[:len(p2.nodes)-1], part) {
					valid = false
				}
			}
			if valid {
				newscore := p1.score + p2.score
				if newscore > maxscore {
					maxscore = newscore
				}
			}
		}
	}
	fmt.Println(maxscore)
}

func bestPath(area map[string]Valve, location string, time int, opened []string) (string, int) {
	if time <= 0 {
		return location, 0
	}
	bestloc := location
	bestreward := 0
	paths := area[location].paths
	for l, d := range paths {
		if !utils.Contains(opened, l) {
			newopened := append(opened, l)
			newloc, newreward := bestPath(area, l, time-d, newopened)
			if newreward > bestreward {
				bestreward = newreward
				bestloc = newloc
			}
		}
	}
	result := area[location].flow * time
	return bestloc, result + bestreward
}

func pathLengths(area map[string]Valve, location string) map[string]int {
	paths := make(map[string]int)
	paths[location] = 1
	current := make([]string, 0)
	current = append(current, location)
	distance := 2
	for true {
		nextround := make([]string, 0)
		for _, loc := range current {
			for _, next := range area[loc].neighbours {
				if paths[next] == 0 {
					if area[next].flow == 0 {
						paths[next] = 40
					} else {
						paths[next] = distance
					}
					nextround = append(nextround, next)
				}
			}
		}
		if len(paths) == len(area) {
			break
		}
		current = nextround
		distance++
	}
	return paths
}

func calcPaths(area map[string]Valve) map[string]Valve {
	for name, valve := range area {
		paths := make(map[string]int)
		for n, d := range pathLengths(area, name) {
			if area[n].flow > 0 {
				paths[n] = d
			}
		}
		valve.paths = paths
		area[name] = valve
	}
	return area
}

func getPaths(area map[string]Valve, time int, loc string, opened []string) []Path {
	result := make([]Path, 0)

	if time <= 0 {
		result = append(result, Path{make([]string, 0), 0})
		return result
	}
	valve := area[loc]
	newopened := append(opened, loc)

	for n, d := range valve.paths {
		if !utils.Contains(newopened, n) {
			response := getPaths(area, time-d, n, newopened)
			for _, p := range response {
				p.nodes = append(p.nodes, loc)
				p.score += time * valve.flow
				result = append(result, p)
			}
		}
	}
	return result
}
