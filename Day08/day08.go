package Day08

import "strconv"

type treeGrid [][]int

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
			myHeight := tg[row][column]
			if visible(tg, myHeight, row, column) {
				total++
				continue
			}
		}
	}

	return total
}

func part2(inputLines []string) int {

	tg := parse(inputLines)

	total := 0
	for row := 1; row < tg.Rows()-1; row++ {
		for column := 1; column < tg.Columns()-1; column++ {
			myHeight := tg[row][column]
			sc := scenicScore(tg, myHeight, row, column)
			if sc > total {
				total = sc
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
		if tg[row][i] >= myHeight {
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
		if tg[row][i] >= myHeight {
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
		if tg[i][column] >= myHeight {
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
		if tg[i][column] >= myHeight {
			return false
		}
	}

	return true
}

func scenicScore(tg treeGrid, myHeight, row, column int) int {
	return scenicScoreToLeft(tg, myHeight, row, column) *
		scenicScoreToRight(tg, myHeight, row, column) *
		scenicScoreAbove(tg, myHeight, row, column) *
		scenicScoreBelow(tg, myHeight, row, column)
}

func scenicScoreToLeft(tg treeGrid, myHeight, row, column int) int {
	if column == 0 {
		return 0
	}

	visibleTrees := 0
	for i := column - 1; i >= 0; i-- {
		visibleTrees++
		if tg[row][i] >= myHeight {
			break
		}
	}

	return visibleTrees
}

func scenicScoreToRight(tg treeGrid, myHeight, row, column int) int {
	if column == tg.Columns()-1 {
		return 0
	}

	visibleTrees := 0
	for i := column + 1; i < tg.Columns(); i++ {
		visibleTrees++
		if tg[row][i] >= myHeight {
			break
		}
	}

	return visibleTrees
}

func scenicScoreAbove(tg treeGrid, myHeight, row, column int) int {
	if row == 0 {
		return 0
	}

	visibleTrees := 0
	for i := row - 1; i >= 0; i-- {
		visibleTrees++
		if tg[i][column] >= myHeight {
			break
		}
	}

	return visibleTrees
}

func scenicScoreBelow(tg treeGrid, myHeight, row, column int) int {
	if row == tg.Rows()-1 {
		return 0
	}

	visibleTrees := 0
	for i := row + 1; i < tg.Rows(); i++ {
		visibleTrees++
		if tg[i][column] >= myHeight {
			break
		}
	}

	return visibleTrees
}
