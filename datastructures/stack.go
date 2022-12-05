package datastructures

// inspired by Jon Bodner's "Learning Go" example

type Stack[T any] struct {
	values []T
}

func (ps *Stack[T]) Push(t T) {
	ps.values = append(ps.values, t)
}

func (ps *Stack[T]) Pop() (T, bool) {
	value, valid := ps.Peek()
	if valid {
		ps.values = ps.values[:len(ps.values)-1]
	}

	return value, valid
}

func (ps *Stack[T]) Peek() (T, bool) {
	length := len(ps.values)
	if length == 0 {
		var zero T
		return zero, false
	}

	return ps.values[length-1], true
}
