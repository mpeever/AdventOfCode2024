package lib

import "fmt"

type Set[T any] struct {
	data map[string]T
}

func NewSet[T any](data []T) Set[T] {
	d := make(map[string]T, len(data))
	for _, v := range data {
		key := fmt.Sprintf("%v", v)
		d[key] = v
	}

	return Set[T]{
		data: d,
	}
}

func (s *Set[T]) Size() int {
	return len(s.data)
}

func (s *Set[T]) Add(value T) int {
	key := fmt.Sprintf("%v", value)
	// we might be able to skip this and just trust the map
	_, ok := s.data[key]
	if !ok {
		s.data[key] = value
	}
	return len(s.data)
}

func (s *Set[T]) Remove(value T) int {
	key := fmt.Sprintf("%v", value)
	_, ok := s.data[key]
	if !ok {
		return len(s.data)
	}

	delete(s.data, key)

	return len(s.data)
}

func (s *Set[T]) Contains(value T) bool {
	_, ok := s.data[fmt.Sprintf("%v", value)]
	return ok
}

func (s *Set[T]) Values() []T {
	values := make([]T, 0, len(s.data))
	for _, v := range s.data {
		values = append(values, v)
	}
	return values
}
