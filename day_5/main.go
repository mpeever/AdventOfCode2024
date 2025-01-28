package main

import (
	. "AdventOfCode2024/lib"
	"bufio"
	"flag"
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	Earlier, Later int
}

// Matches Check whether BOTH pages in our Rule is in this list of page numbers
func (r *Rule) Matches(pages []int) bool {
	pp := NewSet(pages)
	return pp.Contains(r.Earlier) && pp.Contains(r.Later)
}

func puzzle1(rules []Rule, sections [][]int) (sum int) {
	applicableRules := make(map[int][]Rule)
	for sectionNumber, section := range sections {
		matchingRules := []Rule{}
		for _, rule := range rules {
			if rule.Matches(section) {
				slog.Debug("section contains both pages for Rule", "section", sectionNumber, "rule", rule)
				matchingRules = append(matchingRules, rule)
			}
		}
		applicableRules[sectionNumber] = matchingRules
	}
	slog.Info("applicableRules", "applicableRules", applicableRules)
	return
}

func puzzle2(rules []Rule, sections [][]int) (sortedSections [][]int) {

	return
}

// word search, ugh
func main() {
	flag.BoolFunc("debug", "enable debug logging", func(s string) (err error) {
		slog.SetLogLoggerLevel(slog.Level(slog.LevelDebug))
		return
	})

	flag.Parse()

	ruleLineRe := regexp.MustCompile(`^\s*(\d+)\|(\d+)\s*$`)
	sectionLineRe := regexp.MustCompile(`^\s*[\d,]+\s*$`)

	rules := []Rule{}
	sections := [][]int{}

	stdioScanner := bufio.NewScanner(os.Stdin)
	for stdioScanner.Scan() {
		line := stdioScanner.Text()

		if ruleLineRe.MatchString(line) {
			slog.Debug("line matches Rule pattern", "line", line)

			matches := ruleLineRe.FindStringSubmatch(line)
			first, err := strconv.Atoi(matches[1])
			second, err := strconv.Atoi(matches[2])
			if err != nil {
				slog.Error(err.Error())
				continue
			}
			rules = append(rules, Rule{Earlier: first, Later: second})
		} else if sectionLineRe.MatchString(line) {
			slog.Debug("line matches Section pattern", "line", line)

			strs := strings.Split(line, ",")
			pages := Map(strs, func(s string) int {
				i, err := strconv.Atoi(s)
				if err != nil {
					slog.Error(err.Error())
					return -1
				}
				return i
			})
			sections = append(sections, pages)
		} else {
			slog.Debug("line matches neither pattern", "line", line)
		}
	}

	slog.Debug("parsed input", "rules", rules, "sections", sections)

	sum := puzzle1(rules, sections)
	slog.Info("found puzzle1 word matches", "sum", sum)

}
