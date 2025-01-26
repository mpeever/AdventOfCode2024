package main

import (
	. "AdventOfCode2024/lib"
	"bufio"
	"log/slog"
	"os"
	"strings"
)

func puzzle1(grid CharacterGrid, word string) (vectors []Vector) {

	return
}

func puzzle2(grid CharacterGrid, word string) (xMasCount int) {

	return
}

// word search, ugh
func main() {
	if Any(os.Args, func(s string) bool { return strings.EqualFold(s, "debug") }) {
		slog.SetLogLoggerLevel(slog.Level(slog.LevelDebug))
	}

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
