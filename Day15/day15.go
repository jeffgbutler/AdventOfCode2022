package Day15

import (
	"sort"
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
	ranges, beaconsInTheRow := calculateRowRanges(sensorsAndReach, row)
	total := 0
	for _, rg := range ranges {
		total += rg.max - rg.min + 1
	}

	return total - beaconsInTheRow
}

func calculateRowRanges(sensorsAndReach []sensorAndReach, row int) ([]integerRange, int) {
	var ranges []integerRange
	beaconSet := map[point]bool{}

	for _, sr := range sensorsAndReach {
		beaconSet[sr.beacon] = true
		rowToSensorDistance := abs(row - sr.sensor.y)
		overlaps := rowToSensorDistance <= sr.reach
		if overlaps {
			overlap := abs(rowToSensorDistance - sr.reach)

			rg := integerRange{min: sr.sensor.x - overlap, max: sr.sensor.x + overlap}
			ranges = append(ranges, rg)
		}
	}

	ranges = collapseRanges(ranges)

	return ranges, beaconsInRow(beaconSet, row, ranges)
}

func part2(inputLines []string, maxCoordinate int) int {
	sensorsAndReach := parse(inputLines)
	row, slot := findRowWithOneOpenSlot(sensorsAndReach, maxCoordinate)

	return slot*4000000 + row
}

func findRowWithOneOpenSlot(sensorsAndReach []sensorAndReach, maxCoordinate int) (row, slot int) {
	// find a row with only 1 open slot
	for i := 0; i < maxCoordinate; i++ {
		ranges, _ := calculateRowRanges(sensorsAndReach, i)
		ranges = normalizeRanges(ranges, maxCoordinate)
		slots := findSingleSlots(ranges, maxCoordinate)
		if len(slots) == 1 {
			return i, slots[0]
		}
	}

	return -1, -1
}

func findSingleSlots(ranges []integerRange, maxCoordinate int) []int {
	var slots []int

	switch len(ranges) {
	case 1:
		if ranges[0].min == 1 {
			slots = append(slots, 0)
		} else if ranges[0].max == maxCoordinate-1 {
			slots = append(slots, maxCoordinate)
		}
	default:
		for i := 0; i < len(ranges)-1; i++ {
			r1 := ranges[i]
			r2 := ranges[i+1]
			if r2.min == r1.max+2 {
				slots = append(slots, r1.max+1)
			}
		}
		// TODO check for open slot at beginning and end
	}

	return slots
}

func collapseRanges(ranges []integerRange) []integerRange {
	currentRanges := ranges
	sort.Slice(currentRanges, func(i, j int) bool { return currentRanges[i].min < currentRanges[j].min })
	for {
		newRanges := collapseRangeIteration(currentRanges)
		if len(newRanges) == len(currentRanges) {
			return newRanges
		} else {
			currentRanges = newRanges
		}
	}
}

func beaconsInRow(beaconSet map[point]bool, row int, ranges []integerRange) int {
	count := 0

	for beacon := range beaconSet {
		if beacon.y != row {
			continue
		}

		if beaconIsInRow(beacon, ranges) {
			count++
		}
	}

	return count
}

func beaconIsInRow(beacon point, ranges []integerRange) bool {
	for _, rg := range ranges {
		if beacon.x >= rg.min && beacon.x <= rg.max {
			return true
		}
	}

	return false
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
		case currentRange.max == nextRange.min-1:
			// ranges adjacent
			currentRange.max = max(currentRange.max, nextRange.max)
		case currentRange.min == nextRange.max+1:
			// ranges adjacent
			currentRange.min = min(currentRange.min, nextRange.min)
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

func normalizeRanges(ranges []integerRange, maxCoordinate int) []integerRange {
	var answer []integerRange

	for _, rg := range ranges {
		switch {
		case rg.min >= 0 && rg.max <= maxCoordinate:
			answer = append(answer, rg)
		case rg.min < 0:
			if rg.max <= maxCoordinate {
				answer = append(answer, integerRange{0, rg.max})
			} else {
				answer = append(answer, integerRange{0, maxCoordinate})
			}
		case rg.max > maxCoordinate:
			if rg.min < 0 {
				answer = append(answer, integerRange{0, maxCoordinate})
			} else {
				answer = append(answer, integerRange{rg.min, maxCoordinate})
			}
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
