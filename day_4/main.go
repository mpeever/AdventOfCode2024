package main

import (
	. "AdventOfCode2024/lib"
	"bufio"
	"flag"
	"log/slog"
	"os"
	"strings"
)

func puzzle1(grid CharacterGrid, word string) (vectors []Vector) {
	length := Size(len(word))
	slog.Debug("target word length", "length", length, "word", word)

	for y, l := range grid.Content {
		for x := range l {
			p := Point{X: Width(x), Y: Height(y)}
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

	return
}

func puzzle2(grid CharacterGrid, word string) (xMasCount int) {
	grid.PrettyPrint(true)

	for y, l := range grid.Content {
		for x := range l {
			p := Point{X: Width(x), Y: Height(y)}
			slog.Debug("examining point", "p", p, "l", l, "char", grid.Char(p))
			if !strings.EqualFold(grid.Char(p), string(word[1])) {
				// we KNOW the word here is "MAS", we want to match "A"
				slog.Debug("skipping point", "p", p, "char", grid.Char(p))
				continue
			}

			pointsAround := grid.Corners(p)
			if len(pointsAround) < 4 { // we're on an edge, where a solution can't exist
				slog.Info("skipping point, because we have too few corners", "p", p, "corner count", len(pointsAround))
				continue
			}

			diagonals, err := grid.Diagonals(p)
			if err != nil {
				slog.Info("skipping point because Diagonals threw an error", "p", p, "err", err)
				continue
			}

			if !All(diagonals, func(vector Vector) bool {
				s := vector.String()
				return strings.EqualFold(s, "MAS") || strings.EqualFold(s, "SAM")
			}) {
				slog.Info("skipping point, diagonals don't spell 'MAS' or 'SAM'", "p", p, "diagonals", diagonals)
				continue
			}
			xMasCount++
		}
	}
	return
}

// word search, ugh
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

	slog.Info("generated grid", "grid", grid)

	puzzle1Words := puzzle1(grid, "XMAS")
	slog.Info("found puzzle1 word matches", "count", len(puzzle1Words))

	puzzle2Words := puzzle2(grid, "MAS")
	slog.Info("found puzzle2 word matches", "count", puzzle2Words)
}
