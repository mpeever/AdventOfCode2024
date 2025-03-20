package main

import (
	. "AdventOfCode2024/lib"
	"bufio"
	"flag"
	"log/slog"
	"os"
	"strings"
)

func Demux(g CharacterGrid) (output map[string]CharacterGrid) {
	output = make(map[string]CharacterGrid)

	// first, find all the characters in the map
	for _, row := range g.Content {
		for _, val := range row {
			if val == "." {
				// we don't care about non-antenna locations
				continue
			}
			if _, ok := output[val]; ok {
				// we already have this one in our map
				continue
			}
			output[val] = g.Clone()
		}
	}

	// so far, so good... now we need to scrub each grid of other characters
	for key, _ := range output {
		grid := output[key]
		for y, row := range grid.Content {
			for x, _ := range row {
				p := Point{X: Width(x), Y: Height(y)}
				if grid.Char(p) == "." || grid.Char(p) == key {
					// this is a valid Char
					continue
				}
				grid.Update(p, ".")
			}
		}
	}

	return
}

type Antinode struct {
	grid       CharacterGrid
	Char       string
	P0, P1, P2 Point
}

func (a *Antinode) Valid() bool {
	d0, err := a.grid.FloatDistance(a.P0, a.P1)
	if err != nil {
		return false
	}
	d1, err := a.grid.FloatDistance(a.P0, a.P2)
	if err != nil {
		return false
	}
	slog.Debug("antinode valid?", "d0", d0, "d1", d1)

	return float64(d0) == float64(d1)/2
}

func puzzle1(g CharacterGrid) int {
	allAntiNodes := NewSet[Point]([]Point{})

	for char, grid := range Demux(g) {
		antennaNodes := RemoveIfNot[Point](grid.AllPoints(), func(p Point) bool {
			return grid.Char(p) == char
		})

		slog.Debug("all nodes", "len", len(antennaNodes))

		pairs := Pairs[Point](antennaNodes)
		for _, pair := range pairs {
			// get the line thru the pair of points
			line := grid.Line(pair[0], pair[1])
			slog.Debug("line through", "line", line, "points", pair)

			for _, p := range line.Points.Values() {
				an := Antinode{
					grid: grid,
					Char: char,
					P1:   pair[0],
					P2:   pair[1],
					P0:   p,
				}
				if an.Valid() {
					slog.Debug("antinode valid", "antinode", an)
					allAntiNodes.Add(an.P0)
				}
			}
		}
	}

	slog.Debug("valid Antinodes", "antinodes", allAntiNodes.Values())

	slog.Info("puzzle 1 found Antinodes", "count", allAntiNodes.Size())

	return allAntiNodes.Size()
}

func puzzle2(g CharacterGrid) int {
	allAntiNodes := NewSet[Point]([]Point{})

	for char, grid := range Demux(g) {
		antennaNodes := RemoveIfNot[Point](grid.AllPoints(), func(p Point) bool {
			return grid.Char(p) == char
		})

		slog.Debug("all nodes", "len", len(antennaNodes))

		pairs := Pairs[Point](antennaNodes)

		for _, pair := range pairs {
			// get the line thru the pair of points
			line := grid.Line(pair[0], pair[1])
			slog.Debug("line through", "line", line, "points", pair)

			for _, p := range line.Points.Values() {
				allAntiNodes.Add(p)
			}
		}
	}

	slog.Debug("all Antinodes", "antinodes", allAntiNodes.Values())

	slog.Info("puzzle 2 found Antinodes", "count", allAntiNodes.Size())

	return allAntiNodes.Size()
}

func main() {
	flag.BoolFunc("debug", "enable debug logging", func(s string) (err error) {
		slog.SetLogLoggerLevel(slog.Level(slog.LevelDebug))
		return
	})

	flag.Parse()

	stdioScanner := bufio.NewScanner(os.Stdin)
	lines := [][]string{}
	for stdioScanner.Scan() {
		line := stdioScanner.Text()
		bLine := strings.Split(line, "")
		lines = append(lines, bLine)
	}
	yMax := Height(len(lines))
	xMax := Width(len(lines[0]))
	grid := CharacterGrid{
		Min:     Point{0, 0},
		Max:     Point{xMax - 1, yMax - 1},
		Content: lines,
	}

	slog.Debug("generated grid", "grid", grid)

	sum := puzzle1(grid)
	slog.Info("found puzzle1 sum", "sum", sum)

	sum = puzzle2(grid)
	slog.Info("found puzzle2 sum", "sum", sum)
}
