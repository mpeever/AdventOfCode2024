package lib

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func testGrid() CharacterGrid {
	return CharacterGrid{
		Min: Point{X: 0, Y: 0},
		Max: Point{X: 4, Y: 4},
		Content: [][]string{
			{"a", "b", "c", "d", "e"},
			{"f", "g", "h", "i", "j"},
			{"k", "l", "m", "n", "o"},
			// Skip p, you know why
			{"q", "r", "s", "t", "u"},
			{"v", "w", "x", "y", "z"},
		},
	}
}

func TestCharacterGrid_Clone(t *testing.T) {
	instance := testGrid()
	clone := instance.Clone()

	for y, row := range clone.Content {
		for x, col := range row {
			p := Point{X: Width(x), Y: Height(y)}

			// check the unaltered values match
			v0 := instance.Char(p)
			v1 := clone.Char(p)
			if v0 != v1 {
				t.Errorf("expected values to match: %v, %v", v0, v1)
			}

			// update the clone and confirm the original is unchanged
			clone.Update(p, fmt.Sprintf("%v_%v", col, col))
			v2 := clone.Char(p)
			if v0 == v2 {
				t.Errorf("expected values NOT to match: %v, %v", v0, v2)
			}
		}
	}
}

func TestCharacterGrid_Boundaries(T *testing.T) {
	instance := testGrid()

	// Check grid boundaries
	if !instance.Includes(Point{X: 2, Y: 2}) {
		T.Fail()
	}
	if instance.Includes(Point{X: 0, Y: 5}) {
		T.Fail()
	}
	if instance.Includes(Point{X: 5, Y: 5}) {
		T.Fail()
	}
}

func TestCharacterGrid_Values(T *testing.T) {
	instance := testGrid()

	// Check we get the right characters for a given Point
	if instance.Char(Point{X: 0, Y: 0}) != "a" {
		T.Fail()
	}
	if instance.Char(Point{X: 0, Y: 4}) != "v" {
		T.Fail()
	}

	if instance.Char(Point{X: 2, Y: 2}) != "m" {
		T.Fail()
	}
}

func TestCharacterGrid_Update(t *testing.T) {
	instance := testGrid()

	val := "Updated"
	point := Point{X: 2, Y: 2}

	instance.Update(point, val)
	if instance.Char(point) != val {
		t.Fail()
	}
}

func TestCharacterGrid_Neighbors_LEFT(T *testing.T) {
	instance := testGrid()
	p0 := Point{X: 2, Y: 2}
	d := LEFT
	enp := Point{X: 1, Y: 2}
	np, _ := instance.NextPoint(p0, d)
	if !np.Equals(&enp) {
		T.Fail()
	}
}

func TestCharacterGrid_Neighbors_RIGHT(T *testing.T) {
	instance := testGrid()
	p0 := Point{X: 2, Y: 2}
	d := RIGHT
	enp := Point{X: 3, Y: 2}
	np, _ := instance.NextPoint(p0, d)
	if !np.Equals(&enp) {
		T.Fail()
	}
}

func TestCharacterGrid_Neighbors_DOWN(T *testing.T) {
	instance := testGrid()
	p0 := Point{X: 2, Y: 2}
	d := DOWN
	enp := Point{X: 2, Y: 3}
	np, _ := instance.NextPoint(p0, d)
	if !np.Equals(&enp) {
		T.Fail()
	}
}

func TestCharacterGrid_Neighbors_DOWNRIGHT(T *testing.T) {
	instance := testGrid()
	p0 := Point{X: 2, Y: 2}
	d := DOWNRIGHT
	enp := Point{X: 3, Y: 3}
	np, _ := instance.NextPoint(p0, d)
	if !np.Equals(&enp) {
		T.Fail()
	}
}

func TestCharacterGrid_PointsAround(T *testing.T) {
	instance := testGrid()
	p0 := Point{X: 2, Y: 2}

	expected := []Point{
		{X: 1, Y: 2}, // LEFT
		{X: 1, Y: 1}, // UPLEFT
		{X: 2, Y: 1}, // UP
		{X: 3, Y: 1}, // UPRIGHT
		{X: 3, Y: 2}, // RIGHT
		{X: 3, Y: 3}, // DOWNRIGHT
		{X: 2, Y: 3}, // DOWN
		{X: 1, Y: 3}, // DOWNLEFT
	}

	np := instance.PointsAround(p0)
	if len(np) != len(expected) {
		T.Fail()
	}
	for _, exp := range expected {
		if !Any(np, func(p Point) bool { return p.Equals(&exp) }) {
			T.Fail()
		}
	}
}

func TestCharacterGrid_PointsAround_OnEdge(T *testing.T) {
	instance := testGrid()
	p0 := Point{X: 4, Y: 2}

	expected := []Point{
		{X: 3, Y: 2}, // LEFT
		{X: 3, Y: 1}, // UPLEFT
		{X: 4, Y: 1}, // UP
		{X: 4, Y: 3}, // DOWN
		{X: 3, Y: 3}, // DOWNLEFT
	}

	np := instance.PointsAround(p0)
	if len(np) != len(expected) {
		T.Fail()
	}
	for _, exp := range expected {
		if !Any(np, func(p Point) bool { return p.Equals(&exp) }) {
			T.Fail()
		}
	}
}

func TestCharacterGrid_Corners(T *testing.T) {
	instance := testGrid()
	p0 := Point{X: 2, Y: 2}

	expected := []Point{
		{X: 1, Y: 1}, // UPLEFT
		{X: 3, Y: 1}, // UPRIGHT
		{X: 3, Y: 3}, // DOWNRIGHT
		{X: 1, Y: 3}, // DOWNLEFT
	}

	np := instance.Corners(p0)
	if len(np) != len(expected) {
		T.Errorf("Expected %d points, got %d", len(expected), len(np))
		T.Fail()
	}
	for _, exp := range expected {
		if !Any(np, func(p Point) bool { return p.Equals(&exp) }) {
			T.Fail()
		}
	}
}

func TestVector_String(t *testing.T) {
	instance := testGrid()
	v := Vector{Points: []Point{{X: 1, Y: 1}, {X: 2, Y: 2}, {X: 3, Y: 3}},
		Grid: &instance,
	}
	expected := "gmt"
	actual := v.String()

	if expected != actual {
		t.Fail()

	}
}

func TestVector_Through(T *testing.T) {
	instance := testGrid()
	p0 := Point{X: 1, Y: 1}
	p1 := Point{X: 2, Y: 1}
	s := Size(3)

	expected := Vector{
		Points: []Point{{X: 1, Y: 1}, {X: 2, Y: 1}, {X: 3, Y: 1}},
	}

	actual, _ := instance.VectorThrough(p0, p1, s)
	if !expected.Equals(&actual) {
		T.Fail()
	}
}

func TestVector_Through_WithSize4(T *testing.T) {
	instance := testGrid()
	p0 := Point{X: 1, Y: 1}
	p1 := Point{X: 2, Y: 1}
	s := Size(4)

	expected := Vector{
		Points: []Point{{X: 1, Y: 1}, {X: 2, Y: 1}, {X: 3, Y: 1}, {X: 4, Y: 1}},
	}

	actual, _ := instance.VectorThrough(p0, p1, s)
	if !expected.Equals(&actual) {
		T.Fail()
	}
}

func TestValuesAround(T *testing.T) {
	instance := testGrid()
	p0 := Point{X: 1, Y: 1}
	expected := []string{"a", "b", "c", "f", "h", "k", "l", "m"}
	actual := instance.ValuesAround(p0)
	if len(actual) != len(expected) {
		T.Errorf("Expected %d values, got %d", len(expected), len(actual))
		T.Fail()
	}
	for _, exp := range expected {
		if !Any(actual, func(s string) bool { return s == exp }) {
			T.Errorf("expected %s to be in %v", exp, actual)
			T.Fail()
		}
	}
}

func TestValuesAround_WithBorder(T *testing.T) {
	instance := testGrid()
	p0 := Point{X: 4, Y: 1}
	expected := []string{
		"d", "e",
		"i",
		"n", "o",
	}
	actual := instance.ValuesAround(p0)
	if len(actual) != len(expected) {
		T.Errorf("Expected %d values, got %d", len(expected), len(actual))
		T.Fail()
	}
	for _, exp := range expected {
		if !Any(actual, func(s string) bool { return s == exp }) {
			T.Errorf("expected %s to be in %v", exp, actual)
			T.Fail()
		}
	}
}

func TestIntersections(t *testing.T) {
	grid := testGrid()
	v0 := Vector{
		Grid:   &grid,
		Points: []Point{{X: 1, Y: 1}, {X: 2, Y: 1}, {X: 3, Y: 1}},
	}
	v1 := Vector{
		Grid:   &grid,
		Points: []Point{{X: 2, Y: 0}, {X: 2, Y: 1}, {X: 2, Y: 3}},
	}
	expected := []Intersection{{V0: v0, V: v1}}
	actual := Intersections([]Vector{v0, v1})
	if len(actual) != len(expected) {
		t.Fail()
	}
	for _, exp := range expected {
		if !Any(actual, func(intersection Intersection) bool { return intersection.Equals(&exp) }) {
			t.Fail()
		}
	}
}

func TestCharacterGrid_Diagonals(t *testing.T) {
	grid := testGrid()
	p0 := Point{X: 2, Y: 2}
	expected := []Vector{
		{
			Grid:   &grid,
			Points: []Point{{X: 1, Y: 1}, {X: 2, Y: 2}, {X: 3, Y: 3}},
		},
		{
			Grid:   &grid,
			Points: []Point{{X: 3, Y: 1}, {X: 2, Y: 2}, {X: 1, Y: 3}},
		},
	}

	actual, err := grid.Diagonals(p0)
	if err != nil {
		t.Fail()
	}
	if len(actual) != len(expected) {
		t.Fail()
	}
	for _, exp := range expected {
		if !Any(actual, func(v Vector) bool {
			return v.Equals(&exp)
		}) {
			t.Fail()
		}
	}
}

func TestCharacterGrid_VectorToEdge(t *testing.T) {
	grid := testGrid()
	p0 := Point{X: 1, Y: 1} // "g"
	v := grid.VectorToEdge(p0, DOWNRIGHT)
	expected := Vector{
		Grid:   &grid,
		Points: []Point{{X: 1, Y: 1}, {X: 2, Y: 2}, {X: 3, Y: 3}, {X: 4, Y: 4}},
	}
	if !reflect.DeepEqual(expected, v) {
		t.Errorf("expected: %v, got: %v", expected, v)
	}
}

func TestCharacterGrid_FloatDistance(t *testing.T) {
	grid := testGrid()
	p0 := Point{X: 1, Y: 1}
	p1 := Point{X: 2, Y: 2}
	d, err := grid.FloatDistance(p0, p1)
	if err != nil || d != Distance(math.Sqrt(2.0)) {
		t.Errorf("expected: %v, got: %v", 1.414, d)
	}
}

func TestCharacterGrid_Line(t *testing.T) {
	grid := testGrid()
	p0 := Point{X: 0, Y: 0}
	p1 := Point{X: 1, Y: 2}
	line := grid.Line(p0, p1)

	expectedPoints := []Point{{X: 0, Y: 0}, {X: 1, Y: 2}, {X: 2, Y: 4}}
	expected := Line{
		Grid:   &grid,
		Points: NewSet(expectedPoints),
	}

	if !reflect.DeepEqual(line, expected) {
		t.Errorf("expected %v, got %v", expected, line)
		t.Fail()
	}
}

func TestSlope_Reduce_WhenXgtY(t *testing.T) {
	slope := Slope{X: 4, Y: 2}
	r := slope.Reduce()
	if !reflect.DeepEqual(r, Slope{X: 2, Y: 1}) {
		t.Fail()
	}
}

func TestSlope_Reduce_WhenXltY(t *testing.T) {
	slope := Slope{X: 4, Y: 12}
	r := slope.Reduce()
	if !reflect.DeepEqual(r, Slope{X: 1, Y: 3}) {
		t.Fail()
	}
}

func TestSlope_Reduce_WithPoints_Negative(t *testing.T) {
	p0 := Point{X: 2, Y: 2}
	p1 := Point{X: 1, Y: 6}
	m := NewSlope(p0, p1)
	if !reflect.DeepEqual(m, Slope{X: -1, Y: 4}) && !reflect.DeepEqual(m, Slope{X: 1, Y: -4}) {
		// Negative sign should be on the Y here, but could end up on either
		t.Errorf("expected: %v, got: %v", Slope{X: 1, Y: -4}, m)
		t.Fail()
	}
}

func TestSlope_Reduce_WithPoints_Positive(t *testing.T) {
	p0 := Point{X: 1, Y: 1}
	p1 := Point{X: 2, Y: 4}
	m := NewSlope(p0, p1)
	if !reflect.DeepEqual(m, Slope{X: 1, Y: 3}) {
		t.Errorf("expected: %v, got: %v", Slope{X: 1, Y: 3}, m)
		t.Fail()
	}
}

func TestLine_WithProblemInputs(t *testing.T) {
	// line="{Points:{data:map[{0 2}:{X:0 Y:2} {1 2}:{X:1 Y:2} {2 1}:{X:2 Y:1} {3 1}:{X:3 Y:1} {4 0}:{X:4 Y:0} {5 0}:{X:5 Y:0}]} Grid:0x14000110100}"
	// points="[{X:3 Y:1} {X:1 Y:2}]"
	// m = -1/2
	// b = 1/2
	grid := testGrid()
	p0 := Point{X: 3, Y: 1}
	p1 := Point{X: 1, Y: 2}
	l := grid.Line(p0, p1)

	if !l.Points.Contains(Point{X: 5, Y: 0}) {
		t.Errorf("expected: %v to contain %v", l.Points, Point{X: 5, Y: 0})
		t.Fail()
	}
}
