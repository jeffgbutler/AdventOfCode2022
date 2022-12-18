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

	ranges = collapseRanges(ranges)
	total := 0
	for _, rg := range ranges {
		total += rg.max - rg.min + 1
	}

	return total - 1 // this assumes there is one and only one beacon in the range!
}
func collapseRanges(ranges []integerRange) []integerRange {
	currentRanges := ranges
	for {
		newRanges := collapseRangeIteration(currentRanges)
		if len(newRanges) == len(currentRanges) {
			return newRanges
		} else {
			currentRanges = newRanges
		}
	}
}

func collapseRangeIteration(ranges []integerRange) []integerRange {
	var answer []integerRange
	currentRange := ranges[0]

	for i := 1; i < len(ranges); i++ {
		nextRange := ranges[i]
		switch {
		case nextRange.min >= currentRange.min && nextRange.max <= currentRange.max:
			// next range contained within current range, ignore
			continue
		case nextRange.min > currentRange.max || nextRange.max < currentRange.min:
			// ranges don't overlap
			answer = append(answer, currentRange)
			currentRange = nextRange
		default:
			// ranges overlap
			currentRange.min = min(currentRange.min, nextRange.min)
			currentRange.max = max(currentRange.max, nextRange.max)
		}
	}

	answer = append(answer, currentRange)
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

func min(i1, i2 int) int {
	if i1 < i2 {
		return i1
	} else {
		return i2
	}
}

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	} else {
		return i2
	}
}
