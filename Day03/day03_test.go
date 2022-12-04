package Day03

import (
	"AdventOfCode2022/inputfile"
	"testing"
)

func Test_Part1ExampleData(t *testing.T) {
	inputLines, _ := inputfile.Read("testdata/example.txt")

	if len(inputLines) != 6 {
		t.Error("Expected 6 input lines, got", len(inputLines))
	}
}

func Test_WhatIsARune(t *testing.T) {
	s := "Fred"
	s1 := s[0]

	inta := int(s1)  // 97
	intz := int('z') // 122
	intA := int('A') // 65
	intZ := int('Z') // 90

	if inta != 22 {
		t.Error("inta is", inta)
	}

	if intz != 22 {
		t.Error("intz is", intz)
	}

	if intA != 22 {
		t.Error("intA is", intA)
	}

	if intZ != 22 {
		t.Error("intZ is", intZ)
	}
}
