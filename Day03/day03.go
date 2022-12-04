package Day03

import "strings"

type rucksackContents struct {
	compartment1 string
	compartment2 string
}

func convertToRucksacks(inputLines []string) []rucksackContents {
	var rucksacks []rucksackContents

	for _, inputLine := range inputLines {
		length := len(inputLine)
		mid := length / 2

		rucksack := rucksackContents{inputLine[:mid], inputLine[mid:length]}
		rucksacks = append(rucksacks, rucksack)
	}

	return rucksacks
}

func calculateRucksackTotal(rucksacks []rucksackContents) int {
	value := 0

	for _, rucksack := range rucksacks {
		value += calculateRucksack(rucksack)
	}

	return value
}

func calculateRucksack(rucksack rucksackContents) int {
	for _, r := range rucksack.compartment1 {
		if strings.ContainsRune(rucksack.compartment2, r) {
			return calculateLetterValue(r)
		}
	}

	return -1
}

func calculateLetterValue(r rune) int {
	val := int(r)
	if val > 90 {
		//lower case
		return val - int('a') + 1
	} else {
		// upper case
		return val - int('A') + 27
	}
}

func calculateGroupPriorityTotal(inputLines []string) (int, int) {
	total := 0
	groups := 0
	for {
		if len(inputLines) < 3 {
			break
		}
		groups++

		total += calculateGroupPriority(inputLines[0], inputLines[1], inputLines[2])
		inputLines = inputLines[3:]
	}

	return groups, total
}

func calculateGroupPriority(line1, line2, line3 string) int {
	for _, v := range line1 {
		if strings.ContainsRune(line2, v) && strings.ContainsRune(line3, v) {
			return calculateLetterValue(v)
		}
	}

	return -1
}
