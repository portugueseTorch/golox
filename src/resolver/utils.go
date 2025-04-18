package resolver

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		items: make([]T, 128),
	}
}

func (s *Stack[T]) push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) pop() {
	if len(s.items) == 0 {
		return
	}

	s.items = s.items[:len(s.items)-1]
}

func (s *Stack[T]) peek() (*T, bool) {
	if len(s.items) == 0 {
		return nil, false
	}

	return &s.items[len(s.items)-1], true
}
