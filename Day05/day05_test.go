package Day05

import (
	"github.com/jeffgbutler/AdventOfCode2022/inputfile"
	"testing"
)

func Test_Part1WithExampleData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/example.txt")
	s := part1(inputLines)

	if s != "CMZ" {
		t.Error("Expected CMZ, got", s)
	}
}

func Test_Part1WithActualData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/actual.txt")
	s := part1(inputLines)

	if s != "MQSHJMWNH" {
		t.Error("Expected MQSHJMWNH, got", s)
	}
}

func Test_Part2WithExampleData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/example.txt")
	s := part2(inputLines)

	if s != "MCD" {
		t.Error("Expected MCD, got", s)
	}
}

func Test_Part2WithActualData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/actual.txt")
	s := part2(inputLines)

	if s != "LLWJRBHVZ" {
		t.Error("Expected LLWJRBHVZ, got", s)
	}
}
