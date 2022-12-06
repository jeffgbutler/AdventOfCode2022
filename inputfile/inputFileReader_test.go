package inputfile

import "testing"

func Test_Read(t *testing.T) {
	lines, _ := Read("testdata/testData.txt")

	if len(lines) != 21 {
		t.Error("Expected 21, got", len(lines))
	}
}

func Test_ReadGroups(t *testing.T) {
	groups, _ := ReadGroups("testdata/testData.txt")

	if len(groups) != 6 {
		t.Error("Expected 6 groups, got", len(groups))
	}

	expectedLengths := []int{3, 1, 2, 3, 1, 1}
	for i, v := range groups {
		if len(v) != expectedLengths[i] {
			t.Error("Expected length", expectedLengths[i], "in group", i, "got", len(v))
		}
	}
}
