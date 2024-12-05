package main

import (
	"bufio"
	"log/slog"
	"os"
	"regexp"
	"strconv"
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

}
