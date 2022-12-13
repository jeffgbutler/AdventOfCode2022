package Day12

import (
	"github.com/jeffgbutler/AdventOfCode2022/inputfile"
	"testing"
)

func Test_Part1WithExampleData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/example.txt")

	answer := part1(inputLines)

	if answer != 31 {
		t.Error("Expected 31, got", answer)
	}
}

func Test_Part1WithActualData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/actual.txt")

	answer := part1(inputLines)

	if answer != 497 {
		t.Error("Expected 497, got", answer)
	}
}

func Test_Part2WithExampleData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/example.txt")

	answer := part2(inputLines)

	if answer != 29 {
		t.Error("Expected 29, got", answer)
	}
}

func Test_Part2WithActualData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/actual.txt")

	answer := part2(inputLines)

	if answer != 492 {
		t.Error("Expected 492, got", answer)
	}
}
