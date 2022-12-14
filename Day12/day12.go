package Day12

import (
	"github.com/jeffgbutler/AdventOfCode2022/datastructures"
	"sort"
)

type point struct {
	x, y int
}

type heightMap struct {
	start, end    point
	grid          map[point]int
	rows, columns int
}

func newHeightMap() heightMap {
	return heightMap{grid: make(map[point]int)}
}

func part1(inputLines []string) int {
	theMap := parse(inputLines)
	return bfs(theMap.start, theMap.end, theMap.grid)
}

func part2(inputLines []string) int {
	theMap := parse(inputLines)

	var allSteps []int
	for x := 0; x < theMap.columns; x++ {
		for y := 0; y < theMap.rows; y++ {
			p := point{x, y}
			height, ok := theMap.grid[p]
			if ok && height == 1 {
				steps := bfs(p, theMap.end, theMap.grid)
				if steps != -1 {
					allSteps = append(allSteps, steps)
				}
			}
		}
	}

	sort.Slice(allSteps, func(i, j int) bool { return allSteps[i] < allSteps[j] })
	return allSteps[0]
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
				answer.grid[p] = 1
			case 'E':
				answer.end = p
				answer.grid[p] = 26
			default:
				val := int(char) - int('a') + 1
				answer.grid[p] = val
			}
		}
	}

	return answer
}

func bfs(start, end point, grid map[point]int) int {
	var queue datastructures.Queue[point]
	visitedMap := map[point]bool{}
	queue.Enqueue(start)
	visitedMap[start] = true

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
			if visitedMap[adjacentPoint] {
				continue
			}

			// set parent here!
			// adjacentPoint.parent = p
			// or
			// parentsMap[adjacentPoint] = p
			visitedMap[adjacentPoint] = true
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

func adjacentPoints(p point, theMap map[point]int) []point {
	var answer []point
	var myHeight = theMap[p]

	// brute force
	adjPoint := point{p.x + 1, p.y}
	if pointIsEligible(adjPoint, myHeight, theMap) {
		answer = append(answer, adjPoint)
	}

	adjPoint = point{p.x - 1, p.y}
	if pointIsEligible(adjPoint, myHeight, theMap) {
		answer = append(answer, adjPoint)
	}

	adjPoint = point{p.x, p.y + 1}
	if pointIsEligible(adjPoint, myHeight, theMap) {
		answer = append(answer, adjPoint)
	}

	adjPoint = point{p.x, p.y + -1}
	if pointIsEligible(adjPoint, myHeight, theMap) {
		answer = append(answer, adjPoint)
	}

	return answer
}

func pointIsEligible(adjPoint point, myHeight int, theMap map[point]int) bool {
	nodeHeight, ok := theMap[adjPoint]
	return ok && nodeHeight-1 <= myHeight
}
