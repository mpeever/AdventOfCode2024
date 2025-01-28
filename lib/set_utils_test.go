package lib

import "testing"

func TestSet_Size(t *testing.T) {
	input := []string{"a", "b", "c", "d", "e", "f", "g", "a", "b", "c", "d", "e", "f", "g"}
	instance := NewSet(input)
	if instance.Size() != 7 {
		t.Errorf("Set size should be 7")
		t.Fail()
	}
}

func TestSet_Add(t *testing.T) {
	input := []string{"a", "b", "c", "d", "e", "f", "g", "a", "b", "c", "d", "e", "f", "g"}
	instance := NewSet(input)
	instance.Add("a")
	if instance.Size() != 7 {
		t.Errorf("Set size should be 7")
		t.Fail()
	}
}
