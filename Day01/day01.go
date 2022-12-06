package Day01

import (
	"github.com/jeffgbutler/AdventOfCode2022/functions"
	"sort"
	"strconv"
)

type elf struct {
	elfNumber int
	calories  int
}

func convertToElves(lineGroups [][]string) ([]elf, error) {
	var allElves []elf

	for i, group := range lineGroups {
		total := functions.Reduce(group, 0, addSingleLineValue)
		currentElf := elf{i + 1, total}
		allElves = append(allElves, currentElf)
	}

	return allElves, nil
}

func addSingleLineValue(i int, line string) int {
	num, _ := strconv.Atoi(line)
	return i + num
}

func findGreatestElf(elves []elf) elf {
	sortElvesByCaloriesMostToLeast(elves)
	return elves[0]
}

func findTopThreeElves(elves []elf) []elf {
	sortElvesByCaloriesMostToLeast(elves)
	return elves[:3]
}

func sortElvesByCaloriesMostToLeast(elves []elf) {
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].calories > elves[j].calories
	})
}

func compareElfSlices(a, b []elf) bool {
	if len(a) != len(b) {
		return false
	}

	for i, e := range a {
		if e != b[i] {
			return false
		}
	}

	return true
}

func calculateTotalCalories(elves []elf) int {
	cals := functions.Map(elves, func(e elf) int { return e.calories })
	return functions.Reduce(cals, 0, func(i1, i2 int) int { return i1 + i2 })
}
