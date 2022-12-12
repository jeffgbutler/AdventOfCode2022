package Day10

import (
	"github.com/jeffgbutler/AdventOfCode2022/inputfile"
	"testing"
)

func Test_Part1WithExampleData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/example.txt")

	answer := part1(inputLines)

	if answer != 13140 {
		t.Error("Expected 13140, got", answer)
	}
}

func Test_Part1WithActualData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/actual.txt")

	answer := part1(inputLines)

	if answer != 14160 {
		t.Error("Expected 14160, got", answer)
	}
}

func Test_Part2WithExampleData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/example.txt")

	expected := []string{
		"##..##..##..##..##..##..##..##..##..##..",
		"###...###...###...###...###...###...###.",
		"####....####....####....####....####....",
		"#####.....#####.....#####.....#####.....",
		"######......######......######......####",
		"#######.......#######.......#######....."}

	answer := part2(inputLines)

	if len(answer) != len(expected) {
		t.Error("Expected equal lengths")
	}

	for i, v := range expected {
		if answer[i] != v {
			t.Error("String", i, "doesn't match")
		}
	}
}

func Test_Part2WithActualData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/actual.txt")

	expected := []string{
		"###....##.####.###..###..####.####..##..",
		"#..#....#.#....#..#.#..#.#....#....#..#.",
		"#..#....#.###..#..#.#..#.###..###..#....",
		"###.....#.#....###..###..#....#....#....",
		"#.#..#..#.#....#.#..#....#....#....#..#.",
		"#..#..##..####.#..#.#....####.#.....##.."}

	answer := part2(inputLines)

	if len(answer) != len(expected) {
		t.Error("Expected equal lengths")
	}

	for i, v := range expected {
		if answer[i] != v {
			t.Error("String", i, "doesn't match")
		}
	}
}
