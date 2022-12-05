package datastructures

import "testing"

func Test_Stack(t *testing.T) {
	var myStack Stack[rune]
	var myValue rune
	var valid bool

	myStack.Push('A')
	myStack.Push('B')
	myStack.Push('C')
	myStack.Push('D')
	myStack.Push('E')

	myValue, valid = myStack.Pop()
	if valid && myValue != 'E' {
		t.Error("Expected E, got", myValue)
	}

	myValue, valid = myStack.Pop()
	if valid && myValue != 'D' {
		t.Error("Expected D, got", myValue)
	}

	myValue, valid = myStack.Pop()
	if valid && myValue != 'C' {
		t.Error("Expected C, got", myValue)
	}

	myValue, valid = myStack.Pop()
	if valid && myValue != 'B' {
		t.Error("Expected B, got", myValue)
	}

	myValue, valid = myStack.Peek()
	if valid && myValue != 'A' {
		t.Error("Expected A (peek), got", myValue)
	}

	myValue, valid = myStack.Pop()
	if valid && myValue != 'A' {
		t.Error("Expected A, got", myValue)
	}

	myValue, valid = myStack.Peek()
	if valid {
		t.Error("Expected invalid Peek")
	}
}
