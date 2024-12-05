package main

import (
	"bufio"
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func puzzle1(input []string) int {
	// This works in Perl: /mul\((\d+),(\d+)\)/g
	re := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")

	slog.Debug("regexp", "regexp", re)
	sum := 0
	for lineNumber, line := range input {
		slog.Debug("checking out line", "lineNumber", lineNumber, "line", line)
		matches := re.FindAllStringSubmatch(line, -1)
		slog.Debug("regexp matches", "matches", matches)
		for idx, m := range matches {
			slog.Debug("examining submatches", "idx", idx, "submatches", m)
			a, err := strconv.Atoi(m[1])
			if err != nil {
				continue
			}
			b, err := strconv.Atoi(m[2])
			if err != nil {
				continue
			}
			sum += a * b
		}
	}
	return sum
}

func puzzle2(input []string) int {
	eligible := make([]string, 5000) // holds all the strings we need to check

	// Argh!! This is my first n+1 failures! It's only ONE input, so join the lines!
	line := strings.Join(input, "")

	fields := strings.Split(line, "don't()")

	// we always need to keep fields[0], as it comes before the first don't()
	eligible = append(eligible, fields[0])

	// for fields[1:], we only keep what comes after "do()"
	for _, field := range fields[1:] {
		slog.Info("checking out field", "field", field)
		// we could have multiple "do()" in our line, but we can't have "don't()"
		subfields := strings.Split(field, "do()")

		// so we reject subfields[0] and keep subfields[1:]
		eligible = append(eligible, subfields[1:]...)
	}

	// Now that we have the "do" line segments all in one place, send them to the last puzzle for processing.
	return puzzle1(eligible)
}

func main() {
	stdioScanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	// we're just going to slurp the whole input
	for stdioScanner.Scan() {
		line := stdioScanner.Text()
		lines = append(lines, line)
	}

	answer1 := puzzle1(lines)
	slog.Info("puzzle1 solution", "sum", answer1)

	answer2 := puzzle2(lines)
	slog.Info("puzzle2 solution", "sum", answer2)
}
