package Day11

import (
	"github.com/jeffgbutler/AdventOfCode2022/functions"
	"sort"
	"strconv"
	"strings"
)

type operation struct {
	left, right int
	op          string
}

type monkey struct {
	monkeyNumber                                      int
	items                                             []int
	operation                                         operation
	divisor, trueTarget, falseTarget, inspectionCount int
}

func part1(inputLineGroups [][]string) int {
	var monkeys []monkey

	for _, g := range inputLineGroups {
		monkeys = append(monkeys, parseMonkey(g))
	}

	for i := 0; i < 20; i++ {
		round(monkeys)
	}

	counts := functions.Map(monkeys, func(m monkey) int { return m.inspectionCount })
	sort.Slice(counts, func(i, j int) bool { return counts[i] > counts[j] })
	return counts[0] * counts[1]
}

func round(monkeys []monkey) {
	for i := 0; i < len(monkeys); i++ {
		mky := &monkeys[i]
		for _, item := range mky.items {
			mky.inspectionCount++
			a := performOperation(item, mky.operation)
			a = a / 3
			if a%mky.divisor == 0 {
				throwToMonkey(a, mky.trueTarget, monkeys)
			} else {
				throwToMonkey(a, mky.falseTarget, monkeys)
			}
		}
		mky.items = []int{}
	}
}

func performOperation(old int, op operation) int {
	if op.left == -99999 {
		op.left = old
	}

	if op.right == -99999 {
		op.right = old
	}

	answer := 0
	switch op.op {
	case "*":
		answer = op.left * op.right
	case "+":
		answer = op.left + op.right
	}

	return answer
}

func throwToMonkey(item, targetMonkey int, monkeys []monkey) {
	mky := &monkeys[targetMonkey]
	mky.items = append(mky.items, item)
}

func parseMonkey(inputLines []string) monkey {

	answer := monkey{}

	for _, line := range inputLines {
		switch {
		case strings.HasPrefix(line, "Monkey"):
			answer.monkeyNumber = parseMonkeyLine(line)
		case strings.HasPrefix(line, "  Starting"):
			answer.items = parseStartingItemsLine(line)
		case strings.HasPrefix(line, "  Operation"):
			answer.operation = parseOperation(line)
		case strings.HasPrefix(line, "  Test"):
			answer.divisor = parseDivisorLine(line)
		case strings.HasPrefix(line, "    If true"):
			answer.trueTarget = parseTargetLine(line)
		case strings.HasPrefix(line, "    If false"):
			answer.falseTarget = parseTargetLine(line)
		}
	}

	return answer
}

func parseTargetLine(line string) int {
	f := strings.Fields(line)
	i, _ := strconv.Atoi(f[5])
	return i
}

func parseDivisorLine(line string) int {
	f := strings.Fields(line)
	i, _ := strconv.Atoi(f[3])
	return i
}

func parseStartingItemsLine(line string) []int {
	index := strings.Index(line, ":")
	itemString := line[index+2:]
	f := strings.Split(itemString, ", ")

	var answer []int

	for i := 0; i < len(f); i++ {
		num, _ := strconv.Atoi(f[i])
		answer = append(answer, num)
	}

	return answer
}

func parseMonkeyLine(line string) int {
	startIndex := strings.Index(line, " ")
	endIndex := strings.Index(line, ":")

	f := line[startIndex+1 : endIndex]
	i, _ := strconv.Atoi(f)
	return i
}

func parseOperation(line string) operation {
	startIndex := strings.Index(line, "= ")
	f := strings.Fields(line[startIndex+2:])

	answer := operation{}

	if f[0] == "old" {
		answer.left = -99999
	} else {
		i, _ := strconv.Atoi(f[0])
		answer.left = i
	}

	answer.op = f[1]

	if f[2] == "old" {
		answer.right = -99999
	} else {
		i, _ := strconv.Atoi(f[2])
		answer.right = i
	}

	return answer
}
