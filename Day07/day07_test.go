package Day07

import (
	"github.com/jeffgbutler/AdventOfCode2022/inputfile"
	"testing"
)

func Test_Part1WithSampleData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/example.txt")
	res := part1(inputLines)

	if res != 95437 {
		t.Error("Expected 95437, got", res)
	}
}

func Test_Part1WithActualData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/actual.txt")
	res := part1(inputLines)

	if res != 1307902 {
		t.Error("Expected 1307902, got", res)
	}
}

func Test_Part2WithSampleData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/example.txt")
	res := part2(inputLines)

	if res != 24933642 {
		t.Error("Expected 24933642, got", res)
	}
}

func Test_Part2WithActualData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/actual.txt")
	res := part2(inputLines)

	if res != 7068748 {
		t.Error("Expected 7068748, got", res)
	}
}
