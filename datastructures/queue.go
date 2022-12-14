package datastructures

type Queue[T any] struct {
	values []T
}

func (q *Queue[T]) Enqueue(t T) {
	q.values = append(q.values, t)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.values) > 0 {
		answer := q.values[0]
		q.values = q.values[1:]
		return answer, true
	} else {
		var zero T
		return zero, false
	}
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.values) == 0
}
