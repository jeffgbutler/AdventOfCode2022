package Day22

type point struct {
	x, y int
}

type positionAndHeading struct {
	position point
	heading  string // U, D, L, R
}

type vector struct {
	dx, dy int
}

type pathSegment struct {
	steps    int
	nextTurn string
}

var forwardVectors = map[string]vector{
	"U": {0, -1},
	"D": {0, 1},
	"L": {-1, 0},
	"R": {1, 0},
}

var backwardVectors = map[string]vector{
	"U": {0, 1},
	"D": {0, -1},
	"L": {1, 0},
	"R": {-1, 0},
}

func part1(inputGroups [][]string) int {
	mapLines := parseMapLines(inputGroups[0])
	path := parsePathLine(inputGroups[1][0])
	currentPosition := findInitialPosition(inputGroups[0][0])

	for _, ps := range path {
		currentPosition = processPathSegment(ps, currentPosition, mapLines)
	}

	answer := 1000 * (currentPosition.position.y + 1)
	answer += 4 * (currentPosition.position.x + 1)

	switch currentPosition.heading {
	case "U":
		answer += 3
	case "D":
		answer += 1
	case "L":
		answer += 2
	}

	return answer
}

func parseMapLines(mapLines []string) map[point]bool {
	answer := map[point]bool{}

	for i, line := range mapLines {
		parseMapLine(i, line, answer)
	}

	return answer
}

func parseMapLine(lineNumber int, line string, mapPoints map[point]bool) {
	for i := 0; i < len(line); i++ {
		switch line[i] {
		case '.':
			mapPoints[point{x: i, y: lineNumber}] = true
		case '#':
			mapPoints[point{x: i, y: lineNumber}] = false
		}
	}
}

func parsePathLine(pathLine string) []pathSegment {
	var answer []pathSegment

	num := 0

	for i := 0; i < len(pathLine); i++ {
		c := pathLine[i]
		switch c {
		case 'R':
			answer = append(answer, pathSegment{steps: num, nextTurn: "R"})
			num = 0
		case 'L':
			answer = append(answer, pathSegment{steps: num, nextTurn: "L"})
			num = 0
		default:
			num = num*10 + int(c-'0')
		}
	}

	answer = append(answer, pathSegment{steps: num, nextTurn: ""})

	return answer
}

func findInitialPosition(inputLine string) positionAndHeading {
	for i := 0; i < len(inputLine); i++ {
		if inputLine[i] == '.' {
			return positionAndHeading{position: point{x: i, y: 0}, heading: "R"}
		}
	}

	panic("no initial position found")
}

func calculateNewHeading(currentHeading string, leftOrRight string) string {
	switch currentHeading {
	case "U":
		if leftOrRight == "L" {
			return "L"
		} else {
			return "R"
		}
	case "D":
		if leftOrRight == "L" {
			return "R"
		} else {
			return "L"
		}
	case "L":
		if leftOrRight == "L" {
			return "D"
		} else {
			return "U"
		}
	case "R":
		if leftOrRight == "L" {
			return "U"
		} else {
			return "D"
		}
	}

	panic("unknown heading")
}

func processPathSegment(ps pathSegment, start positionAndHeading, mapPoints map[point]bool) positionAndHeading {

	// mapLines[point] == true means OK to move there
	// mapLines[point] == false means wall - stop
	// mapLines[point] ! ok means off the map

	currentPosition := start
	forwardVector := forwardVectors[currentPosition.heading]

	for i := 0; i < ps.steps; i++ {
		newPosition := addVector(currentPosition.position, forwardVector)

		openSpace, ok := mapPoints[newPosition]
		if !ok {
			// we're off the map, need to wrap
			// go backwards until we're off the map, then go forward one step
			backwardsVector := backwardVectors[currentPosition.heading]
			newPosition = addVector(currentPosition.position, backwardsVector)
			for {
				openSpace, ok = mapPoints[newPosition]
				if ok {
					newPosition = addVector(newPosition, backwardsVector)
				} else {
					// off the map
					newPosition = addVector(newPosition, forwardVector)
					openSpace, ok = mapPoints[newPosition]
					break
				}
			}
		}

		if openSpace {
			currentPosition.position = newPosition
			continue
		} else {
			// hit a wall
			break
		}
	}

	if ps.nextTurn != "" {
		currentPosition.heading = calculateNewHeading(currentPosition.heading, ps.nextTurn)
	}

	return currentPosition
}

func addVector(p point, v vector) point {
	return point{p.x + v.dx, p.y + v.dy}
}
