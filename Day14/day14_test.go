package Day14

import (
	"github.com/jeffgbutler/AdventOfCode2022/inputfile"
	"testing"
)

func Test_Part1WithExampleData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/example.txt")

	answer := part1(inputLines)

	if answer != 24 {
		t.Error("Expected 24, got", answer)
	}
}

func Test_Part1WithActualData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/actual.txt")

	answer := part1(inputLines)

	if answer != 638 {
		t.Error("Expected 638, got", answer)
	}
}

func Test_Part2WithExampleData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/example.txt")

	answer := part2(inputLines)

	if answer != 93 {
		t.Error("Expected 93, got", answer)
	}
}

func Test_Part2WithActualData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/actual.txt")

	answer := part2(inputLines)

	if answer != 31722 {
		t.Error("Expected 31722, got", answer)
	}
}

func Test_ToString(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/example.txt")
	grid := parseInputLines(inputLines)
	strings := toString(grid)

	if len(strings) != 6 {
		t.Error("Expected 6, got", len(strings))
	}
}
