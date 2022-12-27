package Day21

import (
	"strconv"
	"strings"
)

type node interface {
	value(nodes map[string]node) int
}

type integerNode struct {
	val int
}

func (n integerNode) value(_ map[string]node) int {
	return n.val
}

type mathNode struct {
	operation, left, right string
}

func (n mathNode) value(nodes map[string]node) int {
	leftValue := nodes[n.left].value(nodes)
	rightValue := nodes[n.right].value(nodes)

	switch n.operation {
	case "+":
		return leftValue + rightValue
	case "-":
		return leftValue - rightValue
	case "*":
		return leftValue * rightValue
	case "/":
		return leftValue / rightValue
	default:
		return -1
	}
}

func part1(inputLines []string) int {
	nodes := parse(inputLines)

	rootNode := nodes["root"]

	return rootNode.value(nodes)
}

func parse(inputLines []string) map[string]node {
	nodes := map[string]node{}

	for _, inputLine := range inputLines {
		name, node := parseInputLine(inputLine)
		nodes[name] = node
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
