package Day09

import "testing"

func Test_Test1(t *testing.T) {
	head := position{2, 0}
	tail := position{0, 0}
	expected := position{1, 0}

	tail = moveTail(head, tail)

	if tail.x != expected.x || tail.y != expected.y {
		t.Error("expected", expected, "got", tail)
	}
}

func Test_Test2(t *testing.T) {
	head := position{2, -1}
	tail := position{0, 0}

	expected := position{1, -1}

	tail = moveTail(head, tail)

	if tail.x != expected.x || tail.y != expected.y {
		t.Error("expected", expected, "got", tail)
	}
}

func Test_Test3(t *testing.T) {
	head := position{1, -2}
	tail := position{0, 0}

	expected := position{1, -1}

	tail = moveTail(head, tail)

	if tail.x != expected.x || tail.y != expected.y {
		t.Error("expected", expected, "got", tail)
	}
}

func Test_Test4(t *testing.T) {
	head := position{0, -2}
	tail := position{0, 0}

	expected := position{0, -1}

	tail = moveTail(head, tail)

	if tail.x != expected.x || tail.y != expected.y {
		t.Error("expected", expected, "got", tail)
	}
}

func Test_Test5(t *testing.T) {
	head := position{-1, -2}
	tail := position{0, 0}

	expected := position{-1, -1}

	tail = moveTail(head, tail)

	if tail.x != expected.x || tail.y != expected.y {
		t.Error("expected", expected, "got", tail)
	}
}

func Test_Test6(t *testing.T) {
	head := position{-2, -1}
	tail := position{0, 0}

	expected := position{-1, -1}

	tail = moveTail(head, tail)

	if tail.x != expected.x || tail.y != expected.y {
		t.Error("expected", expected, "got", tail)
	}
}

func Test_Test7(t *testing.T) {
	head := position{-2, 0}
	tail := position{0, 0}

	expected := position{-1, 0}

	tail = moveTail(head, tail)

	if tail.x != expected.x || tail.y != expected.y {
		t.Error("expected", expected, "got", tail)
	}
}

func Test_Test8(t *testing.T) {
	head := position{-2, 1}
	tail := position{0, 0}

	expected := position{-1, 1}

	tail = moveTail(head, tail)

	if tail.x != expected.x || tail.y != expected.y {
		t.Error("expected", expected, "got", tail)
	}
}

func Test_Test9(t *testing.T) {
	head := position{-1, 2}
	tail := position{0, 0}

	expected := position{-1, 1}

	tail = moveTail(head, tail)

	if tail.x != expected.x || tail.y != expected.y {
		t.Error("expected", expected, "got", tail)
	}
}

func Test_Test10(t *testing.T) {
	head := position{0, 2}
	tail := position{0, 0}

	expected := position{0, 1}

	tail = moveTail(head, tail)

	if tail.x != expected.x || tail.y != expected.y {
		t.Error("expected", expected, "got", tail)
	}
}

func Test_Test11(t *testing.T) {
	head := position{1, 2}
	tail := position{0, 0}

	expected := position{1, 1}

	tail = moveTail(head, tail)

	if tail.x != expected.x || tail.y != expected.y {
		t.Error("expected", expected, "got", tail)
	}
}

func Test_Test12(t *testing.T) {
	head := position{2, 1}
	tail := position{0, 0}

	expected := position{1, 1}

	tail = moveTail(head, tail)

	if tail.x != expected.x || tail.y != expected.y {
		t.Error("expected", expected, "got", tail)
	}
}

func Test_Test13(t *testing.T) {
	head := position{4, -1}
	tail := position{2, 1}

	expected := position{3, 0}

	tail = moveTail(head, tail)

	if tail.x != expected.x || tail.y != expected.y {
		t.Error("expected", expected, "got", tail)
	}
}
