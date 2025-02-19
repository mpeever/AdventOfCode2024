package main

import (
	//. "AdventOfCode2024/lib"
	"bufio"
	"flag"
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Equation struct {
	Expected int
	Inputs   []string
}

func (eq *Equation) Clone() Equation {
	cex := eq.Expected
	cin := make([]string, len(eq.Inputs))
	copy(cin, eq.Inputs)
	return Equation{cex, cin}
}

func (eq *Equation) Permutations() (output []Equation) {
	if len(eq.Inputs) == 0 {
		output = append(output, eq.Clone())
		return
	}

	if len(eq.Inputs) == 2 {
		for _, operation := range []string{"+", "*"} {
			eq1 := eq.Clone()
			inputs := make([]string, len(eq.Inputs)+1)
			inputs[0] = eq1.Inputs[0]
			inputs[1] = operation
			inputs[2] = eq1.Inputs[1]
			eq1.Inputs = inputs

			output = append(output, eq1)
		}
		return
	}

	return
}

func puzzle1(eqns []Equation) (sum int) {

	return
}

func puzzle2(eqns []Equation) (sum int) {

	return
}

func main() {
	flag.BoolFunc("debug", "enable debug logging", func(s string) (err error) {
		slog.SetLogLoggerLevel(slog.Level(slog.LevelDebug))
		return
	})

	flag.Parse()

	var equations []Equation

	eqnRe := regexp.MustCompile(`^(\d+):\s+(.+)$`)

	stdioScanner := bufio.NewScanner(os.Stdin)
	for stdioScanner.Scan() {
		line := stdioScanner.Text()
		if !eqnRe.MatchString(line) {
			slog.Info("skipping non-matching line")
			continue
		}

		nums := eqnRe.FindAllStringSubmatch(line, 2)

		expected, err := strconv.Atoi(nums[0][1])
		if err != nil {
			slog.Error("can't find expected number", "error", err)
		}
		inputs := strings.Split(nums[0][2], " ")

		eqn := Equation{
			Expected: expected,
			Inputs:   inputs,
		}
		slog.Debug("adding equation", "equation", eqn)

		equations = append(equations, eqn)
	}

	slog.Debug("found equations", "equation count", len(equations))

	sum := puzzle1(equations)
	slog.Info("found puzzle1 sum", "sum", sum)

	sum = puzzle2(equations)
	slog.Info("found puzzle2 sum", "sum", sum)
}
