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

type bounds struct {
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
	lavaCubes, _ := parse(inputLines)

	exposedFaces := 0
	for cube := range lavaCubes {
		exposedFaces += calculateExposedFaces(cube, lavaCubes)
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
	lavaCubes, gridBorder := parse(inputLines)
	steamCubes := calculateSteamCubes(lavaCubes, gridBorder)

	exteriorFaces := 0

	for cube := range lavaCubes {
		exteriorFaces += calculateCoveredFaces(cube, steamCubes)
	}

	return exteriorFaces
}

func calculateSteamCubes(lavaCubes map[point]bool, gridBorder bounds) map[point]bool {
	// do a bfs walk over the grid to find everywhere steam could go
	expandedBorder := bounds{
		addVector(gridBorder.min, vector{-1, -1, -1}),
		addVector(gridBorder.max, vector{1, 1, 1}),
	}

	steamCubes := map[point]bool{}
	toDoList := datastructures.Queue[point]{}
	p := expandedBorder.min
	steamCubes[p] = true
	toDoList.Enqueue(p)

	for !toDoList.IsEmpty() {
		p, _ = toDoList.Dequeue()

		// find adjacent cubes not visited, and not lava
		for _, v := range adjacentVectors {
			adjacent := addVector(p, v)

			// is adjacent beyond border?
			if beyondBounds(adjacent, expandedBorder) {
				continue
			}

			// is adjacent cube lava?
			_, ok := lavaCubes[adjacent]
			if ok {
				continue
			}

			// has adjacent cube been visited?
			_, ok = steamCubes[adjacent]
			if ok {
				continue
			}

			steamCubes[adjacent] = true
			toDoList.Enqueue(adjacent)
		}
	}

	return steamCubes
}

func parse(inputLines []string) (map[point]bool, bounds) {
	answer := map[point]bool{}
	gridBorder := bounds{}

	p := parseInputLine(inputLines[0])
	answer[p] = true
	gridBorder.max = p
	gridBorder.min = p

	for i := 1; i < len(inputLines); i++ {
		p = parseInputLine(inputLines[i])
		gridBorder = updateGridBorder(p, gridBorder)
		answer[p] = true
	}

	return answer, gridBorder
}

func parseInputLine(inputLine string) point {
	f := strings.Split(inputLine, ",")

	x, _ := strconv.Atoi(f[0])
	y, _ := strconv.Atoi(f[1])
	z, _ := strconv.Atoi(f[2])

	return point{x, y, z}
}

func updateGridBorder(p point, gridBorder bounds) bounds {
	if p.x > gridBorder.max.x {
		gridBorder.max.x = p.x
	}

	if p.x < gridBorder.min.x {
		gridBorder.min.x = p.x
	}

	if p.y > gridBorder.max.y {
		gridBorder.max.y = p.y
	}

	if p.y < gridBorder.min.y {
		gridBorder.min.y = p.y
	}

	if p.z > gridBorder.max.z {
		gridBorder.max.z = p.z
	}

	if p.z < gridBorder.min.z {
		gridBorder.min.z = p.z
	}

	return gridBorder
}

func addVector(p point, v vector) point {
	return point{p.x + v.dx, p.y + v.dy, p.z + v.dz}
}

func beyondBounds(p point, b bounds) bool {
	if p.x > b.max.x ||
		p.x < b.min.x ||
		p.y > b.max.y ||
		p.y < b.min.y ||
		p.z > b.max.z ||
		p.z < b.min.z {
		return true
	}

	return false
}
