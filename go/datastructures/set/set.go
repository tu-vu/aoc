package set

type Set[T comparable] struct {
	set map[T]struct{}
}

func New[T comparable]() *Set[T] {
	s := new(Set[T])
	s.set = make(map[T]struct{})
	return s
}

func (s *Set[T]) Has(v T) bool {
	_, ok := s.set[v]
	return ok
}

func (s *Set[T]) Add(v T) {
	s.set[v] = struct{}{}
}

func (s *Set[T]) Remove(v T) {
	delete(s.set, v)
}

func (s *Set[T]) Clear() {
	s.set = make(map[T]struct{})
}

func (s *Set[T]) Size() int {
	return len(s.set)
}
