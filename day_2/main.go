package main

import (
	. "AdventOfCode2024/lib"
	"bufio"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

type Level int
type Step int
type Report []Level

// Delta is always b - a
func delta(a, b Level) Step {
	delta := b - a
	slog.Debug("calculating delta", "a", a, "b", b, "delta", delta)
	return Step(delta)
}

// Calculate a list as a series of deltas
func deltas(input []Level, accumulator []Step) []Step {
	if len(input) < 2 {
		slog.Debug("returning accumlator", "accumlator", accumulator)
		return accumulator
	}

	car := input[0]
	cdr := input[1:]

	d := delta(cdr[0], car)
	accumulator = append(accumulator, d)

	return deltas(cdr, accumulator)
}

func isSafe(input Report, size int) (bool, []Step) {
	acc := deltas(input, []Step{})

	// first, we check that ALL of the deltas are either above or below zero
	if !(All(acc, func(i Step) bool { return i > 0 }) || All(acc, func(i Step) bool { return i < 0 })) {
		return false, acc
	}

	// OK, so All our deltas are in the same direction, now we need to find if Any exceed the size limit
	for _, d := range acc {
		if d > Step(size) || d < Step(0-size) {
			return false, acc
		}
	}

	return true, acc
}

// From a Report, generate a list of smaller Reports by removing each Level once
func perforate(report Report) (perf []Report) {
	for idx, level := range report {
		slog.Debug("removing level", "idx", idx, "level", level)
		if idx == 0 {
			p := make([]Level, len(report)-1)
			copy(p, report[1:])
			perf = append(perf, p)
		} else if idx == len(report)-1 {
			p := make([]Level, len(report)-1)
			copy(p, report[:idx])
			perf = append(perf, p)
		} else {
			front := make([]Level, idx)
			copy(front, report[:idx])
			back := make([]Level, len(report)-idx-1)
			copy(back, report[idx+1:])
			slog.Debug("combining front and back", "front", front, "back", back)
			p := append(front, back...)
			perf = append(perf, p)
		}
	}
	slog.Debug("perforations", "original", report, "perforations", perf)
	return
}

// Parse input lines into a series of Reports
func parse(input []string) (reports []Report) {
	for _, report := range input {
		var levels []Level
		for _, level := range strings.Fields(report) {
			value, err := strconv.Atoi(level)
			if err != nil {
				slog.Error("cannot convert String to Int", "level", level)
				continue
			}
			levels = append(levels, Level(value))
		}
		reports = append(reports, Report(levels))
	}

	slog.Info("parsed report count", "reports", reports, "count", len(reports))

	return
}

// a report is safe if every level changes in the same direction, none have zero deltas, and no delta exceeds 3
func puzzle1(input []string) (output int, err error) {
	safetyMap := make(map[int]bool)
	reports := parse(input)
	for idx, report := range reports {
		safe, _ := isSafe(report, 3)
		safetyMap[idx] = safe
	}

	// we start with no "safe" levels
	output = 0

	slog.Debug("safe levels", "safetyMap", safetyMap)

	for _, safe := range safetyMap {
		if safe {
			output += 1
		}
	}

	return
}

// a report is safe if a single level can be removed to make it safe
func puzzle2(input []string) (safeCount int, err error) {
	reports := parse(input)

	for _, report := range reports {
		safe, _ := isSafe(report, 3)
		if safe {
			safeCount++
			continue
		}

		slog.Debug("report to perforate", "report", report)

		subreports := perforate(report)
		if Any(subreports, func(r Report) bool { s, _ := isSafe(r, 3); return s }) {
			safeCount++
		}
	}

	return
}

func main() {
	stdioScanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	// we're just going to slurp the whole input
	for stdioScanner.Scan() {
		line := stdioScanner.Text()
		lines = append(lines, line)
	}

	// Puzzle 1
	count, err := puzzle1(lines)
	if err != nil {
		slog.Error("puzzle1 failed", "err", err)
	}
	slog.Info("Puzzle 1: found safe lines", "count", count)

	// Puzzle 2
	count, err = puzzle2(lines)
	if err != nil {
		slog.Error("puzzle2 failed", "err", err)
	}
	slog.Info("Puzzle 2: found safe lines", "count", count)
}
