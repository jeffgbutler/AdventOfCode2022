package Day02

import (
	"AdventOfCode2022/inputfile"
	"testing"
)

func Test_Part1WithExampleData(t *testing.T) {

	inputLines, _ := inputfile.Read("testdata/example.txt")
	rounds := convertToRoundsPart1(inputLines)
	total := calculateMatchOutcome(rounds)

	if total != 15 {
		t.Error("Expected total of 15, got", total)
	}
}

func Test_Part1WithRealData(t *testing.T) {

	inputLines, _ := inputfile.Read("testdata/mydata.txt")
	rounds := convertToRoundsPart1(inputLines)
	total := calculateMatchOutcome(rounds)

	if total != 14297 {
		t.Error("Expected total of 14297, got", total)
	}
}

func Test_Part2WithExampleData(t *testing.T) {

	inputLines, _ := inputfile.Read("testdata/example.txt")
	rounds := convertToRoundsPart2(inputLines)
	total := calculateMatchOutcome(rounds)

	if total != 12 {
		t.Error("Expected total of 12, got", total)
	}
}

func Test_Part2WithRealData(t *testing.T) {

	inputLines, _ := inputfile.Read("testdata/mydata.txt")
	rounds := convertToRoundsPart2(inputLines)
	total := calculateMatchOutcome(rounds)

	if total != 10498 {
		t.Error("Expected total of 10498, got", total)
	}
}
