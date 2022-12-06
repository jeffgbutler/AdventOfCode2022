package Day06

import (
	"github.com/jeffgbutler/AdventOfCode2022/inputfile"
	"testing"
)

func Test_Part1WithExampleData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/example.txt")
	expected := []int{7, 5, 6, 10, 11}

	if len(inputLines) != len(expected) {
		t.Error("Part1, Example data, inputs don't match")
	}

	for i, v := range inputLines {
		answer := part1(v)
		if answer != expected[i] {
			t.Error("expected", expected[i], "got", answer)
		}
	}
}

func Test_Part1WithActualData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/actual.txt")
	expected := []int{1034}

	if len(inputLines) != len(expected) {
		t.Error("Part1, Actual data, inputs don't match")
	}

	for i, v := range inputLines {
		answer := part1(v)
		if answer != expected[i] {
			t.Error("expected", expected[i], "got", answer)
		}
	}
}

func Test_Part2WithExampleData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/example.txt")
	expected := []int{19, 23, 23, 29, 26}

	if len(inputLines) != len(expected) {
		t.Error("Part2, Example data, inputs don't match")
	}

	for i, v := range inputLines {
		answer := part2(v)
		if answer != expected[i] {
			t.Error("expected", expected[i], "got", answer)
		}
	}
}

func Test_Part2WithActualData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/actual.txt")
	expected := []int{2472}

	if len(inputLines) != len(expected) {
		t.Error("Part2, Actual data, inputs don't match")
	}

	for i, v := range inputLines {
		answer := part2(v)
		if answer != expected[i] {
			t.Error("expected", expected[i], "got", answer)
		}
	}
}
