package Day05

import (
	"AdventOfCode2022/datastructures"
	"strconv"
	"strings"
)

type move struct {
	count, from, to int
}

func part1(inputLines []string) string {

	stacks, moves := parseInputFile(inputLines)

	processMovesPart1(stacks, moves)

	var answer string
	for i := 0; i < len(stacks); i++ {
		r, valid := stacks[i].Peek()
		if valid {
			answer += r
		}
	}

	return answer
}

func part2(inputLines []string) string {

	stacks, moves := parseInputFile(inputLines)

	processMovesPart2(stacks, moves)

	var answer string
	for i := 0; i < len(stacks); i++ {
		r, valid := stacks[i].Peek()
		if valid {
			answer += r
		}
	}

	return answer
}

func parseInputFile(inputLines []string) ([]datastructures.Stack[string], []move) {
	var stacks []datastructures.Stack[string]
	stacks, inputLines = parseStackLines(inputLines)
	moves := parseMoveLines(inputLines)

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

func parseStackLines(inputLines []string) ([]datastructures.Stack[string], []string) {
	// save initial stack lines
	var stackLines datastructures.Stack[string]

	for {
		currentLine := inputLines[0]
		if strings.ContainsRune(currentLine, '[') {
			stackLines.Push(currentLine)
			inputLines = inputLines[1:]
		} else {
			break
		}
	}

	// input lines now pointing to the number of stacks
	currentLine := inputLines[0]
	inputLines = inputLines[1:]
	s := strings.Fields(currentLine)
	numberOfStacks, _ := strconv.Atoi(s[len(s)-1])
	stacks := make([]datastructures.Stack[string], numberOfStacks)

	var stackLine string
	var valid bool
	for {
		stackLine, valid = stackLines.Pop()
		if valid {
			populateStack(stackLine, stacks)
		} else {
			break
		}
	}

	return stacks, inputLines[1:]
}

func populateStack(inputLine string, stacks []datastructures.Stack[string]) {
	for i := 0; (i*4)+1 < len(inputLine); i++ {
		offset := (i * 4) + 1
		if inputLine[offset] != ' ' {
			stacks[i].Push(string(inputLine[offset]))
		}
	}
}

func parseMoveLines(inputLines []string) []move {
	var moves []move

	for _, inputLine := range inputLines {
		s := strings.Split(inputLine, " ")
		count, _ := strconv.Atoi(s[1])
		from, _ := strconv.Atoi(s[3])
		to, _ := strconv.Atoi(s[5])
		moves = append(moves, move{count, from, to})
	}

	return moves
}
