package lib

import (
	"testing"
)

func TestNewBlockMap_WithSmallInput(t *testing.T) {
	input := "12345"
	expected := "0..111....22222"
	actual, _ := NewBlockMap(input)
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestNewBlockMap_WithLargerInput(t *testing.T) {
	input := "2333133121414131402"
	expected := "00...111...2...333.44.5555.6666.777.888899"
	actual, _ := NewBlockMap(input)
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
