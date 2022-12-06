package Day01

import (
	"github.com/jeffgbutler/AdventOfCode2022/inputfile"
	"testing"
)

func Test_Converter(t *testing.T) {
	groups, err := inputfile.ReadGroups("testdata/example.txt")

	if err != nil {
		t.Fatal(err)
	}

	elves, err := convertToElves(groups)

	if err != nil {
		t.Fatal(err)
	}

	if len(elves) != 5 {
		t.Error("Expected 5 elves, got", len(elves))
	}
}

func Test_findGreatestElfWithExampleData(t *testing.T) {
	groups, err := inputfile.ReadGroups("testdata/example.txt")

	if err != nil {
		t.Fatal(err)
	}

	elves, err := convertToElves(groups)

	if err != nil {
		t.Fatal(err)
	}

	greatestElf := findGreatestElf(elves)
	expectedElf := elf{4, 24000}
	if greatestElf != expectedElf {
		t.Error("Expected elf", expectedElf, "got elf", greatestElf)
	}
}

func Test_findGreatestElfWithActualData(t *testing.T) {
	groups, err := inputfile.ReadGroups("testdata/actual.txt")

	if err != nil {
		t.Fatal(err)
	}

	elves, err := convertToElves(groups)

	if err != nil {
		t.Fatal(err)
	}

	greatestElf := findGreatestElf(elves)
	expectedElf := elf{172, 70509}
	if greatestElf != expectedElf {
		t.Error("Expected elf", expectedElf, "got elf", greatestElf)
	}
}

func Test_findTopThreeElvesWithExampleData(t *testing.T) {
	groups, err := inputfile.ReadGroups("testdata/example.txt")

	if err != nil {
		t.Fatal(err)
	}

	elves, err := convertToElves(groups)

	if err != nil {
		t.Fatal(err)
	}

	topThreeElves := findTopThreeElves(elves)
	expectedTopThree := []elf{{4, 24000}, {3, 11000}, {5, 10000}}

	if !compareElfSlices(topThreeElves, expectedTopThree) {
		t.Error("Expected top three elves", expectedTopThree, "got", topThreeElves)
	}

	totalCalories := calculateTotalCalories(topThreeElves)
	expectedCalories := 45000
	if totalCalories != expectedCalories {
		t.Error("Expected top three elves to have", expectedCalories, "calories, got", totalCalories)
	}
}

func Test_findTopThreeElvesWithActualData(t *testing.T) {
	groups, err := inputfile.ReadGroups("testdata/actual.txt")

	if err != nil {
		t.Fatal(err)
	}

	elves, err := convertToElves(groups)

	if err != nil {
		t.Fatal(err)
	}

	topThreeElves := findTopThreeElves(elves)
	expectedTopThree := []elf{{172, 70509}, {80, 69894}, {28, 68164}}

	if !compareElfSlices(topThreeElves, expectedTopThree) {
		t.Error("Expected top three elves", expectedTopThree, "got", topThreeElves)
	}

	totalCalories := calculateTotalCalories(topThreeElves)
	expectedCalories := 208567
	if totalCalories != expectedCalories {
		t.Error("Expected top three elves to have", expectedCalories, "calories, got", totalCalories)
	}
}
