package Day22

import (
	"github.com/jeffgbutler/AdventOfCode2022/inputfile"
	"testing"
)

func Test_Part1WithExampleData(t *testing.T) {
	inputGroups, _ := inputfile.ReadGroups("testdata/example.txt")

	answer := part1(inputGroups)

	if answer != 6032 {
		t.Error("Expected 6032, got", answer)
	}
}

func Test_Part1WithActualData(t *testing.T) {
	inputGroups, _ := inputfile.ReadGroups("testdata/actual.txt")

	answer := part1(inputGroups)

	if answer != 36518 {
		t.Error("Expected 36518, got", answer)
	}
}
