package Day06

import "strings"

func part1(s string) int {
	return findUniqueString(s, 4)
}

func part2(s string) int {
	return findUniqueString(s, 14)
}

func findUniqueString(s string, length int) int {
	foundIt := -1

outer:
	for i := 0; i < len(s)-length; i++ {
		subString := s[i : i+length]
		for j := 0; j < length; j++ {
			count := strings.Count(subString, string(s[i+j]))
			if count > 1 {
				continue outer
			}
		}
		foundIt = i
		break
	}

	return foundIt + length
}
