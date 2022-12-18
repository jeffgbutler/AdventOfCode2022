package Day15

import (
	"github.com/jeffgbutler/AdventOfCode2022/inputfile"
	"testing"
)

func Test_Part1WithExampleData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/example.txt")

	answer := part1(inputLines, 10)

	if answer != 26 {
		t.Error("Expected 26, got", answer)
	}
}

func Test_Part1WithExampleDataRow11(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/example.txt")

	answer := part1(inputLines, 11)

	if answer != 28 {
		t.Error("Expected 28, got", answer)
	}
}

func Test_Part1WithActualData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/actual.txt")

	answer := part1(inputLines, 2_000_000)

	if answer != 5511201 {
		t.Error("Expected 5,511,201, got", answer)
	}
}

func Test_Part2WithExampleData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/example.txt")

	answer := part2(inputLines, 20)

	if answer != 56000011 {
		t.Error("Expected 56000011, got", answer)
	}
}

func Test_Part2WithActualData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/actual.txt")

	answer := part2(inputLines, 4000000)

	if answer != 11318723411840 {
		t.Error("Expected 11318723411840, got", answer)
	}
}

func Test_RangeCollapse(t *testing.T) {
	ranges := []integerRange{{1, 3}, {4, 14}, {-2, 2}, {14, 26}}

	ranges = collapseAndSortRanges(ranges)

	if len(ranges) != 1 {
		t.Error("Expected 1, got", len(ranges))
	}

	if ranges[0].min != -2 || ranges[0].max != 26 {
		t.Error("Wrong range:", ranges[0])
	}
}
