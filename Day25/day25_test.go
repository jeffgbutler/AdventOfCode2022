package Day25

import (
	"github.com/jeffgbutler/AdventOfCode2022/inputfile"
	"testing"
)

func Test_AddThreeDigits(t *testing.T) {
	r := snafuAddThreeDigits("2", "2", "2")
	// 6 --> 1, 1
	if r.result != "1" {
		t.Error("Result is wrong")
	}
	if r.carry != "1" {
		t.Error("Carry is wrong")
	}
}

func Test_AddThreeDigits2(t *testing.T) {
	r := snafuAddThreeDigits("2", "=", "-")
	// -1 --> 0, -
	if r.result != "-" {
		t.Error("Result is wrong")
	}
	if r.carry != "0" {
		t.Error("Carry is wrong")
	}
}

func Test_SnafuAdd(t *testing.T) {
	s1 := "1=-0-2"
	s2 := "12111"

	r := snafuAdd(s1, s2)

	if r != "1-111=" {
		t.Error("Expected 1-111=, got", r)
	}
}

func Test_Part1WithExampleData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/example.txt")

	answer := part1(inputLines)

	if answer != "2=-1=0" {
		t.Error("Expected 2=-1=0, got", answer)
	}
}

func Test_Part1WithActualData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/actual.txt")

	answer := part1(inputLines)

	if answer != "122-0==-=211==-2-200" {
		t.Error("Expected 122-0==-=211==-2-200, got", answer)
	}
}
