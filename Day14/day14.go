package Day14

import (
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type vector point

type gridRange struct {
	maxX, minX, maxY, minY int
}

type moveResult int

const (
	OK moveResult = iota
	Blocked
	OffGrid
)

func part1(inputLines []string) int {
	grid := parseInputLines(inputLines)
	gr := calculateGridRange(grid)

	ticks := 0
	for {
		if tick(grid, gr) {
			ticks++
		} else {
			break
		}
	}

	return ticks
}

func toString(grid map[point]rune) []string {
	gr := calculateGridRange(grid)
	var lines []string

	for y := gr.minY; y <= gr.maxY; y++ {
		line := ""
		for x := gr.minX; x <= gr.maxX; x++ {
			r, ok := grid[point{x, y}]
			if ok {
				line += string(r)
			} else {
				line += "."
			}
		}
		lines = append(lines, line)
	}

	return lines
}

func tick(grid map[point]rune, gr gridRange) bool {
	downVector := vector{0, 1}
	downLeftVector := vector{-1, 1}
	downRightVector := vector{1, 1}
	start := point{500, 0}

	current := start

	for {
		nextPosition := addVector(current, downVector)
		result := tryMove(nextPosition, grid, gr)
		if result == OK {
			current = nextPosition
			continue
		} else if result == OffGrid {
			return false
		}

		nextPosition = addVector(current, downLeftVector)
		result = tryMove(nextPosition, grid, gr)
		if result == OK {
			current = nextPosition
			continue
		} else if result == OffGrid {
			return false
		}

		nextPosition = addVector(current, downRightVector)
		result = tryMove(nextPosition, grid, gr)
		if result == OK {
			current = nextPosition
			continue
		} else if result == OffGrid {
			return false
		}

		// none of the next positions work, stop where we are
		grid[current] = 'o'
		return true
	}
}

func tryMove(nextPosition point, grid map[point]rune, gr gridRange) moveResult {
	if isPointOffGrid(nextPosition, gr) {
		return OffGrid
	}

	_, ok := grid[nextPosition]
	if ok {
		return Blocked
	} else {
		return OK
	}
}

func isPointOffGrid(p point, gr gridRange) bool {
	return p.x > gr.maxX || p.x < gr.minX || p.y > gr.maxY
}

func calculateGridRange(grid map[point]rune) gridRange {
	answer := gridRange{}

	initialized := false
	for p := range grid {
		if !initialized {
			// initialize
			answer.maxX = p.x
			answer.minX = p.x
			answer.maxY = p.y
			answer.minY = p.y
			initialized = true
			continue
		}

		if p.x > answer.maxX {
			answer.maxX = p.x
		}

		if p.x < answer.minX {
			answer.minX = p.x
		}

		if p.y > answer.maxY {
			answer.maxY = p.y
		}

		if p.y < answer.minY {
			answer.minY = p.y
		}
	}

	return answer
}

func parseInputLines(inputLines []string) map[point]rune {
	answer := map[point]rune{}

	for _, v := range inputLines {
		parseInputLine(v, answer)
	}

	return answer
}

func parseInputLine(inputLine string, answer map[point]rune) {
	points := strings.Split(inputLine, " -> ")

	for i := 0; i < len(points)-1; i++ {
		start := parsePair(points[i])
		end := parsePair(points[i+1])
		populateMap(start, end, answer)
	}
}

func parsePair(s string) point {
	nums := strings.Split(s, ",")

	answer := point{}
	answer.x, _ = strconv.Atoi(nums[0])
	answer.y, _ = strconv.Atoi(nums[1])
	return answer
}

func populateMap(start, end point, grid map[point]rune) {
	grid[start] = '#'
	grid[end] = '#'

	incrementVector := calculateIncrementVector(start, end)
	currentPoint := start

	for currentPoint != end {
		currentPoint = addVector(currentPoint, incrementVector)
		grid[currentPoint] = '#'
	}
}

func calculateIncrementVector(start, end point) vector {
	switch {
	case start.x > end.x:
		return vector{-1, 0}
	case start.x < end.x:
		return vector{1, 0}
	case start.y > end.y:
		return vector{0, -1}
	default:
		return vector{0, 1}
	}
}

func addVector(start point, increment vector) point {
	return point{start.x + increment.x, start.y + increment.y}
}
