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
	columnA := []int{}
	columnB := []int{}
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

		columnA = append(columnA, a)
		columnB = append(columnB, b)
	}

	chanA := make(chan int)
	chanB := make(chan int)

	// ugly, but it works: sort both columns
	go sortAndReport(columnA, chanA)
	go sortAndReport(columnB, chanB)

	// loop over the sorted channels, comparing each pair as it arrives
	for a := range chanA {
		b := <-chanB
		slog.Info("comparing A to B", "A", a, "B", b)
		diffs += diff(a, b)
	}

	fmt.Println("Found diffs:", diffs)
}
