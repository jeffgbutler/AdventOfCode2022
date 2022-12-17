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

func Test_Part1WithActualData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/actual.txt")

	answer := part1(inputLines, 2_000_000)

	if answer != 5511201 {
		t.Error("Expected 5,511,201, got", answer)
	}
}
