package Day08

import "strconv"

type treeGrid [][]int

func (t treeGrid) GetHeight(row, column int) int {
	return t[row][column]
}

func (t treeGrid) Rows() int {
	return len(t)
}

func (t treeGrid) Columns() int {
	return len(t[0])
}

func part1(inputLines []string) int {

	tg := parse(inputLines)

	total := 0
	for row := 0; row < tg.Rows(); row++ {
		for column := 0; column < tg.Columns(); column++ {
			myHeight := tg.GetHeight(row, column)
			if visible(tg, myHeight, row, column) {
				total++
				continue
			}
		}
	}

	return total
}

func parse(inputLines []string) treeGrid {

	var answer treeGrid

	for _, v := range inputLines {
		var row []int
		for j := 0; j < len(v); j++ {
			height, _ := strconv.Atoi(string(v[j]))
			row = append(row, height)
		}
		answer = append(answer, row)
	}

	return answer
}

func visible(tg treeGrid, myHeight, row, column int) bool {
	return visibleFromLeft(tg, myHeight, row, column) ||
		visibleFromRight(tg, myHeight, row, column) ||
		visibleFromAbove(tg, myHeight, row, column) ||
		visibleFromBelow(tg, myHeight, row, column)
}

func visibleFromLeft(tg treeGrid, myHeight, row, column int) bool {
	if column == 0 {
		return true
	}

	for i := column - 1; i >= 0; i-- {
		if tg.GetHeight(row, i) >= myHeight {
			return false
		}
	}

	return true
}

func visibleFromRight(tg treeGrid, myHeight, row, column int) bool {
	if column == tg.Columns()-1 {
		return true
	}

	for i := column + 1; i < tg.Columns(); i++ {
		if tg.GetHeight(row, i) >= myHeight {
			return false
		}
	}

	return true
}

func visibleFromAbove(tg treeGrid, myHeight, row, column int) bool {
	if row == 0 {
		return true
	}

	for i := row - 1; i >= 0; i-- {
		if tg.GetHeight(i, column) >= myHeight {
			return false
		}
	}

	return true
}

func visibleFromBelow(tg treeGrid, myHeight, row, column int) bool {
	if row == tg.Rows()-1 {
		return true
	}

	for i := row + 1; i < tg.Rows(); i++ {
		if tg.GetHeight(i, column) >= myHeight {
			return false
		}
	}

	return true
}
