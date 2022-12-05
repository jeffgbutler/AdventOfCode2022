package Day04

import (
	"strconv"
	"strings"
)

type singleRange struct {
	start, end int
}

type doubleRange struct {
	range1, range2 singleRange
}

func part1(inputLines []string) int {
	total := 0
	for _, v := range inputLines {
		ranges := convertToDoubleRange(v)
		if eitherRangeContained(ranges) {
			total++
		}
	}

	return total
}

func part2(inputLines []string) int {
	total := 0
	for _, v := range inputLines {
		ranges := convertToDoubleRange(v)
		if eitherRangeOverlaps(ranges) {
			total++
		}
	}

	return total
}

func convertToDoubleRange(rangesSpec string) doubleRange {
	f := strings.Split(rangesSpec, ",")
	return doubleRange{convertToSingleRange(f[0]), convertToSingleRange(f[1])}
}

func convertToSingleRange(rangeSpec string) singleRange {
	f := strings.Split(rangeSpec, "-")
	i1, _ := strconv.Atoi(f[0])
	i2, _ := strconv.Atoi(f[1])
	return singleRange{i1, i2}
}

func eitherRangeContained(ranges doubleRange) bool {
	return rangeContained(ranges.range1, ranges.range2) || rangeContained(ranges.range2, ranges.range1)
}

func rangeContained(range1, range2 singleRange) bool {
	return range1.start >= range2.start && range1.end <= range2.end
}

func eitherRangeOverlaps(ranges doubleRange) bool {
	return rangeOverlaps(ranges.range1, ranges.range2) || rangeContained(ranges.range2, ranges.range1)
}

func rangeOverlaps(range1, range2 singleRange) bool {
	return inRange(range1.start, range2) || inRange(range1.end, range2)
}

func inRange(i int, r singleRange) bool {
	return i >= r.start && i <= r.end
}
