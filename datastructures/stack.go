package datastructures

// inspired by Jon Bodner's "Learning Go" example

type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) Push(t T) {
	s.values = append(s.values, t)
}

func (s *Stack[T]) PushMultiple(t []T) {
	s.values = append(s.values, t...)
}

func (s *Stack[T]) Pop() (T, bool) {
	value, valid := s.Peek()
	if valid {
		s.values = s.values[:len(s.values)-1]
	}

	return value, valid
}

func (s *Stack[T]) PopMultiple(i int) ([]T, bool) {
	if i > len(s.values) {
		return nil, false
	}

	values := s.values[len(s.values)-i:]
	s.values = s.values[:len(s.values)-i]

	return values, true
}

func (s *Stack[T]) Peek() (T, bool) {
	length := len(s.values)
	if length == 0 {
		var zero T
		return zero, false
	}

	return s.values[length-1], true
}
