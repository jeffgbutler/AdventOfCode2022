package Day11

import (
	"github.com/jeffgbutler/AdventOfCode2022/inputfile"
	"testing"
)

func Test_Part1WithExampleData(t *testing.T) {
	inputLineGroups, _ := inputfile.ReadGroups("testdata/example.txt")

	answer := part1(inputLineGroups)

	if answer != 10605 {
		t.Error("Expected 10605, got", answer)
	}
}

func Test_Part1WithActualData(t *testing.T) {
	inputLineGroups, _ := inputfile.ReadGroups("testdata/actual.txt")

	answer := part1(inputLineGroups)

	if answer != 58786 {
		t.Error("Expected 58786, got", answer)
	}
}

func Test_Part2WithExampleData(t *testing.T) {
	inputLineGroups, _ := inputfile.ReadGroups("testdata/example.txt")

	answer := part2(inputLineGroups)

	if answer != 2713310158 {
		t.Error("Expected 2713310158, got", answer)
	}
}

func Test_Part2WithActualData(t *testing.T) {
	inputLineGroups, _ := inputfile.ReadGroups("testdata/actual.txt")

	answer := part2(inputLineGroups)

	if answer != 14952185856 {
		t.Error("Expected 14952185856, got", answer)
	}
}
