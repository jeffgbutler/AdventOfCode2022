package Day18

import (
	"strconv"
	"strings"
)

type point struct {
	x, y, z int
}

func part1(inputLines []string) int {
	cubes := parse(inputLines)

	exposedFaces := 0
	for cube := range cubes {
		exposedFaces += calculateExposedFaces(cube, cubes)
	}

	return exposedFaces
}

func calculateExposedFaces(cube point, cubes map[point]bool) int {
	answer := 6

	_, ok := cubes[point{cube.x + 1, cube.y, cube.z}]
	if ok {
		answer--
	}

	_, ok = cubes[point{cube.x - 1, cube.y, cube.z}]
	if ok {
		answer--
	}

	_, ok = cubes[point{cube.x, cube.y + 1, cube.z}]
	if ok {
		answer--
	}

	_, ok = cubes[point{cube.x, cube.y - 1, cube.z}]
	if ok {
		answer--
	}

	_, ok = cubes[point{cube.x, cube.y, cube.z + 1}]
	if ok {
		answer--
	}

	_, ok = cubes[point{cube.x, cube.y, cube.z - 1}]
	if ok {
		answer--
	}

	return answer
}

func parse(inputLines []string) map[point]bool {
	answer := map[point]bool{}

	for _, inputLine := range inputLines {
		f := strings.Split(inputLine, ",")

		x, _ := strconv.Atoi(f[0])
		y, _ := strconv.Atoi(f[1])
		z, _ := strconv.Atoi(f[2])

		answer[point{x, y, z}] = true
	}

	return answer
}
