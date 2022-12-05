package Day04

import (
	"github.com/jeffgbutler/AdventOfCode2022/inputfile"
	"testing"
)

func Test_Part1WihExampleData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/example.txt")
	answer := part1(inputLines)

	if answer != 2 {
		t.Error("Expected 2, got", answer)
	}
}

func Test_Part1WihActualData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/actual.txt")
	answer := part1(inputLines)

	if answer != 528 {
		t.Error("Expected 528, got", answer)
	}
}

func Test_Part2WihExampleData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/example.txt")
	answer := part2(inputLines)

	if answer != 4 {
		t.Error("Expected 4, got", answer)
	}
}

func Test_Part2WihActualData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/actual.txt")
	answer := part2(inputLines)

	if answer != 881 {
		t.Error("Expected 881, got", answer)
	}
}
