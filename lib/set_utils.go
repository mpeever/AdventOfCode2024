package lib

type Set[T comparable] struct {
	Data []T
}

func NewSet[T comparable](data []T) Set[T] {
	data = Unique(data)
	return Set[T]{
		Data: data,
	}
}

func (s *Set[T]) Size() int {
	return len(s.Data)
}

func (s *Set[T]) Add(value T) int {
	s.Data = Unique(append(s.Data, value))
	return len(s.Data)
}

func (s *Set[T]) Remove(value T) int {
	if s.Contains(value) {
		dta := []T{}
		for _, d := range s.Data {
			if d != value {
				dta = append(dta, d)
			}
		}
		s.Data = Unique(dta)
	}
	return len(s.Data)
}

func (s *Set[T]) Contains(value T) bool {
	for _, v := range s.Data {
		if v == value {
			return true
		}
	}
	return false
}
