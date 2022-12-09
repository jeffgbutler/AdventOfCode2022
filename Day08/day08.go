package Day08

import "strconv"

type tree struct {
	height  int
	visible bool
}

type treeGrid [][]tree

func (t *treeGrid) SetVisible(row, column int) {
	(*t)[row][column].visible = true
}

func (t *treeGrid) IsVisible(row, column int) bool {
	return (*t)[row][column].visible
}

func (t *treeGrid) GetHeight(row, column int) int {
	return (*t)[row][column].height
}

func (t *treeGrid) Rows() int {
	return len(*t)
}

func (t *treeGrid) Columns() int {
	return len((*t)[0])
}

func part1(inputLines []string) int {

	tg := parse(inputLines)

	calc1(&tg)
	calc2(&tg)
	calc3(&tg)
	calc4(&tg)

	total := 0
	for row := 0; row < tg.Rows(); row++ {
		for column := 0; column < tg.Columns(); column++ {
			if tg.IsVisible(row, column) {
				total++
			}
		}
	}

	return total
}

func parse(inputLines []string) treeGrid {

	var answer treeGrid

	for _, v := range inputLines {
		var row []tree
		for j := 0; j < len(v); j++ {
			height, _ := strconv.Atoi(string(v[j]))
			row = append(row, tree{height: height})
		}
		answer = append(answer, row)
	}

	return answer
}

func calc1(tg *treeGrid) {
	for row := 0; row < tg.Rows(); row++ {
		height := tg.GetHeight(row, 0)
		tg.SetVisible(row, 0)
		for column := 1; column < tg.Columns()-1; column++ {
			if tg.GetHeight(row, column) > height {
				height = tg.GetHeight(row, column)
				tg.SetVisible(row, column)
			}
		}
	}
}

func calc2(tg *treeGrid) {
	for row := 0; row < tg.Rows(); row++ {
		height := tg.GetHeight(row, tg.Columns()-1)
		tg.SetVisible(row, tg.Columns()-1)
		for column := tg.Columns() - 2; column >= 1; column-- {
			if tg.GetHeight(row, column) > height {
				height = tg.GetHeight(row, column)
				tg.SetVisible(row, column)
			}
		}
	}
}

func calc3(tg *treeGrid) {
	for column := 0; column < tg.Columns(); column++ {
		height := tg.GetHeight(0, column)
		tg.SetVisible(0, column)
		for row := 1; row < tg.Rows()-1; row++ {
			if tg.GetHeight(row, column) > height {
				height = tg.GetHeight(row, column)
				tg.SetVisible(row, column)
			}
		}
	}
}

func calc4(tg *treeGrid) {
	for column := 0; column < tg.Columns(); column++ {
		height := tg.GetHeight(tg.Rows()-1, column)
		tg.SetVisible(tg.Rows()-1, column)
		for row := tg.Rows() - 2; row >= 1; row-- {
			if tg.GetHeight(row, column) > height {
				height = tg.GetHeight(row, column)
				tg.SetVisible(row, column)
			}
		}
	}
}
