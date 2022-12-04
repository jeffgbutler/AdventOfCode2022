package Day03

import (
	"AdventOfCode2022/inputfile"
	"testing"
)

func Test_Part1ExampleData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/example.txt")
	rucksacks := convertToRucksacks(inputLines)
	total := calculateRucksackTotal(rucksacks)

	if total != 157 {
		t.Error("Expected 157, got", total)
	}
}

func Test_Part1ActualData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/actual.txt")
	rucksacks := convertToRucksacks(inputLines)
	total := calculateRucksackTotal(rucksacks)

	if total != 8085 {
		t.Error("Expected 8085, got", total)
	}
}

func Test_Part2ExampleData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/example.txt")
	groups, total := calculateGroupPriorityTotal(inputLines)

	if groups != 2 {
		t.Error("Expected 2 groups, got", groups)
	}

	if total != 70 {
		t.Error("Expected 70, got", total)
	}
}

func Test_Part2ActualData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/actual.txt")
	groups, total := calculateGroupPriorityTotal(inputLines)

	if groups != 100 {
		t.Error("Expected 100 groups, got", groups)
	}

	if total != 2515 {
		t.Error("Expected 2515, got", total)
	}
}

func Test_CalculateValue_a(t *testing.T) {
	value := calculateLetterValue('a')

	if value != 1 {
		t.Error("Expected value of a to be 1, got", value)
	}
}

func Test_CalculateValue_z(t *testing.T) {
	value := calculateLetterValue('z')

	if value != 26 {
		t.Error("Expected value of z to be 26, got", value)
	}
}

func Test_CalculateValue_A(t *testing.T) {
	value := calculateLetterValue('A')

	if value != 27 {
		t.Error("Expected value of A to be 27, got", value)
	}
}

func Test_CalculateValue_Z(t *testing.T) {
	value := calculateLetterValue('Z')

	if value != 52 {
		t.Error("Expected value of Z to be 52, got", value)
	}
}
