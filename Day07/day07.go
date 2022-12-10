package Day07

import (
	"github.com/jeffgbutler/AdventOfCode2022/functions"
	"sort"
	"strconv"
	"strings"
)

type file struct {
	name string
	size int
}

type directory struct {
	name           string
	parent         *directory
	files          []file
	subDirectories []directory
}

type directoryInfo struct {
	fullPath  string
	totalSize int
}

func (d *directory) AddFile(f file) {
	d.files = append(d.files, f)
}

func (d *directory) AddSubDirectory(name string) *directory {
	d.subDirectories = append(d.subDirectories, directory{name: name, parent: d})
	return &d.subDirectories[len(d.subDirectories)-1]
}

func (d *directory) Size() int {
	fileSize := functions.Reduce(d.files, 0, func(total int, f file) int { return total + f.size })
	directorySize := functions.Reduce(d.subDirectories, 0, func(total int, dir directory) int { return total + dir.Size() })
	return fileSize + directorySize
}

func (d *directory) FullPath() string {
	if d.parent == nil {
		return d.name
	} else {
		parentPath := d.parent.FullPath()
		if parentPath == "/" {
			return parentPath + d.name
		} else {
			return parentPath + "/" + d.name
		}
	}
}

func (d *directory) CalculateSummary() []directoryInfo {
	var answer []directoryInfo

	answer = append(answer, directoryInfo{d.FullPath(), d.Size()})

	for _, v := range d.subDirectories {
		answer = append(answer, v.CalculateSummary()...)
	}

	return answer
}

func part1(inputLines []string) int {
	root := parseInput(inputLines)

	summary := root.CalculateSummary()

	summary = functions.Filter(summary, func(d directoryInfo) bool { return d.totalSize <= 100000 })
	answer := functions.Reduce(summary, 0, func(i int, d directoryInfo) int { return i + d.totalSize })
	return answer
}

func part2(inputLines []string) int {
	root := parseInput(inputLines)

	freeSpace := 70000000 - root.Size()
	neededSpace := 30000000 - freeSpace

	summary := root.CalculateSummary()

	summary = functions.Filter(summary, func(d directoryInfo) bool { return d.totalSize >= neededSpace })
	sort.Slice(summary, func(s1, s2 int) bool { return summary[s1].totalSize < summary[s2].totalSize })
	answer := summary[0].totalSize
	return answer
}

func parseInput(inputLines []string) directory {
	// return the root directory
	var root directory
	var currentDirectory *directory

	for _, inputLine := range inputLines {
		switch {
		case inputLine == "$ cd /":
			root = directory{name: "/"}
			currentDirectory = &root
		case inputLine == "$ cd ..":
			currentDirectory = currentDirectory.parent
		case strings.HasPrefix(inputLine, "$ cd"):
			currentDirectory = currentDirectory.AddSubDirectory(parseCDLine(inputLine))
		case inputLine == "$ ls":
			continue
		case strings.HasPrefix(inputLine, "dir "):
			continue
		default:
			currentDirectory.AddFile(parseFileLine(inputLine))
		}
	}

	return root
}

func parseCDLine(inputLine string) string {
	tokens := strings.Split(inputLine, " ")
	if len(tokens) != 3 {
		panic("$ cd line has " + string(rune(len(tokens))) + " tokens")
	}

	dirName := tokens[2]
	return dirName
}

func parseFileLine(inputLine string) file {
	tokens := strings.Split(inputLine, " ")
	if len(tokens) != 2 {
		panic("default line has " + string(rune(len(tokens))) + " tokens")
	}
	size, _ := strconv.Atoi(tokens[0])
	return file{name: tokens[1], size: size}
}
