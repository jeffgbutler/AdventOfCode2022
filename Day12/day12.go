package Day12

import (
	"github.com/jeffgbutler/AdventOfCode2022/datastructures"
)

type point struct {
	x, y int
}

type nodeData struct {
	height  int
	visited bool
}

type heightMap struct {
	start, end    point
	grid          map[point]nodeData
	rows, columns int
}

func newHeightMap() heightMap {
	return heightMap{grid: make(map[point]nodeData)}
}

func part1(inputLines []string) int {
	theMap := parse(inputLines)
	return bfs(theMap.start, theMap.end, theMap.grid)
}

func parse(inputLines []string) heightMap {
	answer := newHeightMap()

	for y := 0; y < len(inputLines); y++ {
		answer.rows++
		answer.columns = 0
		inputLine := inputLines[y]
		for x := 0; x < len(inputLine); x++ {
			answer.columns++
			p := point{x, y}
			char := inputLine[x]
			switch char {
			case 'S':
				answer.start = p
				answer.grid[p] = nodeData{height: 1}
			case 'E':
				answer.end = p
				answer.grid[p] = nodeData{height: 26}
			default:
				val := int(char) - int('a') + 1
				answer.grid[p] = nodeData{height: val}
			}
		}
	}

	return answer
}

func bfs(start, end point, grid map[point]nodeData) int {
	var queue datastructures.Queue[point]
	queue.Enqueue(start)
	nd1 := grid[start]
	nd1.visited = true
	grid[start] = nd1

	steps := 0
	nodesLeftInLayer := 1
	nodesInNextLayer := 0
	reachedEnd := false

	for {
		p, ok := queue.Dequeue()

		if !ok {
			break
		}

		if p == end {
			reachedEnd = true
			break
		}

		adjPoints := adjacentPoints(p, grid)
		for _, adjacentPoint := range adjPoints {
			nd := grid[adjacentPoint]
			if nd.visited {
				continue
			}

			nd.visited = true
			grid[adjacentPoint] = nd
			queue.Enqueue(adjacentPoint)
			nodesInNextLayer++
		}

		nodesLeftInLayer--

		if nodesLeftInLayer == 0 {
			steps++
			nodesLeftInLayer = nodesInNextLayer
			nodesInNextLayer = 0
		}
	}

	if reachedEnd {
		return steps
	} else {
		return -1
	}
}

func adjacentPoints(p point, theMap map[point]nodeData) []point {
	var answer []point
	var myHeight = theMap[p].height

	// brute force
	adjPoint := point{p.x + 1, p.y}
	nd, ok := theMap[adjPoint]
	if ok && nd.height-1 <= myHeight {
		answer = append(answer, adjPoint)
	}

	adjPoint = point{p.x - 1, p.y}
	nd, ok = theMap[adjPoint]
	if ok && nd.height-1 <= myHeight {
		answer = append(answer, adjPoint)
	}

	adjPoint = point{p.x, p.y + 1}
	nd, ok = theMap[adjPoint]
	if ok && nd.height-1 <= myHeight {
		answer = append(answer, adjPoint)
	}

	adjPoint = point{p.x, p.y + -1}
	nd, ok = theMap[adjPoint]
	if ok && nd.height-1 <= myHeight {
		answer = append(answer, adjPoint)
	}

	return answer
}
