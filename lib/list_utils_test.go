package lib

import (
	"strings"
	"testing"
)

func TestAll_WithInt64(t *testing.T) {
	input := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	output := All(input, func(x int64) bool { return x > 0 })
	if !output {
		t.Error("Expected true, got false")
	}
}

func TestAll_WithString(t *testing.T) {
	input := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	output := All(input, func(x string) bool { return len(x) == 1 })
	if !output {
		t.Error("Expected true, got false")
	}
}

func TestAny_WithInt64(t *testing.T) {
	input := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	output := Any(input, func(x int64) bool { return x > 10 })
	if !output {
		t.Error("Expected true, got false")
	}
}

func TestAny_WithString(t *testing.T) {
	input := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	output := Any(input, func(x string) bool { return strings.EqualFold(x, "h") })
	if !output {
		t.Error("Expected true, got false")
	}
}

func TestMap_WithInt64(t *testing.T) {
	input := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	expected := []int64{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200}

	actual := Map(input, func(x int64) int64 { return x * 10 })
	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
		t.Fail()
	}

	for i, a := range actual {
		if a != expected[i] {
			t.Fail()
		}
	}
}

func TestRemoveIfNot(T *testing.T) {
	input := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	expected := []int64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	actual := RemoveIfNot(input, func(x int64) bool { return x%2 == 0 })

	if len(actual) != len(expected) {
		T.Fail()
	}
	for i, a := range actual {
		if a != expected[i] {
			T.Fail()
		}
	}
}

func TestUnique_WithInt64(t *testing.T) {
	input := []int64{1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 4, 4, 4}
	expected := []int64{1, 2, 3, 4}
	actual := Unique(input)
	if len(actual) != len(expected) {
		t.Fail()
	}
}

func TestUnique_WithString(t *testing.T) {
	input := []string{"the", "the", "quick", "brown", "fox", "fox", "jumps", "the"}
	expected := []string{"the", "quick", "brown", "fox", "jumps"}
	actual := Unique(input)
	if len(actual) != len(expected) {
		t.Error("lengths do not match")
		t.Fail()
	}
}

func TestPairs(t *testing.T) {
	input := []string{"a", "b", "c", "d", "e"}
	expected := [][]string{
		{"a", "b"}, {"a", "c"}, {"a", "d"}, {"a", "e"},
		{"b", "a"}, {"b", "c"}, {"b", "d"}, {"b", "e"},
		{"c", "a"}, {"c", "b"}, {"c", "d"}, {"c", "e"},
		{"d", "a"}, {"d", "b"}, {"d", "c"}, {"d", "e"},
		{"e", "a"}, {"e", "b"}, {"e", "c"}, {"e", "d"},
	}
	actual := Pairs(input)
	if len(actual) != len(expected) {
		t.Error("lengths do not match")
		t.Fail()
	}
}
