package Day05

import (
	"github.com/jeffgbutler/AdventOfCode2022/datastructures"
	"github.com/jeffgbutler/AdventOfCode2022/functions"
	"strconv"
	"strings"
)

type move struct {
	count, from, to int
}

func part1(inputLineGroups [][]string) string {
	stacks, moves := parseInputFile(inputLineGroups)
	processMovesPart1(stacks, moves)
	return buildAnswer(stacks)
}

func part2(inputLineGroups [][]string) string {
	stacks, moves := parseInputFile(inputLineGroups)
	processMovesPart2(stacks, moves)
	return buildAnswer(stacks)
}

func parseInputFile(inputLineGroups [][]string) ([]datastructures.Stack[string], []move) {
	stacks := parseStackLines(inputLineGroups[0])
	moves := parseMoveLines(inputLineGroups[1])

	return stacks, moves
}

func processMovesPart1(stacks []datastructures.Stack[string], moves []move) {
	for _, theMove := range moves {
		for i := 0; i < theMove.count; i++ {
			s, _ := stacks[theMove.from-1].Pop()
			stacks[theMove.to-1].Push(s)
		}
	}
}

func processMovesPart2(stacks []datastructures.Stack[string], moves []move) {
	for _, theMove := range moves {
		values, _ := stacks[theMove.from-1].PopMultiple(theMove.count)
		stacks[theMove.to-1].PushMultiple(values)
	}
}

func parseStackLines(inputLines []string) []datastructures.Stack[string] {
	// calculate number of stacks
	s := strings.Fields(inputLines[len(inputLines)-1])
	numberOfStacks, _ := strconv.Atoi(s[len(s)-1])
	stacks := make([]datastructures.Stack[string], numberOfStacks)

	// process stack information backwards
	for i := len(inputLines) - 2; i >= 0; i-- {
		populateStack(inputLines[i], stacks)
	}

	return stacks
}

func populateStack(inputLine string, stacks []datastructures.Stack[string]) {
	// this will fail if there are more than 9 stacks
	for i := 0; (i*4)+1 < len(inputLine); i++ {
		offset := (i * 4) + 1
		if inputLine[offset] != ' ' {
			stacks[i].Push(string(inputLine[offset]))
		}
	}
}

func parseMoveLines(inputLines []string) []move {
	return functions.Map(inputLines, convertLineToMove)
}

func convertLineToMove(line string) move {
	s := strings.Split(line, " ")
	count, _ := strconv.Atoi(s[1])
	from, _ := strconv.Atoi(s[3])
	to, _ := strconv.Atoi(s[5])
	return move{count, from, to}
}

func buildAnswer(stacks []datastructures.Stack[string]) string {
	// build a string from the top element of each stack
	return functions.Reduce(stacks, "", func(s string, stack datastructures.Stack[string]) string {
		value, valid := stack.Peek()
		if valid {
			s += value
		}
		return s
	})
}
