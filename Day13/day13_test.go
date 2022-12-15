package Day13

import (
	"github.com/jeffgbutler/AdventOfCode2022/inputfile"
	"testing"
)

func Test_Part1WithExampleData(t *testing.T) {
	inputLineGroups, _ := inputfile.ReadGroups("testdata/example.txt")

	answer := part1(inputLineGroups)

	if answer != 13 {
		t.Error("expected 13, got", answer)
	}
}

func Test_Part1WithActualData(t *testing.T) {
	inputLineGroups, _ := inputfile.ReadGroups("testdata/actual.txt")

	answer := part1(inputLineGroups)

	if answer != 6484 {
		t.Error("expected 6484, got", answer)
	}
}

func Test_Part2WithExampleData(t *testing.T) {
	inputLineGroups, _ := inputfile.ReadGroups("testdata/example.txt")

	answer := part2(inputLineGroups)

	if answer != 140 {
		t.Error("expected 140, got", answer)
	}
}

func Test_Part2WithActualData(t *testing.T) {
	inputLineGroups, _ := inputfile.ReadGroups("testdata/actual.txt")

	answer := part2(inputLineGroups)

	if answer != 19305 {
		t.Error("expected 19305, got", answer)
	}
}
