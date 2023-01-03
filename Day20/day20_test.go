package Day20

import (
	"github.com/jeffgbutler/AdventOfCode2022/inputfile"
	"testing"
)

func Test_Part1WithExampleData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/example.txt")

	answer := part1(inputLines)

	if answer != 3 {
		t.Error("Expecting 3, got", answer)
	}
}

func Test_Part1WithActualData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/actual.txt")

	answer := part1(inputLines)

	if answer != 15297 {
		t.Error("Expecting 15297, got", answer)
	}
}
