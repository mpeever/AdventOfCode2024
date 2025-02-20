package main

import (
	. "AdventOfCode2024/lib"
	"bufio"
	"flag"
	"log/slog"
	"os"
	"regexp"
	"slices"
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

func (eq *Equation) Eval() (val int, err error) {
	var operator string
	var i int

	for _, input := range eq.Inputs {

		if input == "+" || input == "*" {
			operator = input
			continue
		}

		i, err = strconv.Atoi(input)
		if err != nil {
			return
		}

		if val == 0 {
			val = i
			continue
		}

		if operator == "+" {
			val = i + val
			continue
		}

		val = i * val
	}

	return
}

func (eq *Equation) Permutations() (output []Equation) {
	perms := permutations(eq.Inputs)
	for _, perm := range perms {
		eqn := eq.Clone()
		eqn.Inputs = perm
		output = append(output, eqn)
	}

	return
}

func (eq *Equation) Verify() bool {
	perms := eq.Permutations()
	return Any(perms, func(eq Equation) bool {
		value, err := eq.Eval()
		if err != nil {
			return false
		}
		return value == eq.Expected
	})
}

func Permutations(input []string) (output [][]string) {
	buffer := make([]string, len(input))
	copy(buffer, input)
	slices.Reverse(buffer)

	perms := permutations(buffer)
	for _, perm := range perms {
		p := make([]string, len(perm))
		copy(p, perm)
		slices.Reverse(p)
		output = append(output, p)
	}

	return
}

func permutations(str []string) (output [][]string) {
	if len(str) == 0 {
		output = append(output, str)
		return
	}

	if len(str) == 2 {
		for _, operation := range []string{"+", "*"} {
			perm := make([]string, len(str)+1)
			perm[0] = str[0]
			perm[1] = operation
			perm[2] = str[1]
			output = append(output, perm)
		}
		return
	}

	slog.Debug("permutations", "len(str)", len(str))

	if len(str) > 2 {
		partials := permutations(str[1:])
		for _, partial := range partials {
			for _, operation := range []string{"+", "*"} {
				perm := []string{str[0], operation}
				perm = append(perm, partial...)
				output = append(output, perm)
			}
		}
	}

	return
}

func puzzle1(eqns []Equation) (sum int) {
	valid := RemoveIfNot(eqns, func(e Equation) bool {
		return e.Verify()
	})

	slog.Debug("puzzle1", "valid", valid)

	for _, eq := range valid {
		sum += eq.Expected
	}

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
