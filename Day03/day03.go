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
