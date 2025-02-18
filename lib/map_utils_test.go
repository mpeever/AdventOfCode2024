package lib

import "testing"

func testMapGrid() CharacterGrid {
	return CharacterGrid{
		Min: Point{X: 0, Y: 0},
		Max: Point{X: 4, Y: 4},
		Content: [][]string{
			{".", ".", "#", ".", "."},
			{"#", ".", ".", ".", "."},
			{".", ".", ".", ".", "."},
			{"#", ".", "^", ".", "#"}, // Guard starts at (2, 4)
			{"#", ".", ".", "#", "."},
		},
	}
}

func TestGuard_Move(t *testing.T) {
	grid := testMapGrid()
	p0 := Point{X: Width(2), Y: Height(4)}
	guard := Guard{
		Grid:        grid,
		Position:    p0,
		Orientation: NORTH,
		Path:        []Point{},
	}
	p, o, err := guard.Move()
	if err != nil {
		t.Errorf("Move failed: %v", guard)
	}

	pExpected := Point{X: Width(2), Y: Height(3)}
	if !p.Equals(&pExpected) {
		t.Errorf("Move did not move to correct point")
	}

	if o != NORTH {
		t.Errorf("Orientation should still be NORTH")
	}
}
