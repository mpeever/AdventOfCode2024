package main

import (
	. "AdventOfCode2024/lib"
	"bufio"
	"errors"
	"flag"
	"log/slog"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const (
	ADD  = "+"
	MULT = "*"
	CAT  = "||"
)

type Equation struct {
	Expected  int
	Inputs    []string
	Operators []string
}

func (eq *Equation) Clone() Equation {
	cex := eq.Expected

	cin := make([]string, len(eq.Inputs))
	copy(cin, eq.Inputs)

	return Equation{cex, cin, eq.Operators}
}

func (eq *Equation) Formulae() (output [][]string) {
	perms := permutations(eq.Inputs, eq.Operators)
	for _, perm := range perms {
		output = append(output, perm)
	}
	return
}

func (eq *Equation) Verify() bool {
	formulae := eq.Formulae()
	return Any(formulae, func(f []string) bool {
		value, err := Eval(f)
		if err != nil {
			return false
		}
		return value == eq.Expected
	})
}

func Eval(str []string) (val int, err error) {
	stack := NewStack[string]()

	// simple dispatch table
	dt := make(map[string]func(a, b int) int)
	dt[ADD] = func(a, b int) int { return a + b }
	dt[MULT] = func(a, b int) int { return a * b }
	dt[CAT] = func(a, b int) int {
		astr := strconv.Itoa(a)
		bstr := strconv.Itoa(b)
		ostr := strings.Join([]string{astr, bstr}, "")
		o, _ := strconv.Atoi(ostr)
		return o
	}

	re := regexp.MustCompile(`\d+`)
	for _, el := range str {
		// just push operators onto the stack
		if el == ADD || el == MULT || el == CAT {
			stack.Push(el)
			continue
		}

		if re.MatchString(el) {
			// this is a number, we need to check what our last operator was
			last, ok := stack.Peek()
			if !ok {
				// empty stack, we can just push this onto it and go on
				stack.Push(el)
				continue
			}

			if last == ADD || last == MULT || last == CAT {
				fn := dt[last]

				// get the operator off the stack
				stack.Pop()

				a, ok := stack.Pop()
				if !ok {
					err = errors.New("empty stack")
					return
				}

				i, e := strconv.Atoi(a)
				if e != nil {
					err = e
					return
				}

				j, e := strconv.Atoi(el)
				if e != nil {
					err = e
					return
				}

				value := fn(i, j)

				stack.Push(strconv.Itoa(value))
				continue
			}
			// the last element in the stack isn't CAT, just push el and continue
			//stack.Push(el)
		}
	}
	s, ok := stack.Pop()
	if !ok {
		err = errors.New("empty stack")
		return
	}

	val, err = strconv.Atoi(s)

	return
}

func Permutations(input []string, operators []string) (output [][]string) {
	buffer := make([]string, len(input))
	copy(buffer, input)
	slices.Reverse(buffer)

	perms := permutations(buffer, operators)
	for _, perm := range perms {
		p := make([]string, len(perm))
		copy(p, perm)
		slices.Reverse(p)
		output = append(output, p)
	}

	return
}

func permutations(str []string, operators []string) (output [][]string) {
	if len(str) == 0 || len(str) == 1 {
		output = append(output, str)
		return
	}

	if len(str) == 2 {
		for _, operation := range operators {
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
		partials := permutations(str[1:], operators)
		for _, partial := range partials {
			for _, operation := range operators {
				perm := []string{str[0], operation}
				perm = append(perm, partial...)
				output = append(output, perm)
			}
		}
	}

	return
}

func puzzle1(eqns []Equation) (sum int) {
	slog.Debug("puzzle1", "eqns", eqns)

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
	slog.Debug("puzzle2", "eqns", eqns)

	channels := []chan Equation{}

	for _, eq := range eqns {
		channel := make(chan Equation)
		channels = append(channels, channel)
		go func() {
			if eq.Verify() {
				channel <- eq
			}
			close(channel)
		}()
	}

	for _, channel := range channels {
		eq, ok := <-channel
		if ok {
			sum += eq.Expected
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

	p1Equations := Map(equations, func(e Equation) Equation {
		e.Operators = []string{ADD, MULT}
		return e
	})

	sum := puzzle1(p1Equations)
	slog.Info("found puzzle1 sum", "sum", sum)

	p2Equations := Map(equations, func(e Equation) Equation {
		e.Operators = []string{ADD, MULT, CAT}
		return e
	})

	sum = puzzle2(p2Equations)
	slog.Info("found puzzle2 sum", "sum", sum)
}
