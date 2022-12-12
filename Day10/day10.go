package Day10

import (
	"strconv"
	"strings"
)

func part1(inputLines []string) int {
	register := 1
	cycle := 1
	answer := 0

	for _, v := range inputLines {
		switch {
		case v == "noop":
			answer += calculateSignalStrength(cycle, register)
			cycle++
		default:
			answer += calculateSignalStrength(cycle, register)
			cycle++
			answer += calculateSignalStrength(cycle, register)
			cycle++

			fields := strings.Fields(v)
			addX, _ := strconv.Atoi(fields[1])
			register += addX
		}
	}
	return answer

}

func calculateSignalStrength(cycle, register int) int {
	if cycle == 20 || (cycle-20)%40 == 0 {
		return cycle * register
	} else {
		return 0
	}
}

func part2(inputLines []string) []string {
	register := 1
	cycle := 1
	var answer []string

	for _, v := range inputLines {
		switch {
		case v == "noop":
			answer = append(answer, calculatePixel(cycle, register))
			cycle++
		default:
			answer = append(answer, calculatePixel(cycle, register))
			cycle++

			answer = append(answer, calculatePixel(cycle, register))
			cycle++

			fields := strings.Fields(v)
			addX, _ := strconv.Atoi(fields[1])
			register += addX
		}
	}

	s1 := strings.Join(answer[0:40], "")
	s2 := strings.Join(answer[40:80], "")
	s3 := strings.Join(answer[80:120], "")
	s4 := strings.Join(answer[120:160], "")
	s5 := strings.Join(answer[160:200], "")
	s6 := strings.Join(answer[200:240], "")

	return []string{s1, s2, s3, s4, s5, s6}
}

func calculatePixel(cycle, register int) string {
	var answer string
	cycle = (cycle - 1) % 40
	if cycle >= register-1 && cycle <= register+1 {
		answer = "#"
	} else {
		answer = "."
	}

	return answer
}
