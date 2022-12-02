package Day02

import "testing"

func Test_Part1WithExampleData(t *testing.T) {

	rounds, _ := readInputFilePart1("testdata/example.txt")

	total, _ := calculateMatchOutcome(rounds)

	if total != 15 {
		t.Error("Expected total of 15, got", total)
	}
}

func Test_Part1WithRealData(t *testing.T) {

	rounds, _ := readInputFilePart1("testdata/mydata.txt")

	total, _ := calculateMatchOutcome(rounds)

	if total != 14297 {
		t.Error("Expected total of 14297, got", total)
	}
}

func Test_Part2WithExampleData(t *testing.T) {

	rounds, _ := readInputFilePart2("testdata/example.txt")

	total, _ := calculateMatchOutcome(rounds)

	if total != 12 {
		t.Error("Expected total of 12, got", total)
	}
}

func Test_Part2WithRealData(t *testing.T) {

	rounds, _ := readInputFilePart2("testdata/mydata.txt")

	total, _ := calculateMatchOutcome(rounds)

	if total != 10498 {
		t.Error("Expected total of 10498, got", total)
	}
}
