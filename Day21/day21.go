package Day21

import (
	"strconv"
	"strings"
)

type node interface {
	value(nodes map[string]node) (int, bool)
	calculateMissingValue(nodes map[string]node, expectedValue int) int
}

type integerNode struct {
	val int
}

func (n integerNode) value(_ map[string]node) (int, bool) {
	return n.val, true
}

func (n integerNode) calculateMissingValue(_ map[string]node, _ int) int {
	// should never be called
	panic("panic 2")
}

type humanNode struct{}

func (n humanNode) value(_ map[string]node) (int, bool) {
	return 0, false
}

func (n humanNode) calculateMissingValue(_ map[string]node, expectedValue int) int {
	return expectedValue
}

type mathNode struct {
	operation, left, right string
}

func (n mathNode) value(nodes map[string]node) (int, bool) {
	leftValue, okl := nodes[n.left].value(nodes)
	rightValue, okr := nodes[n.right].value(nodes)

	if !okl || !okr {
		return 0, false
	}

	switch n.operation {
	case "+":
		return leftValue + rightValue, true
	case "-":
		return leftValue - rightValue, true
	case "*":
		return leftValue * rightValue, true
	case "/":
		return leftValue / rightValue, true
	default:
		return -1, false
	}
}

func (n mathNode) calculateMissingValue(nodes map[string]node, expectedValue int) int {
	leftValue, leftOk := nodes[n.left].value(nodes)
	rightValue, rightOk := nodes[n.right].value(nodes)

	if !leftOk {
		switch n.operation {
		case "+":
			// l + r = e ---> l = e - r
			return nodes[n.left].calculateMissingValue(nodes, expectedValue-rightValue)
		case "-":
			// l - r = e ---> l = e + r
			return nodes[n.left].calculateMissingValue(nodes, expectedValue+rightValue)
		case "*":
			// l * r = e ---> l = e / r
			return nodes[n.left].calculateMissingValue(nodes, expectedValue/rightValue)
		case "/":
			// l / r = e ---> l = e * r
			return nodes[n.left].calculateMissingValue(nodes, expectedValue*rightValue)
		}
	}

	if !rightOk {
		switch n.operation {
		case "+":
			// l + r = e ---> r = e - l
			return nodes[n.right].calculateMissingValue(nodes, expectedValue-leftValue)
		case "-":
			// l - r = e ---> l = e + r ---> l - e = r
			return nodes[n.right].calculateMissingValue(nodes, leftValue-expectedValue)
		case "*":
			// l * r = e ---> r = e / l
			return nodes[n.right].calculateMissingValue(nodes, expectedValue/leftValue)
		case "/":
			// l / r = e ---> l = e * r ---> l / e = r
			return nodes[n.right].calculateMissingValue(nodes, leftValue/expectedValue)
		}
	}

	// should never happen
	panic("panic 1")
}

func part1(inputLines []string) int {
	nodes := parse(inputLines)

	rootNode := nodes["root"]

	answer, _ := rootNode.value(nodes)
	return answer
}

func part2(inputLines []string) int {
	nodes := parsePart2(inputLines)

	rootNode := nodes["root"]

	answer := rootNode.calculateMissingValue(nodes, 0)

	return answer
}

func parse(inputLines []string) map[string]node {
	nodes := map[string]node{}

	for _, inputLine := range inputLines {
		name, node := parseInputLine(inputLine)
		nodes[name] = node
	}

	return nodes
}

func parsePart2(inputLines []string) map[string]node {
	nodes := map[string]node{}

	for _, inputLine := range inputLines {
		switch {
		case strings.HasPrefix(inputLine, "humn"):
			nodes["humn"] = humanNode{}
		case strings.HasPrefix(inputLine, "root"):
			name, node := parsePart2Root(inputLine)
			nodes[name] = node
		default:
			name, node := parseInputLine(inputLine)
			nodes[name] = node
		}
	}

	return nodes
}

func parseInputLine(inputLine string) (string, node) {
	p1 := strings.Split(inputLine, ": ")
	p2 := strings.Split(p1[1], " ")

	if len(p2) == 1 {
		// integer node
		i, _ := strconv.Atoi(p2[0])
		return p1[0], integerNode{val: i}
	} else {
		//math node
		return p1[0], mathNode{left: p2[0], operation: p2[1], right: p2[2]}
	}
}

func parsePart2Root(inputLine string) (string, node) {
	p1 := strings.Split(inputLine, ": ")
	p2 := strings.Split(p1[1], " ")

	if len(p2) == 1 {
		// integer node
		i, _ := strconv.Atoi(p2[0])
		return p1[0], integerNode{val: i}
	} else {
		//math node
		return p1[0], mathNode{left: p2[0], operation: "-", right: p2[2]}
	}
}
