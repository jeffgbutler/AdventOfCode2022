package Day13

import (
	"github.com/jeffgbutler/AdventOfCode2022/inputfile"
	"testing"
)

func Test_Part1TreesWithExampleData(t *testing.T) {
	inputLineGroups, _ := inputfile.ReadGroups("testdata/example.txt")

	answer := part1WithTrees(inputLineGroups)

	if answer != 13 {
		t.Error("expected 13, got", answer)
	}
}

func Test_Part1TreesWithActualData(t *testing.T) {
	inputLineGroups, _ := inputfile.ReadGroups("testdata/actual.txt")

	answer := part1WithTrees(inputLineGroups)

	if answer != 6484 {
		t.Error("expected 6484, got", answer)
	}
}

func Test_Part2TreesWithExampleData(t *testing.T) {
	inputLineGroups, _ := inputfile.ReadGroups("testdata/example.txt")

	answer := part2WithTrees(inputLineGroups)

	if answer != 140 {
		t.Error("expected 140, got", answer)
	}
}

func Test_Part2TreesWithActualData(t *testing.T) {
	inputLineGroups, _ := inputfile.ReadGroups("testdata/actual.txt")

	answer := part2WithTrees(inputLineGroups)

	if answer != 19305 {
		t.Error("expected 19305, got", answer)
	}
}

func Test_ToString1(t *testing.T) {
	input := "[]"

	parsedTree := parseInputLineToTree(input)

	output := toString(parsedTree)

	if input != output {
		t.Error("Parse/ToString mismatch.")
	}
}

func Test_ToString2(t *testing.T) {
	startDivider := newStartDivider()
	endDivider := newEndDivider()

	sst := toString(startDivider)
	est := toString(endDivider)

	if sst != "[[2]]" {
		t.Error("Expected [[2]], got", sst)
	}

	if est != "[[6]]" {
		t.Error("Expected [[6]], got", sst)
	}
}
