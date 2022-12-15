package Day13

import (
	"sort"
	"strconv"
	"strings"
)

type packet interface {
	compareTo(other packet) int
	isInt() bool
}

type ip int

func (i ip) compareTo(other packet) int {
	if other.isInt() {
		o := other.(ip)
		switch {
		case int(i) == int(o):
			return 0
		case int(i) < int(o):
			return -1
		default:
			return 1
		}
	} else {
		v2 := lp{i}
		return v2.compareTo(other)
	}
}

func (i ip) isInt() bool {
	return true
}

type lp []packet

func (l lp) compareTo(other packet) int {
	if other.isInt() {
		return l.compareTo(lp{other})
	} else {
		return compareTo(l, other.(lp))
	}
}

func (l lp) isInt() bool {
	return false
}

func compareTo(left, right lp) int {
	for i := 0; i < len(left) && i < len(right); i++ {
		v := left[i].compareTo(right[i])
		if v == 0 {
			continue
		} else {
			return v
		}
	}

	switch {
	case len(left) == len(right):
		return 0
	case len(left) < len(right):
		return -1
	default:
		return 1
	}
}

func part1(inputLineGroups [][]string) int {

	count := 0
	lps := parse(inputLineGroups)
	for i, v := range lps {
		if v[0].compareTo(v[1]) <= 0 {
			count += i + 1
		}
	}

	return count
}

func part2(inputLineGroups [][]string) int {

	lps := parse(inputLineGroups)
	//flatten the list of packets
	var allPackets []lp
	for _, v := range lps {
		allPackets = append(allPackets, v...)
	}

	// sort the packets
	sort.Slice(allPackets, func(i, j int) bool { return allPackets[i].compareTo(allPackets[j]) <= 0 })

	startDivider := lp{lp{ip(2)}}
	endDivider := lp{lp{ip(6)}}

	startIndex := 0
	endIndex := 0

	for i, v := range allPackets {
		if startIndex == 0 && startDivider.compareTo(v) <= 0 {
			startIndex = i + 1
		}

		if startIndex != 0 && endIndex == 0 && endDivider.compareTo(v) <= 0 {
			endIndex = i + 2 // account for startDivider position
			break
		}
	}

	return startIndex * endIndex
}

func parse(inputLineGroups [][]string) [][]lp {

	var answer [][]lp

	for _, v := range inputLineGroups {
		answer = append(answer, parseInputLines(v))
	}

	return answer
}

func parseInputLines(inputLines []string) []lp {
	if len(inputLines) != 2 {
		panic("invalid input")
	}

	var answer []lp
	answer = append(answer, parseInputLine(inputLines[0]))
	answer = append(answer, parseInputLine(inputLines[1]))
	return answer
}

func parseInputLine(inputLine string) lp {
	var answer lp

	if !strings.Contains(inputLine, "[") {
		// no sub lists
		answer = append(answer, parseListOfIntegers(inputLine)...)
		return answer
	}

	startIndex := 0
	endIndex := -1
	for startIndex < len(inputLine) {
		// find index of string in slice
		nextOpenBracketIndex := strings.Index(inputLine[startIndex:], "[")
		if nextOpenBracketIndex == -1 {
			// no more sub-lists
			break
		}
		nextOpenBracketIndex += startIndex // add the start index so we're now pointing to the index in the full string

		if nextOpenBracketIndex > startIndex {
			// some list of integers before the first open bracket
			answer = append(answer, parseListOfIntegers(inputLine[startIndex:nextOpenBracketIndex])...)
		}

		openBrackets := 0
		endIndex = -1
	forLoop:
		for i := nextOpenBracketIndex; i < len(inputLine); i++ {
			switch inputLine[i] {
			case '[':
				openBrackets++
			case ']':
				openBrackets--
				if openBrackets == 0 {
					endIndex = i
					break forLoop
				}
			}
		}

		if endIndex == -1 {
			panic("unmatched brackets in string " + inputLine)
		}

		answer = append(answer, parseInputLine(inputLine[nextOpenBracketIndex+1:endIndex]))
		endIndex++
		startIndex = endIndex
	}

	if endIndex < len(inputLine) {
		answer = append(answer, parseListOfIntegers(inputLine[endIndex:])...)
	}

	return answer
}

func parseListOfIntegers(s string) []packet {
	if strings.Contains(s, "[") {
		panic("parseListOfInts received a string with sub lists")
	}

	var answer []packet

	intStrings := strings.Split(s, ",")
	for _, intString := range intStrings {
		theInt, err := strconv.Atoi(intString)
		if err == nil {
			answer = append(answer, ip(theInt))
		}
	}

	return answer
}
