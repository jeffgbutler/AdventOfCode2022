package Day18

import (
	"strconv"
	"strings"
)

type point struct {
	x, y, z int
}

type face struct {
	corner1, corner2 point
}

func part1(inputLines []string) int {
	points := parse(inputLines)

	faceMap := map[face]int{}

	for _, p := range points {
		faces := pointToFaces(p)
		for _, f := range faces {
			i, _ := faceMap[f]
			faceMap[f] = i + 1
		}
	}

	answer := 0
	for _, i := range faceMap {
		if i == 1 {
			answer++
		}
	}

	return answer
}

func pointToFaces(p point) []face {
	faces := make([]face, 6)

	faces[0] = face{p, point{p.x + 1, p.y + 1, p.z}}
	faces[1] = face{p, point{p.x, p.y + 1, p.z + 1}}
	faces[2] = face{p, point{p.x + 1, p.y, p.z + 1}}

	faces[3] = face{point{p.x + 1, p.y, p.z}, point{p.x + 1, p.y + 1, p.z + 1}}
	faces[4] = face{point{p.x, p.y + 1, p.z}, point{p.x + 1, p.y + 1, p.z + 1}}
	faces[5] = face{point{p.x, p.y, p.z + 1}, point{p.x + 1, p.y + 1, p.z + 1}}

	return faces
}

func parse(inputLines []string) []point {
	var answer []point

	for _, inputLine := range inputLines {
		f := strings.Split(inputLine, ",")

		x, _ := strconv.Atoi(f[0])
		y, _ := strconv.Atoi(f[1])
		z, _ := strconv.Atoi(f[2])

		answer = append(answer, point{x, y, z})
	}

	return answer
}
