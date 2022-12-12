package Day10

import (
	"github.com/jeffgbutler/AdventOfCode2022/functions"
	"strconv"
	"strings"
)

type snapshot struct {
	cycle, register int
}

func part1(inputLines []string) int {
	register := 1
	tickCount := 0
	var snapshots []snapshot

	for _, v := range inputLines {
		switch {
		case v == "noop":
			takeSnapshot(tickCount, register, &snapshots)
			tickCount++
		default:
			fields := strings.Fields(v)
			addX, _ := strconv.Atoi(fields[1])
			takeSnapshot(tickCount, register, &snapshots)
			tickCount++
			takeSnapshot(tickCount, register, &snapshots)
			tickCount++
			register += addX
		}
	}
	return functions.Reduce(snapshots, 0, func(i int, ss snapshot) int { return i + (ss.cycle * ss.register) })
}

func takeSnapshot(beginningTickCount, register int, snapshots *[]snapshot) {
	cycle := beginningTickCount + 1
	if cycle == 20 || (cycle-20)%40 == 0 {
		*snapshots = append(*snapshots, snapshot{cycle, register})
	}
}
