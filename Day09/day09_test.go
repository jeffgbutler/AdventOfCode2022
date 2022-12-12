package Day09

import (
	"github.com/jeffgbutler/AdventOfCode2022/inputfile"
	"testing"
)

func Test_Part1WithExampleData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/example.txt")

	answer := part1(inputLines)

	if answer != 13 {
		t.Error("Expected 13, got", answer)
	}
}

func Test_Part1WithActualData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/actual.txt")

	answer := part1(inputLines)

	if answer != 6464 {
		t.Error("Expected 6464, got", answer)
	}
}

func Test_Part2WithExampleData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/example.txt")

	answer := part2(inputLines)

	if answer != 1 {
		t.Error("Expected 1, got", answer)
	}
}

func Test_Part2WithExampleData2(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/example2.txt")

	answer := part2(inputLines)

	if answer != 36 {
		t.Error("Expected 36, got", answer)
	}
}

func Test_Part2WithActualData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/actual.txt")

	answer := part2(inputLines)

	if answer != 2604 {
		t.Error("Expected 2604, got", answer)
	}
}
