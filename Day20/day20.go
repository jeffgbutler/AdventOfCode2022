package Day20

import (
	"strconv"
)

type s1 struct {
	value, initialOrder int
}

func part1(inputLines []string) int {
	var mixedOrder []s1

	for i, inputLine := range inputLines {
		mixedOrder = append(mixedOrder, s1{value: toInt(inputLine), initialOrder: i})
	}

	for i := 0; i < len(inputLines); i++ {
		mixIt(i, mixedOrder)
	}

	zeroIndex := indexOf(0, mixedOrder)

	oneThousandIndex := (1000 + zeroIndex) % len(mixedOrder)
	twoThousandIndex := (2000 + zeroIndex) % len(mixedOrder)
	threeThousandIndex := (3000 + zeroIndex) % len(mixedOrder)

	oneThousand := mixedOrder[oneThousandIndex]
	twoThousand := mixedOrder[twoThousandIndex]
	threeThousand := mixedOrder[threeThousandIndex]

	return oneThousand.value + twoThousand.value + threeThousand.value
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func mixIt(initialIndex int, ints []s1) {
	currentIndex := findIt(initialIndex, ints)
	currentValue := ints[currentIndex]
	val := currentValue.value

	slots := len(ints) - 1
	moves := val % slots

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

	// cut out val at index and shift higher values down
	ints = append(ints[:currentIndex], ints[currentIndex+1:]...)

	// insert val at index
	if newIndex == 0 && val < 0 {
		ints = append(ints, currentValue)
	} else {
		ints = append(ints, s1{0, 0})
		copy(ints[newIndex+1:], ints[newIndex:])
		ints[newIndex] = currentValue
	}
}

func findIt(initialIndex int, ints []s1) int {
	for i := 0; i < len(ints); i++ {
		if ints[i].initialOrder == initialIndex {
			return i
		}
	}

	return -1
}

func indexOf(value int, intSlice []s1) int {
	for index := 0; index < len(intSlice); index++ {
		if intSlice[index].value == value {
			return index
		}
	}

	return -1
}
