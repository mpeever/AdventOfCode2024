package lib

import (
	"errors"
	"fmt"
	"log/slog"
	"math"
	"slices"
	"strings"
)

type Direction int

const ( // This is complicated, because we're not in the first quadrant
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
	UPLEFT
	UPRIGHT
	DOWNLEFT
	DOWNRIGHT
)

type Size int
type Width int
type Height int

type Point struct {
	X Width
	Y Height
}

type CharacterGrid struct {
	Min, Max Point
	Content  [][]string
}

type Vector struct {
	Points []Point
	Grid   *CharacterGrid
}

type Intersection struct {
	V0, V Vector
	Grid  *CharacterGrid
}

func (p *Point) Equals(other *Point) bool {
	return p.X == other.X && p.Y == other.Y
}

func (v *Vector) Size() Size {
	return Size(len(v.Points))
}

func (v *Vector) Equals(other *Vector) bool {
	if v.Size() != other.Size() {
		// if they're not the same size, they can't be equal
		return false
	}

	for _, p := range v.Points {
		if !slices.Contains(other.Points, p) {
			// if they don't both contain p, they can't be equal
			return false
		}
	}
	return true
}

func (v *Vector) String() string {
	buffer := make([]string, v.Size())
	for p := range v.Points {
		b := v.Grid.Char(v.Points[p])
		buffer = append(buffer, b)
	}
	return strings.Join(buffer, "")
}

func (v *Vector) Intersect(other *Vector) (bool, Point) {
	for _, p := range v.Points {
		for _, o := range other.Points {
			if p.Equals(&o) {
				return true, p
			}
		}
	}
	return false, Point{}
}

func (v *Vector) Colinear(other *Vector) bool {
	coinc := []Point{}
	for _, p := range v.Points {
		for _, o := range other.Points {
			if p.Equals(&o) {
				coinc = append(coinc, p)
			}
		}
	}
	return len(coinc) > 1
}

func (i *Intersection) Contains(v Vector) bool {
	return i.V0.Equals(&v) || i.V.Equals(&v)
}

func (i *Intersection) Equals(other *Intersection) bool {
	return other.Contains(i.V0) && other.Contains(i.V)
}

func (i *Intersection) String() string {
	return fmt.Sprintf("'%v X %v'", i.V0.String(), i.V.String())
}

func (grid *CharacterGrid) Clone() CharacterGrid {
	cloneRows := make([][]string, len(grid.Content))

	for y, row := range grid.Content {
		cloneRow := make([]string, len(row))
		copy(cloneRow, row)
		cloneRows[y] = cloneRow
	}

	return CharacterGrid{
		Min:     grid.Min,
		Max:     grid.Max,
		Content: cloneRows,
	}
}

func (grid *CharacterGrid) Includes(p Point) bool {
	if p.X >= grid.Min.X && p.X <= grid.Max.X && p.Y >= grid.Min.Y && p.Y <= grid.Max.Y {
		slog.Debug("point is on grid", "p", p, "min", grid.Min, "max", grid.Max)
		return true
	}
	slog.Debug("point is NOT on grid", "p", p, "min", grid.Min, "max", grid.Max)
	return false
}

func (grid *CharacterGrid) Char(p Point) string {
	return grid.Content[p.Y][p.X]
}

func (grid *CharacterGrid) Update(p Point, c string) {
	grid.Content[p.Y][p.X] = c
}

func (grid *CharacterGrid) NextPoint(p0 Point, direction Direction) (p Point, err error) {
	if !grid.Includes(p0) {
		err = errors.New("point p0 is not on grid")
		return
	}

	dispatch := make(map[Direction]func(p0 Point) Point)
	dispatch[RIGHT] = func(p0 Point) Point { return Point{p0.X + 1, p0.Y} }
	dispatch[LEFT] = func(p0 Point) Point { return Point{p0.X - 1, p0.Y} }
	dispatch[UP] = func(p0 Point) Point { return Point{p0.X, p0.Y - 1} }
	dispatch[DOWN] = func(p0 Point) Point { return Point{p0.X, p0.Y + 1} }
	dispatch[DOWNLEFT] = func(p0 Point) Point { return Point{p0.X - 1, p0.Y + 1} }
	dispatch[DOWNRIGHT] = func(p0 Point) Point { return Point{p0.X + 1, p0.Y + 1} }
	dispatch[UPLEFT] = func(p0 Point) Point { return Point{p0.X - 1, p0.Y - 1} }
	dispatch[UPRIGHT] = func(p0 Point) Point { return Point{p0.X + 1, p0.Y - 1} }

	pTemp := dispatch[direction](p0)

	if !grid.Includes(pTemp) {
		err = errors.New(fmt.Sprintf("point pTemp (%v) is not on grid", pTemp))
		return
	}
	p = pTemp
	return
}

func (grid *CharacterGrid) Distance(p0, p1 Point) (s Size, err error) {
	dx := int(math.Abs(float64(p0.X - p1.X)))
	dy := int(math.Abs(float64(p0.Y - p1.Y)))
	l := math.Hypot(float64(dx), float64(dy))
	s = Size(math.Floor(l))
	return
}

func (grid *CharacterGrid) Vector(p0 Point, d Direction, length Size) (v Vector, err error) {
	if !grid.Includes(p0) {
		err = errors.New("point p0 is not on grid")
		return
	}
	slog.Debug("creating Vector from point", "p0", p0, "direction", d, "length", length)

	points := []Point{p0}
	pCur := p0

	for Size(len(points)) < length && grid.Includes(pCur) {
		pTemp, err := grid.NextPoint(pCur, d)
		if err != nil {
			slog.Debug("error getting next point", "pCurrent", pCur, "err", err)
			return v, err
		}
		pCur = pTemp
		vpts := make([]Point, len(points))
		copy(vpts, points)
		points = append(vpts, pCur)
	}
	v = Vector{Points: points, Grid: grid}

	return
}

func (grid *CharacterGrid) DirectionOf(p0, p1 Point) (d Direction, err error) {
	dx := p1.X - p0.X
	dy := p1.Y - p0.Y

	if dx == 0 && dy == 0 {
		err = errors.New("can't calculate Direction of one Point to itself")
		return
	}

	// let's shortcut cardinal directions
	if dx == 0 {
		// we can only go up or down
		if dy > 0 {
			d = DOWN
		} else {
			d = UP
		}
		return
	} else if dy == 0 {
		if dx > 0 {
			d = RIGHT
		} else {
			d = LEFT
		}
		return
	}

	if dx > 0 && dy > 0 {
		d = DOWNRIGHT
	} else if dx > 0 && dy < 0 {
		d = UPRIGHT
	} else if dx < 0 && dy < 0 {
		d = DOWNLEFT
	} else if dx < 0 && dy > 0 {
		d = UPLEFT
	}

	return
}

func (grid *CharacterGrid) VectorThrough(p0, p1 Point, l Size) (v Vector, err error) {
	d, err := grid.DirectionOf(p0, p1)
	if err != nil {
		return
	}
	// if l is less than zero, calculate the length we need
	if l < 0 {
		l, err = grid.Distance(p0, p1)
		if err != nil {
			return
		}
	}
	v, err = grid.Vector(p0, d, l)
	if err != nil {
		return
	}

	return
}

func (grid *CharacterGrid) PointsAround(p0 Point) (points []Point) {
	for _, d := range []Direction{UP, DOWN, LEFT, RIGHT, UPLEFT, DOWNRIGHT, UPRIGHT, DOWNLEFT} {
		p, err := grid.NextPoint(p0, d)
		if err != nil {
			slog.Debug("error calculating Point", "p", p, "d", d, "err", err)
			continue
		}
		points = append(points, p)
	}
	return
}

func (grid *CharacterGrid) ValuesAround(p0 Point) []string {
	return Map(grid.PointsAround(p0), func(p Point) string { return grid.Char(p) })
}

func (grid *CharacterGrid) Corners(p0 Point) (points []Point) {
	for _, d := range []Direction{UPLEFT, DOWNRIGHT, UPRIGHT, DOWNLEFT} {
		p, err := grid.NextPoint(p0, d)
		if err != nil {
			slog.Debug("error calculating Point", "p", p, "d", d, "err", err)
			continue
		}
		points = append(points, p)
	}
	return
}

func (grid *CharacterGrid) CornerValues(p0 Point) []string {
	return Map(grid.Corners(p0), func(p Point) string { return grid.Char(p) })
}

func (grid *CharacterGrid) Diagonals(p0 Point) (diagonals []Vector, err error) {
	if !grid.Includes(p0) {
		err = errors.New("point p0 is not on grid")
		return
	}

	if p0.X == grid.Min.X || p0.Y == grid.Min.Y || p0.X == grid.Max.X || p0.Y == grid.Max.Y {
		err = errors.New("point p0 is on the edge of the grid, cannot compute diagonals")
		return
	}

	ul, err := grid.NextPoint(p0, UPLEFT)
	if err != nil {
		return
	}
	dr, err := grid.NextPoint(p0, DOWNRIGHT)
	if err != nil {
		return
	}

	v0 := Vector{
		Points: []Point{ul, p0, dr},
		Grid:   grid,
	}

	ur, err := grid.NextPoint(p0, UPRIGHT)
	if err != nil {
		return
	}
	dl, err := grid.NextPoint(p0, DOWNLEFT)
	if err != nil {
		return
	}

	v1 := Vector{
		Points: []Point{ur, p0, dl},
		Grid:   grid,
	}

	diagonals = []Vector{v0, v1}

	return
}

func (grid *CharacterGrid) PrettyPrint(debug bool) {
	for y, l := range grid.Content {
		line := []string{}
		for x, c := range l {
			if debug {
				line = append(line, fmt.Sprintf("%s (%d, %d)", c, x, y))
			} else {
				line = append(line, fmt.Sprintf("%s", c))
			}
		}
		println("\t", strings.Join(line, " "))
	}
}

func (grid *CharacterGrid) PrettyPrintPoints(points []Point, debug bool) {
	println("===================")
	for y, l := range grid.Content {
		line := []string{}
		for x, c := range l {
			p := Point{X: Width(x), Y: Height(y)}
			if !slices.Contains(points, p) {
				c = "."
			}
			if debug {
				line = append(line, fmt.Sprintf("%s (%d, %d)", c, x, y))
			} else {
				line = append(line, fmt.Sprintf("%s", c))
			}
		}
		println("\t", strings.Join(line, " "))
	}
	println("===================")
}

func Intersections(vs []Vector) (inxs []Intersection) {
	// this is basically a permutations problem: pairs of Vectors that intersect
	for i, v0 := range vs {
		for j, v := range vs {
			if i == j || v0.Colinear(&v) {
				// skip if we found ourselves or if v and v0 are colinear
				continue
			}
			ins, _ := v0.Intersect(&v)
			if !ins {
				// skip if we don't intersect
				continue
			}
			slog.Debug("intersection", "v0", v0, "v", v, "i", i, "j", j)
			ix0 := Intersection{V0: v0, V: v, Grid: v0.Grid}

			// If we already have this one, just skip
			if !Any(inxs, func(ix Intersection) bool { return ix.Equals(&ix0) }) {
				slog.Debug("intersection is new", "ix0", ix0)
				inxs = append(inxs, ix0)
			}
		}
	}
	return
}
