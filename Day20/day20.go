package Day20

import (
	"strconv"
)

type entry struct {
	value, initialIndex int
}

func part1(inputLines []string) int {
	var mixedEntries []entry

	for i, inputLine := range inputLines {
		mixedEntries = append(mixedEntries, entry{value: toInt(inputLine), initialIndex: i})
	}

	for i := 0; i < len(inputLines); i++ {
		mixIt(i, mixedEntries)
	}

	zeroIndex := findByValue(0, mixedEntries)

	oneThousandIndex := (1000 + zeroIndex) % len(mixedEntries)
	twoThousandIndex := (2000 + zeroIndex) % len(mixedEntries)
	threeThousandIndex := (3000 + zeroIndex) % len(mixedEntries)

	oneThousand := mixedEntries[oneThousandIndex]
	twoThousand := mixedEntries[twoThousandIndex]
	threeThousand := mixedEntries[threeThousandIndex]

	return oneThousand.value + twoThousand.value + threeThousand.value
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func mixIt(initialIndex int, entries []entry) {
	currentIndex := findByInitialIndex(initialIndex, entries)
	currentEntry := entries[currentIndex]

	slots := len(entries) - 1
	moves := currentEntry.value % slots

	if moves == 0 {
		return
	}

	if moves < 0 {
		moves = slots + moves
	}

	newIndex := (currentIndex + moves) % slots

	if currentIndex == newIndex {
		return
	}

	// cut out entry at index and shift higher entries down
	entries = append(entries[:currentIndex], entries[currentIndex+1:]...)

	entries = append(entries, entry{0, 0})
	copy(entries[newIndex+1:], entries[newIndex:])
	entries[newIndex] = currentEntry
}

func findByInitialIndex(initialIndex int, entries []entry) int {
	for i := 0; i < len(entries); i++ {
		if entries[i].initialIndex == initialIndex {
			return i
		}
	}

	return -1
}

func findByValue(value int, entries []entry) int {
	for i := 0; i < len(entries); i++ {
		if entries[i].value == value {
			return i
		}
	}

	return -1
}
