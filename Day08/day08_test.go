package Day08

import (
	"github.com/jeffgbutler/AdventOfCode2022/inputfile"
	"testing"
)

func Test_Part1WithExampleData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/example.txt")
	answer := part1(inputLines)

	if answer != 21 {
		t.Error("Expected 21, got", answer)
	}
}

func Test_Part1WithActualData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/actual.txt")
	answer := part1(inputLines)

	if answer != 1798 {
		t.Error("Expected 1798, got", answer)
	}
}

func Test_Part2WithExampleData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/example.txt")
	answer := part2(inputLines)

	if answer != 8 {
		t.Error("Expected 8, got", answer)
	}
}

func Test_Part2WithActualData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/actual.txt")
	answer := part2(inputLines)

	if answer != 259308 {
		t.Error("Expected 259308, got", answer)
	}
}
