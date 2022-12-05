package functions

func Map[T1, T2 any](source []T1, mapper func(T1) T2) []T2 {
	answer := make([]T2, len(source))
	for i, v := range source {
		answer[i] = mapper(v)
	}
	return answer
}

func Filter[T any](source []T, filter func(T) bool) []T {
	var answer []T
	for _, v := range source {
		if filter(v) {
			answer = append(answer, v)
		}
	}
	return answer
}

func Reduce[T1, T2 any](source []T1, initialValue T2, reducer func(T2, T1) T2) T2 {
	r := initialValue
	for _, v := range source {
		r = reducer(r, v)
	}

	return r
}
