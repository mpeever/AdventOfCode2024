package lib

import "slices"

type Stack[T any] struct {
	elements []T
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{elements: make([]T, 0)}
}

func (stack *Stack[T]) Peek() (value T, ok bool) {
	if len(stack.elements) == 0 {
		ok = false
		return
	}
	value = stack.elements[len(stack.elements)-1]
	ok = true
	return
}

func (stack *Stack[T]) Pop() (value T, ok bool) {
	if len(stack.elements) == 0 {
		ok = false
		return
	}
	ok = true
	value = stack.elements[len(stack.elements)-1]
	stack.elements = stack.elements[:len(stack.elements)-1]
	return
}

func (stack *Stack[T]) Push(value T) {
	stack.elements = append(stack.elements, value)
}

func (stack *Stack[T]) Len() int {
	return len(stack.elements)
}

func (stack *Stack[T]) List() []T {
	buffer := make([]T, len(stack.elements))
	copy(buffer, stack.elements)
	slices.Reverse(buffer)
	return buffer
}
