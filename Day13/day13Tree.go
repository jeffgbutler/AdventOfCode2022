package Day13

import (
	"sort"
	"strconv"
)

type tree struct {
	value    int
	children []*tree
	parent   *tree
}

func compareTrees(left, right tree) int {
	switch {
	case left.value >= 0 && right.value >= 0:
		// both are ints
		switch {
		case left.value == right.value:
			return 0
		case left.value < right.value:
			return -1
		default:
			return 1
		}
	case left.value >= 0:
		// left is an int, right is a tree - convert left to a tree
		return compareTrees(tree{value: -1, children: []*tree{&left}}, right)
	case right.value >= 0:
		// left is a tree, right is an int - convert right to a tree
		return compareTrees(left, tree{value: -1, children: []*tree{&right}})
	default:
		// both are trees
		for i := 0; i < len(left.children) && i < len(right.children); i++ {
			v := compareTrees(*left.children[i], *right.children[i])
			if v == 0 {
				continue
			} else {
				return v
			}
		}

		switch {
		case len(left.children) == len(right.children):
			return 0
		case len(left.children) < len(right.children):
			return -1
		default:
			return 1
		}
	}
}

func toString(tr tree) string {
	if tr.value >= 0 {
		return strconv.Itoa(tr.value)
	}

	if len(tr.children) == 0 {
		return ""
	}

	answer := "["
	first := true
	for _, t := range tr.children {
		if first {
			first = false
		} else {
			answer += ","
		}
		answer += toString(*t)
	}

	answer += "]"

	return answer
}

func part1WithTrees(inputLineGroups [][]string) int {
	treeGroups := parseToTrees(inputLineGroups)
	count := 0

	for i, treeGroup := range treeGroups {
		if compareTrees(treeGroup[0], treeGroup[1]) <= 0 {
			count += i + 1
		}
	}

	return count
}

func part2WithTrees(inputLineGroups [][]string) int {
	treeGroups := parseToTrees(inputLineGroups)
	//flatten the list of trees
	var allTrees []tree
	for _, treeGroup := range treeGroups {
		allTrees = append(allTrees, treeGroup...)
	}

	// sort the trees
	sort.Slice(allTrees, func(i, j int) bool { return compareTrees(allTrees[i], allTrees[j]) <= 0 })

	startDivider := newStartDivider()
	endDivider := newEndDivider()

	startIndex := 0
	endIndex := 0

	for i, v := range allTrees {
		if startIndex == 0 && compareTrees(startDivider, v) <= 0 {
			startIndex = i + 1
		}

		if startIndex != 0 && endIndex == 0 && compareTrees(endDivider, v) <= 0 {
			endIndex = i + 2 // account for startDivider position
			break
		}
	}

	return startIndex * endIndex
}

func parseToTrees(inputLineGroups [][]string) [][]tree {

	var answer [][]tree

	for _, v := range inputLineGroups {
		answer = append(answer, parseInputLinesToTrees(v))
	}

	return answer
}

func parseInputLinesToTrees(inputLines []string) []tree {
	if len(inputLines) != 2 {
		panic("invalid input")
	}

	var answer []tree
	answer = append(answer, parseInputLineToTree(inputLines[0]))
	answer = append(answer, parseInputLineToTree(inputLines[1]))
	return answer
}

func parseInputLineToTree(inputLine string) tree {
	root := tree{value: -1}
	currentTree := &root

	theNumber := ""
	for _, r := range inputLine {
		switch r {
		case '[':
			nt := tree{value: -1, parent: currentTree}
			currentTree.children = append(currentTree.children, &nt)
			currentTree = &nt
		case ']':
			if len(theNumber) > 0 {
				num, _ := strconv.Atoi(theNumber)
				currentTree.value = num
				theNumber = ""
			}
			currentTree = currentTree.parent
		case ',':
			if len(theNumber) > 0 {
				num, _ := strconv.Atoi(theNumber)
				currentTree.value = num
				theNumber = ""
			}
			currentTree = currentTree.parent
			nt := tree{value: -1, parent: currentTree}
			currentTree.children = append(currentTree.children, &nt)
			currentTree = &nt
		default:
			theNumber += string(r)
		}
	}
	return root
}

func newStartDivider() tree {
	startDivider := tree{value: -1}
	subTree := tree{value: -1, parent: &startDivider}
	subSubTree := tree{value: 2, parent: &subTree}
	subTree.children = append(subTree.children, &subSubTree)
	startDivider.children = append(startDivider.children, &subTree)

	return startDivider
}

func newEndDivider() tree {
	startDivider := tree{value: -1}
	subTree := tree{value: -1, parent: &startDivider}
	subSubTree := tree{value: 6, parent: &subTree}
	subTree.children = append(subTree.children, &subSubTree)
	startDivider.children = append(startDivider.children, &subTree)

	return startDivider
}
