package Day25

import "strings"

// 2 + 2 (4) -> 1-
// 2 + 1 (3) -> 1=
// 2 + 0 (2) -> 2
// 2 + - (1) -> 1
// 2 + = (0) -> 0
//
// 1 + 1 (2) -> 2
// 1 + 0 (1) -> 1
// 1 + - (0) -> 0
// 1 + = (-1) -> -
//
// 0 + 0 (0) -> 0
// 0 + - (-1) -> -
// 0 + = (-2) -> =
//
// - + - (-2) -> =
// - + = (-3) -> -2
//
// = + = (-4) -> -1

type addends struct {
	a1, a2 string
}

type addResult struct {
	result, carry string
}

var addResults = map[addends]addResult{
	addends{"2", "2"}: {"-", "1"},
	addends{"2", "1"}: {"=", "1"},
	addends{"2", "0"}: {"2", "0"},
	addends{"2", "-"}: {"1", "0"},
	addends{"2", "="}: {"0", "0"},

	addends{"1", "1"}: {"2", "0"},
	addends{"1", "0"}: {"1", "0"},
	addends{"1", "-"}: {"0", "0"},
	addends{"1", "="}: {"-", "0"},

	addends{"0", "0"}: {"0", "0"},
	addends{"0", "-"}: {"-", "0"},
	addends{"0", "="}: {"=", "0"},

	addends{"-", "-"}: {"=", "0"},
	addends{"-", "="}: {"2", "-"},

	addends{"=", "="}: {"1", "-"},
}

func snafuAddTwoDigits(a1, a2 string) addResult {
	r, ok := addResults[addends{a1, a2}]
	if ok {
		return r
	} else {
		return addResults[addends{a2, a1}]
	}
}

func snafuAddThreeDigits(a1, a2, a3 string) addResult {
	r1 := snafuAddTwoDigits(a1, a2)
	r2 := snafuAddTwoDigits(r1.result, a3)
	c1 := snafuAddTwoDigits(r1.carry, r2.carry)
	return addResult{result: r2.result, carry: c1.result}
}

func snafuAdd(s1, s2 string) string {
	var result []string

	// make strings the same length by prepending zeros to the shorter string
	switch {
	case len(s1) < len(s2):
		s1 = makeLength(s1, len(s2))
	case len(s2) < len(s1):
		s2 = makeLength(s2, len(s1))
	}

	carry := "0"
	for i := len(s1) - 1; i >= 0; i-- {
		a1 := s1[i]
		a2 := s2[i]
		r := snafuAddThreeDigits(string(a1), string(a2), carry)
		result = append(result, r.result)
		carry = r.carry
	}

	if carry != "0" {
		result = append(result, carry)
	}

	// reverse the slice...
	for i := len(result)/2 - 1; i >= 0; i-- {
		opp := len(result) - 1 - i
		result[i], result[opp] = result[opp], result[i]
	}

	return strings.Join(result, "")
}

func makeLength(s string, length int) string {
	for len(s) < length {
		s = "0" + s
	}

	return s
}

func part1(inputLines []string) string {
	answer := "0"
	for _, s := range inputLines {
		answer = snafuAdd(answer, s)
	}

	return answer
}
