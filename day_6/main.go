package main

import (
	. "AdventOfCode2024/lib"
	"bufio"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

func puzzle1(p0 Point, g CharacterGrid, orientation Orientation) (sum int) {
	guard := Guard{
		Position:    p0,
		Path:        []Point{p0},
		Grid:        g,
		Orientation: orientation,
		Cache:       make(map[string]GuardState),
	}
	slog.Debug("starting with guard", "guard", guard)

	guard.Grid.Update(p0, ".")

	guard.WalkToEdge(func(state *GuardState) bool {
		slog.Error("loop detected", "state", state)
		return true
	})
	slog.Debug("finished with guard", "guard", guard)

	for _, p := range guard.Path {
		slog.Debug("guard position", "position", p)
	}

	locationMap := make(map[string]Point)
	for _, l := range guard.Path {
		key := fmt.Sprintf("(%d, %d)", l.X, l.Y)
		locationMap[key] = l
	}

	return len(locationMap)
}

func puzzle2(p0 Point, g CharacterGrid, orientation Orientation) (sum int) {
	channels := []chan bool{}

	for y, row := range g.Content {
		for x, value := range row {
			if value != "#" {
				clone := g.Clone()
				point := Point{X: Width(x), Y: Height(y)}
				clone.Update(point, "#")
				guard := Guard{
					Position:    p0,
					Path:        []Point{p0},
					Grid:        clone,
					Orientation: orientation,
					Cache:       make(map[string]GuardState),
				}
				output := make(chan bool)
				channels = append(channels, output)
				go func() {
					guard.WalkToEdge(func(state *GuardState) bool {
						output <- true
						return false
					})

					close(output)
				}()
			}
		}
	}

	for i, channel := range channels {
		didLoop, ok := <-channel
		if didLoop && ok {
			slog.Debug("loop detected", "guard", i)
			sum++
		}
	}

	return
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

	var p0 Point
	var orientation Orientation

	// Find the guard's position on tha map
	for y, line := range grid.Content {
		for x, char := range line {
			if char == "^" {
				p0 = Point{X: Width(x), Y: Height(y)}
				orientation = NORTH
			}
		}
	}

	// Gotcha! Gotta scrub that starting point
	grid.Content[p0.X][p0.Y] = "."

	sum := puzzle1(p0, grid, orientation)
	slog.Info("found puzzle1 sum", "sum", sum)

	sum = puzzle2(p0, grid, orientation)
	slog.Info("found puzzle2 sum", "sum", sum)
}
