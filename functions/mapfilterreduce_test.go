package functions

import "testing"

func Test_Map(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}

	result := Map(source, func(i int) int { return i + 5 })

	if result[0] != 6 {
		t.Error("Expected 6, got", result[0])
	}
}

func Test_Filter(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}

	result := Filter(source, func(i int) bool { return i%2 == 0 })

	if result[0] != 2 {
		t.Error("Expected 2, got", result[0])
	}

	if len(result) != 2 {
		t.Error("Expected length of 2, got", result[0])
	}
}

func Test_Reduce(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}

	result := Reduce(source, 0, func(i1, i2 int) int { return i1 + i2 })

	if result != 15 {
		t.Error("Expected 15, got", result)
	}
}
