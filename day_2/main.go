package main

import (
	"bufio"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

// check if fn evaluates to true for all elements of a list.
// why doesn't Go come standard with decent list operators
func all(input []int, fn func(int) bool) bool {
	if len(input) == 0 {
		return true
	}
	car := input[0]
	cdr := input[1:]

	if !fn(car) {
		return false
	}

	return all(cdr, fn)
}

// Delta is always b - a
func delta(a, b int) int {
	delta := b - a
	slog.Debug("calculating delta", "a", a, "b", b, "delta", delta)
	return delta
}

// Calculate a list as a series of deltas
func deltas(input, accumulator []int) []int {
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

func isSafe(input []int, size int) (bool, []int) {
	acc := deltas(input, []int{})

	// first, we check that ALL of the deltas are either above or below zero
	if !(all(acc, func(i int) bool { return i > 0 }) || all(acc, func(i int) bool { return i < 0 })) {
		return false, acc
	}

	// OK, so all our deltas are in the same direction, now we need to find if any exceed the size limit
	for _, d := range acc {
		if d > size || d < (0-size) {
			return false, acc
		}
	}

	return true, acc
}

func puzzle1(input []string) (output int, err error) {
	safetyMap := make(map[int]bool)
	for idx, report := range input {
		levels := []int{}
		for _, level := range strings.Fields(report) {
			value, err := strconv.Atoi(level)
			if err != nil {
				slog.Error("cannot convert String to Int", "level", level)
				continue
			}
			levels = append(levels, value)
		}
		safe, _ := isSafe(levels, 3)
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

func puzzle2(input string) int {
	return 0
}

func main() {
	stdioScanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	// we're just going to slurp the whole input
	for stdioScanner.Scan() {
		line := stdioScanner.Text()
		lines = append(lines, line)
	}
	count, err := puzzle1(lines)
	if err != nil {
		slog.Error("puzzle1 failed", "err", err)
	}
	slog.Info("found safe lines", "count", count)
}
