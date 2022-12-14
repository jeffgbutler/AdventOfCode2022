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
	answer.rows = len(inputLines)
	answer.columns = len(inputLines[0])

	for y := 0; y < len(inputLines); y++ {
		inputLine := inputLines[y]
		for x := 0; x < len(inputLine); x++ {
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
	isVisited := map[point]bool{}
	parents := map[point]point{}

	queue.Enqueue(start)
	isVisited[start] = true

	reachedEnd := false

	for !queue.IsEmpty() {
		p, _ := queue.Dequeue()

		if p == end {
			reachedEnd = true
			break
		}

		for _, adjacentPoint := range adjacentPoints(p, grid) {
			if isVisited[adjacentPoint] {
				continue
			}

			parents[adjacentPoint] = p
			isVisited[adjacentPoint] = true
			queue.Enqueue(adjacentPoint)
		}
	}

	if reachedEnd {
		shortestPath := calculateShortestPath(end, parents)
		if shortestPath[0] == start {
			// answers do not include the start as a step
			return len(shortestPath) - 1
		} else {
			return -2 // no path from start to end
		}
	} else {
		return -1 // never found the end
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

func calculateShortestPath(end point, parents map[point]point) []point {
	var shortestPath []point
	currentPoint := end
	ok := true

	for ; ok == true; currentPoint, ok = parents[currentPoint] {
		shortestPath = append(shortestPath, currentPoint)
	}

	// reverse... this is from https://github.com/golang/go/wiki/SliceTricks
	for i := len(shortestPath)/2 - 1; i >= 0; i-- {
		opp := len(shortestPath) - 1 - i
		shortestPath[i], shortestPath[opp] = shortestPath[opp], shortestPath[i]
	}

	return shortestPath
}
