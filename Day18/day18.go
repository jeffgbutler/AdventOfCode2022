package Day18

import (
	"github.com/jeffgbutler/AdventOfCode2022/datastructures"
	"strconv"
	"strings"
)

type point struct {
	x, y, z int
}

type vector struct {
	dx, dy, dz int
}

type extant struct {
	min, max point
}

var adjacentVectors = []vector{
	{1, 0, 0},
	{-1, 0, 0},
	{0, 1, 0},
	{0, -1, 0},
	{0, 0, 1},
	{0, 0, -1},
}

func part1(inputLines []string) int {
	cubes, _ := parse(inputLines)

	exposedFaces := 0
	for cube := range cubes {
		exposedFaces += calculateExposedFaces(cube, cubes)
	}

	return exposedFaces
}

func calculateExposedFaces(cube point, cubes map[point]bool) int {
	return 6 - calculateCoveredFaces(cube, cubes)
}

func calculateCoveredFaces(cube point, cubes map[point]bool) int {
	answer := 0

	for _, v := range adjacentVectors {
		_, ok := cubes[addVector(cube, v)]
		if ok {
			answer++
		}
	}

	return answer
}

func part2(inputLines []string) int {
	lavaCubes, gridExtant := parse(inputLines)
	steamCubes := calculateSteamCubes(lavaCubes, gridExtant)

	exteriorFaces := 0

	for cube := range lavaCubes {
		exteriorFaces += calculateCoveredFaces(cube, steamCubes)
	}

	return exteriorFaces
}

func calculateSteamCubes(lavaCubes map[point]bool, gridExtant extant) map[point]bool {
	// do a bfs walk over the extant to find everywhere steam could go
	expanded := extant{
		addVector(gridExtant.min, vector{-1, -1, -1}),
		addVector(gridExtant.max, vector{1, 1, 1}),
	}

	answer := map[point]bool{}
	toDoList := datastructures.Queue[point]{}
	p := expanded.min
	answer[p] = true
	toDoList.Enqueue(p)

	for !toDoList.IsEmpty() {
		p, _ = toDoList.Dequeue()

		// find adjacent cubes not visited, and not lava
		for _, v := range adjacentVectors {
			adjacent := addVector(p, v)

			// if adjacent beyond border...
			if adjacent.x > expanded.max.x || adjacent.y > expanded.max.y || adjacent.z > expanded.max.z {
				continue
			}

			if adjacent.x < expanded.min.x || adjacent.y < expanded.min.y || adjacent.z < expanded.min.z {
				continue
			}

			_, ok := lavaCubes[adjacent]
			if ok {
				// adjacent is lava
				continue
			}

			_, ok = answer[adjacent]
			if ok {
				// adjacent has been visited
				continue
			}

			answer[adjacent] = true
			toDoList.Enqueue(adjacent)
		}
	}

	return answer
}

func parse(inputLines []string) (map[point]bool, extant) {
	answer := map[point]bool{}
	gridExtant := extant{}

	p := parseInputLine(inputLines[0])
	answer[p] = true
	gridExtant.max.x = p.x
	gridExtant.min.x = p.x
	gridExtant.max.y = p.y
	gridExtant.min.y = p.y
	gridExtant.max.z = p.z
	gridExtant.min.z = p.z

	for i := 1; i < len(inputLines); i++ {
		p = parseInputLine(inputLines[i])
		gridExtant = updateGridExtant(p, gridExtant)
		answer[p] = true
	}

	return answer, gridExtant
}

func parseInputLine(inputLine string) point {
	f := strings.Split(inputLine, ",")

	x, _ := strconv.Atoi(f[0])
	y, _ := strconv.Atoi(f[1])
	z, _ := strconv.Atoi(f[2])

	return point{x, y, z}
}

func updateGridExtant(p point, gridExtant extant) extant {
	if p.x > gridExtant.max.x {
		gridExtant.max.x = p.x
	}

	if p.x < gridExtant.min.x {
		gridExtant.min.x = p.x
	}

	if p.y > gridExtant.max.y {
		gridExtant.max.y = p.y
	}

	if p.y < gridExtant.min.y {
		gridExtant.min.y = p.y
	}

	if p.z > gridExtant.max.z {
		gridExtant.max.z = p.z
	}

	if p.z < gridExtant.min.z {
		gridExtant.min.z = p.z
	}

	return gridExtant
}

func addVector(p point, v vector) point {
	return point{p.x + v.dx, p.y + v.dy, p.z + v.dz}
}
