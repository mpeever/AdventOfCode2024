package main

import (
	"bufio"
	"errors"
	"log/slog"
	"os"
	"strings"
)

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
	UPLEFT
	UPRIGHT
	DOWNLEFT
	DOWNRIGHT
)

type size int
type width int
type height int

type Point struct {
	X width
	Y height
}

type CharacterGrid struct {
	Min, Max Point
	Content  [][]string
}

type Vector struct {
	Points []Point
	Grid   *CharacterGrid
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
	return grid.Content[p.X][p.Y]
}

func (grid *CharacterGrid) NextPoint(p0 Point, direction Direction) (p Point, err error) {
	if !grid.Includes(p0) {
		err = errors.New("point p0 is not on grid")
		return
	}

	dispatch := make(map[Direction]func(p0 Point) Point)
	dispatch[RIGHT] = func(p0 Point) Point { return Point{p0.X + 1, p0.Y} }
	dispatch[LEFT] = func(p0 Point) Point { return Point{p0.X - 1, p0.Y} }
	dispatch[UP] = func(p0 Point) Point { return Point{p0.X, p0.Y + 1} }
	dispatch[DOWN] = func(p0 Point) Point { return Point{p0.X, p0.Y - 1} }
	dispatch[DOWNLEFT] = func(p0 Point) Point { return Point{p0.X - 1, p0.Y - 1} }
	dispatch[DOWNRIGHT] = func(p0 Point) Point { return Point{p0.X + 1, p0.Y - 1} }
	dispatch[UPLEFT] = func(p0 Point) Point { return Point{p0.X - 1, p0.Y + 1} }
	dispatch[UPRIGHT] = func(p0 Point) Point { return Point{p0.X + 1, p0.Y + 1} }

	pTemp := dispatch[direction](p0)

	if !grid.Includes(pTemp) {
		err = errors.New("point pTemp is not on grid")
		return
	}
	p = pTemp
	return
}

func (grid *CharacterGrid) Vector(p0 Point, d Direction, length size) (v Vector, err error) {
	if !grid.Includes(p0) {
		err = errors.New("point p0 is not on grid")
		return
	}
	slog.Debug("creating Vector from point", "p0", p0, "direction", d, "length", length)

	points := []Point{p0}
	pCur := p0

	for size(len(points)) < length && grid.Includes(pCur) {
		pTemp, err := grid.NextPoint(pCur, d)
		if err != nil {
			//slog.Error("error getting next point", "pCurrent", pCur, "err", err)
			return v, err
		}
		pCur = pTemp
		points = append(points, pCur)
	}
	v = Vector{Points: points, Grid: grid}

	return
}

func (v *Vector) size() size {
	return size(len(v.Points))
}

func (v *Vector) String() string {
	buffer := make([]string, v.size())
	for p := range v.Points {
		b := v.Grid.Char(v.Points[p])
		buffer = append(buffer, b)
	}
	return strings.Join(buffer, "")
}

// word search, ugh
func main() {
	word := "XMAS"
	length := size(len(word))
	slog.Debug("target word length", "length", length, "word", word)

	vectors := []Vector{}

	stdioScanner := bufio.NewScanner(os.Stdin)
	lines := [][]string{}
	for stdioScanner.Scan() {
		line := stdioScanner.Text()
		bLine := strings.Split(line, "")
		lines = append(lines, bLine)
	}
	yMax := height(len(lines))
	xMax := width(len(lines[0]))
	grid := CharacterGrid{
		Min:     Point{0, 0},
		Max:     Point{xMax - 1, yMax - 1},
		Content: lines,
	}

	slog.Info("generated grid", "grid", grid)

	for y, l := range grid.Content {
		for x := range l {
			p := Point{X: width(x), Y: height(y)}
			if !strings.EqualFold(grid.Char(p), string(word[0])) {
				// We only care if we found the first character in our word
				continue
			}

			// I can't find the Golang equivalent of Direction.values
			for d := range []Direction{UP, DOWN, LEFT, RIGHT, UPLEFT, DOWNRIGHT, UPRIGHT, DOWNLEFT} {
				vTemp, err := grid.Vector(p, Direction(d), length)
				if err != nil {
					slog.Debug("skipping error Vector", "err", err, "p0", p, "direction", d)
					continue
				}
				if !strings.EqualFold(vTemp.String(), word) {
					slog.Debug("skipping Vector because of word mismatch", "p0", p, "direction", d, "value", vTemp.String())
					continue
				}
				slog.Debug("storing matching vector", "vector", vTemp, "value", vTemp.String())
				vectors = append(vectors, vTemp)
			}
		}
	}

	slog.Info("found word matches", "count", len(vectors))
}
