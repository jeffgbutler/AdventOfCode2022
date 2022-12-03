package Day01

import (
	"sort"
	"strconv"
)

type elf struct {
	elfNumber int
	calories  int
}

func convertToElves(lines []string) ([]elf, error) {
	var allElves []elf
	var currentElf elf

	for _, line := range lines {
		if len(line) == 0 {
			if currentElf.calories > 0 {
				allElves = append(allElves, currentElf)
				currentElf = elf{len(allElves) + 1, 0}
			} else {
				continue
			}
		} else {
			num, err := strconv.Atoi(line)
			if err != nil {
				return allElves, err
			}
			currentElf.calories += num
		}
	}

	if currentElf.calories > 0 {
		allElves = append(allElves, currentElf)
	}
	return allElves, nil
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
	total := 0
	for _, e := range elves {
		total += e.calories
	}

	return total
}
