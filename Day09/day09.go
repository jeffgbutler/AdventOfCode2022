package Day09

import (
	"strconv"
	"strings"
)

type position struct {
	x, y int
}

func part1(inputLines []string) int {
	return moveRope(inputLines, 2)
}

func part2(inputLines []string) int {
	return moveRope(inputLines, 10)
}

func moveRope(inputLines []string, numberOfKnots int) int {

	positions := make([]position, numberOfKnots)

	tailPositionSet := map[position]bool{}
	tailPositionSet[positions[len(positions)-1]] = true

	for _, v := range inputLines {
		f := strings.Fields(v)
		steps, _ := strconv.Atoi(f[1])
		switch f[0] {
		case "U":
			processUp(steps, positions, tailPositionSet)
		case "D":
			processDown(steps, positions, tailPositionSet)
		case "L":
			processLeft(steps, positions, tailPositionSet)
		case "R":
			processRight(steps, positions, tailPositionSet)
		}
	}

	return len(tailPositionSet)
}

func processUp(steps int, positions []position, tailPositionSet map[position]bool) {
	for i := 0; i < steps; i++ {
		up(positions)
		tailPositionSet[positions[len(positions)-1]] = true
	}
}

func up(positions []position) {
	positions[0].y++
	for i := 0; i < len(positions)-1; i++ {
		positions[i+1] = moveTail(positions[i], positions[i+1])
	}
}

func processDown(steps int, positions []position, tailPositionSet map[position]bool) {
	for i := 0; i < steps; i++ {
		down(positions)
		tailPositionSet[positions[len(positions)-1]] = true
	}
}

func down(positions []position) {
	positions[0].y--
	for i := 0; i < len(positions)-1; i++ {
		positions[i+1] = moveTail(positions[i], positions[i+1])
	}
}

func processLeft(steps int, positions []position, tailPositionSet map[position]bool) {
	for i := 0; i < steps; i++ {
		left(positions)
		tailPositionSet[positions[len(positions)-1]] = true
	}
}

func left(positions []position) {
	positions[0].x--
	for i := 0; i < len(positions)-1; i++ {
		positions[i+1] = moveTail(positions[i], positions[i+1])
	}
}

func processRight(steps int, positions []position, tailPositionSet map[position]bool) {
	for i := 0; i < steps; i++ {
		right(positions)
		tailPositionSet[positions[len(positions)-1]] = true
	}
}

func right(positions []position) {
	positions[0].x++
	for i := 0; i < len(positions)-1; i++ {
		positions[i+1] = moveTail(positions[i], positions[i+1])
	}
}

func moveTail(headPosition, tailPosition position) position {
	dx := headPosition.x - tailPosition.x
	dy := headPosition.y - tailPosition.y

	newTail := tailPosition

	switch {
	case dx > 1:
		newTail.x++
		newTail.y += limit(dy)
	case dx < -1:
		newTail.x--
		newTail.y += limit(dy)
	case dy > 1:
		newTail.y++
		newTail.x += limit(dx)
	case dy < -1:
		newTail.y--
		newTail.x += limit(dx)
	}

	return newTail
}

func limit(i int) int {
	switch {
	case i > 0:
		return 1
	case i < 0:
		return -1
	default:
		return 0
	}
}
