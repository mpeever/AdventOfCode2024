package main

import (
	. "AdventOfCode2024/lib"
	"testing"
)

func testGrid() CharacterGrid {
	return CharacterGrid{
		Min: Point{X: 0, Y: 0},
		Max: Point{X: 4, Y: 4},
		Content: [][]string{
			{"a", ".", "A", ".", "."},
			{".", "a", ".", "A", "."},
			{".", ".", ".", ".", "."},
			{".", "a", ".", ".", "A"},
			{".", ".", ".", ".", "a"},
		},
	}
}

func TestDemux(t *testing.T) {
	g := testGrid()
	grids := Demux(g)
	if len(grids) != 2 {
		t.Errorf("Demux returned wrong number of grids: %d", len(grids))
	}
}
