package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func sortAndReport(col []int, output chan int) {
	defer func() {
		slog.Info("closing channel A")
		close(output)
	}()
	slices.Sort(col)
	for _, i := range col {
		output <- i
	}
}

func main() {
	column_a := []int{}
	column_b := []int{}
	diffs := int(0)

	re := regexp.MustCompile("^(\\d+)\\s+(\\d+)$")

	scanner := bufio.NewScanner(os.Stdin)

	// we're just going to slurp the whole input
	for scanner.Scan() {
		line := scanner.Text()

		if !re.MatchString(line) {
			slog.Info("skipping non-matching line")
			continue
		}

		nums := re.FindAllStringSubmatch(line, 3)

		a, err := strconv.Atoi(nums[0][1])
		if err != nil {
			slog.Error("error finding Range numbers", "error", err)
		}
		b, err := strconv.Atoi(nums[0][2])
		if err != nil {
			slog.Error("error finding Range numbers", "error", err)
		}

		column_a = append(column_a, a)
		column_b = append(column_b, b)
	}

	chan_a := make(chan int)
	chan_b := make(chan int)

	// ugly, but it works: sort both columns
	go sortAndReport(column_a, chan_a)
	go sortAndReport(column_b, chan_b)

	// loop over the sorted channels, comparing each pair as it arrives
	for a := range chan_a {
		b := <-chan_b
		slog.Info("comparing A to B", "A", a, "B", b)
		diffs += diff(a, b)
	}

	fmt.Println("Found diffs:", diffs)
}
