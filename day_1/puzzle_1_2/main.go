package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strconv"
)

func main() {
	numbers := []int{}
	frequencies := make(map[int]int)

	similarityScore := int(0)

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
			slog.Error("error finding column A numbers", "error", err)
		}
		numbers = append(numbers, a)

		b, err := strconv.Atoi(nums[0][2])
		if err != nil {
			slog.Error("error finding column B numbers", "error", err)
		}

		_, found := frequencies[b]
		if !found {
			frequencies[b] = 0
		}
		frequencies[b] += 1
	}

	for _, num := range numbers {
		freq, found := frequencies[num]
		if !found {
			freq = 0
		}

		score := num * freq
		similarityScore += score
	}

	fmt.Println("Found similarityScore:", similarityScore)
}
