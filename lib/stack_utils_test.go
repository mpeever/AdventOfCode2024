package lib

import "testing"

func TestStack_Len(t *testing.T) {
	stack := NewStack[int]()
	if stack.Len() != 0 {
		t.Errorf("stack.Len() returned %d, expected %d", stack.Len(), 0)
	}

	stack.Push(42)
	if stack.Len() != 1 {
		t.Errorf("stack.Len() returned %d, expected %d", stack.Len(), 1)
	}

	stack.Push(3)
	stack.Push(4)
	if stack.Len() != 3 {
		t.Errorf("stack.Len() returned %d, expected %d", stack.Len(), 3)
	}

	value, ok := stack.Pop()
	if !ok {
		t.Errorf("stack.Pop() returned %v", ok)
	}
	if value != 4 {
		t.Errorf("stack.Pop() returned %d, expected %d", value, 4)
	}
}
