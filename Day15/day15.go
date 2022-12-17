package Day15

import (
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type integerRange struct {
	min, max int
}

type sensorAndReach struct {
	sensor point
	beacon point
	reach  int
}

func part1(inputLines []string, row int) int {
	sensorsAndReach := parse(inputLines)
	var ranges []integerRange

	for _, sr := range sensorsAndReach {
		rowToSensorDistance := abs(row - sr.sensor.y)
		overlaps := rowToSensorDistance <= sr.reach
		if overlaps {
			overlap := abs(rowToSensorDistance - sr.reach)

			rg := integerRange{min: sr.sensor.x - overlap, max: sr.sensor.x + overlap}
			ranges = append(ranges, rg)
		}
	}

	rg := collapseRanges(ranges)
	return rg.max - rg.min // this assumes there is one and only one beacon in the range!
}

func collapseRanges(ranges []integerRange) integerRange {
	// this assumes that there are no gaps in the ranges!
	answer := integerRange{min: ranges[0].min, max: ranges[0].max}
	for _, rg := range ranges {
		if rg.min < answer.min {
			answer.min = rg.min
		}
		if rg.max > answer.max {
			answer.max = rg.max
		}
	}
	return answer
}

func parse(inputLines []string) []sensorAndReach {
	var answer []sensorAndReach

	for _, inputLine := range inputLines {
		parts := strings.Split(inputLine, ":")

		sb := sensorAndReach{}
		sb.sensor = parseSensorOrBeacon(parts[0])
		sb.beacon = parseSensorOrBeacon(parts[1])
		sb.reach = manhattanDistance(sb.sensor, sb.beacon)
		answer = append(answer, sb)
	}

	return answer
}

func parseSensorOrBeacon(s string) point {
	index := strings.Index(s, "x=")
	return parseXAndY(s[index:])
}

func parseXAndY(s string) point {
	// x=20, y=32
	parts := strings.Split(s, ", ")
	return point{x: parseNum(parts[0]), y: parseNum(parts[1])}
}

func parseNum(s string) int {
	parts := strings.Split(s, "=")
	i, _ := strconv.Atoi(parts[1])
	return i
}

func manhattanDistance(p1, p2 point) int {
	dx := abs(p1.x - p2.x)
	dy := abs(p1.y - p2.y)
	return dx + dy
}

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}
