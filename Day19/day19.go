package main

import (
	"adventofcode2022/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Blueprint struct {
	id       int
	orebot   int
	claybot  int
	obsboto  int
	obsbotc  int
	geoboto  int
	geobotob int
}

func main() {
	input := utils.ReadFile("../Datafiles/day19.txt")
	blueprints := make([]Blueprint, 0)
	for _, line := range input {
		parts := strings.Split(line, " ")
		num, _ := strconv.ParseInt(parts[1][:len(parts[1])-1], 0, 64)
		id := int(num)
		num, _ = strconv.ParseInt(parts[6], 0, 64)
		ore := int(num)
		num, _ = strconv.ParseInt(parts[12], 0, 64)
		clay := int(num)
		num, _ = strconv.ParseInt(parts[18], 0, 64)
		obso := int(num)
		num, _ = strconv.ParseInt(parts[21], 0, 64)
		obsc := int(num)
		num, _ = strconv.ParseInt(parts[27], 0, 64)
		geoo := int(num)
		num, _ = strconv.ParseInt(parts[30], 0, 64)
		geoob := int(num)
		blueprint := Blueprint{
			id:       id,
			orebot:   ore,
			claybot:  clay,
			obsboto:  obso,
			obsbotc:  obsc,
			geoboto:  geoo,
			geobotob: geoob}
		blueprints = append(blueprints, blueprint)
	}
	Assignment1(blueprints)
	Assignment2(blueprints)
}

func Assignment1(input []Blueprint) {
	result := 0
	for _, blue := range input {
		reward := calcUse(blue, 24, 1, 0, 0, 0, 0, 0, 0, 0)
		result += blue.id * reward
	}
	fmt.Println(result)
}
func Assignment2(input []Blueprint) {
	result := 1
	blueprints := input[:3]
	for _, blue := range blueprints {
		reward := calcUse(blue, 32, 1, 0, 0, 0, 0, 0, 0, 0)
		result *= reward
	}
	fmt.Println(result)
}

func calcUse(blueprint Blueprint, time, boto, botc, botob, botg, o, c, ob, g int) int {
	bestGather := 0
	if botob != 0 && boto != 0 {
		obam := blueprint.geobotob - ob
		oam := blueprint.geoboto - o
		if oam < 0 {
			oam = 0
		}
		if obam < 0 {
			obam = 0
		}
		obtime := int(math.Ceil(float64(obam) / float64(botob)))
		otime := int(math.Ceil(float64(oam) / float64(boto)))
		timespent := utils.Max(obtime, otime) + 1
		if timespent < time {
			newReward := calcUse(blueprint, time-timespent, boto, botc, botob, botg+1, o+boto*timespent-blueprint.geoboto, c+botc*timespent, ob+botob*timespent-blueprint.geobotob, g+botg*timespent)
			if newReward > bestGather {
				bestGather = newReward
			}
		}
	}
	if botc != 0 && boto != 0 {
		cam := blueprint.obsbotc - c
		oam := blueprint.obsboto - o
		if oam < 0 {
			oam = 0
		}
		if cam < 0 {
			cam = 0
		}
		ctime := int(math.Ceil(float64(cam) / float64(botc)))
		otime := int(math.Ceil(float64(oam) / float64(boto)))
		timespent := utils.Max(ctime, otime) + 1
		if timespent < time {
			newReward := calcUse(blueprint, time-timespent, boto, botc, botob+1, botg, o+boto*timespent-blueprint.obsboto, c+botc*timespent-blueprint.obsbotc, ob+botob*timespent, g+botg*timespent)
			if newReward > bestGather {
				bestGather = newReward
			}
		}
	}
	if botc <= blueprint.obsbotc {
		oam := blueprint.claybot - o
		if oam < 0 {
			oam = 0
		}
		timespent := int(math.Ceil(float64(oam)/float64(boto))) + 1
		if timespent < time {
			newReward := calcUse(blueprint, time-timespent, boto, botc+1, botob, botg, o+boto*timespent-blueprint.claybot, c+botc*timespent, ob+botob*timespent, g+botg*timespent)
			if newReward > bestGather {
				bestGather = newReward
			}
		}
	}
	if boto <= utils.Max(utils.Max(blueprint.geoboto, blueprint.obsboto), blueprint.claybot) {
		oam := blueprint.orebot - o
		if oam < 0 {
			oam = 0
		}
		timespent := int(math.Ceil(float64(oam)/float64(boto))) + 1
		if timespent < time {
			newReward := calcUse(blueprint, time-timespent, boto+1, botc, botob, botg, o+boto*timespent-blueprint.orebot, c+botc*timespent, ob+botob*timespent, g+botg*timespent)
			if newReward > bestGather {
				bestGather = newReward
			}
		}
	}
	o += boto * time
	c += botc * time
	ob += botob * time
	g += botg * time
	if g > bestGather {
		bestGather = g
	}
	return bestGather
}
