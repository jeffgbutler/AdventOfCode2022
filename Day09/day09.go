package Day09

import (
	"strconv"
	"strings"
)

type position struct {
	x, y int
}

func part1(inputLines []string) int {

	headPosition := position{0, 0}
	tailPosition := position{0, 0}
	tailPositionSet := map[position]bool{}
	tailPositionSet[tailPosition] = true

	for _, v := range inputLines {
		f := strings.Fields(v)
		steps, _ := strconv.Atoi(f[1])
		switch f[0] {
		case "U":
			processUp(steps, &headPosition, &tailPosition, tailPositionSet)
		case "D":
			processDown(steps, &headPosition, &tailPosition, tailPositionSet)
		case "L":
			processLeft(steps, &headPosition, &tailPosition, tailPositionSet)
		case "R":
			processRight(steps, &headPosition, &tailPosition, tailPositionSet)
		}
	}

	return len(tailPositionSet)
}

func processUp(steps int, headPosition, tailPosition *position, tailPositionSet map[position]bool) {
	for i := 0; i < steps; i++ {
		headPosition.y++
		if headPosition.y-tailPosition.y > 1 {
			tailPosition.y++
			switch {
			case tailPosition.x > headPosition.x:
				tailPosition.x--
			case tailPosition.x < headPosition.x:
				tailPosition.x++
			}
			tailPositionSet[*tailPosition] = true
		}
	}
}

func processDown(steps int, headPosition, tailPosition *position, tailPositionSet map[position]bool) {
	for i := 0; i < steps; i++ {
		headPosition.y--
		if tailPosition.y-headPosition.y > 1 {
			tailPosition.y--
			switch {
			case tailPosition.x > headPosition.x:
				tailPosition.x--
			case tailPosition.x < headPosition.x:
				tailPosition.x++
			}
			tailPositionSet[*tailPosition] = true
		}
	}
}

func processLeft(steps int, headPosition, tailPosition *position, tailPositionSet map[position]bool) {
	for i := 0; i < steps; i++ {
		headPosition.x--
		if tailPosition.x-headPosition.x > 1 {
			tailPosition.x--
			switch {
			case tailPosition.y > headPosition.y:
				tailPosition.y--
			case tailPosition.y < headPosition.y:
				tailPosition.y++
			}
			tailPositionSet[*tailPosition] = true
		}
	}
}

func processRight(steps int, headPosition, tailPosition *position, tailPositionSet map[position]bool) {
	for i := 0; i < steps; i++ {
		headPosition.x++
		if headPosition.x-tailPosition.x > 1 {
			tailPosition.x++
			switch {
			case tailPosition.y > headPosition.y:
				tailPosition.y--
			case tailPosition.y < headPosition.y:
				tailPosition.y++
			}
			tailPositionSet[*tailPosition] = true
		}
	}
}
