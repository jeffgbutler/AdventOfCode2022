package datastructures

import "testing"

func Test_Queue(t *testing.T) {
	intQ := Queue[int]{}

	intQ.Enqueue(1)
	intQ.Enqueue(2)
	intQ.Enqueue(3)

	val, valid := intQ.Dequeue()
	if !valid || val != 1 {
		t.Error("Expected 1")
	}

	val, valid = intQ.Dequeue()
	if !valid || val != 2 {
		t.Error("Expected 2")
	}

	val, valid = intQ.Dequeue()
	if !valid || val != 3 {
		t.Error("Expected 3")
	}

	val, valid = intQ.Dequeue()
	if valid {
		t.Error("Expected Invalid")
	}
}
