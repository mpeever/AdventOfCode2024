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

func puzzle1(rules []Rule, sections []Section) (sum int) {
	allSections := []Section{}
	for sectionNumber, section := range sections {
		for _, rule := range rules {
			if rule.MatchesSection(section) {
				slog.Debug("section contains both pages for Rule", "section", sectionNumber, "rule", rule)
				section.AddRule(rule)
			}
		}
		slog.Debug("found applicableRules", "section", section)
		allSections = append(allSections, section)
	}

	slog.Debug("added section rules", "sections", allSections)

	correct := RemoveIfNot(allSections, func(s Section) bool {
		return s.IsCorrect()
	})

	sum = 0
	for _, section := range correct {
		// We only have the correct Sections here
		pages := section.Pages
		center, err := Center(pages)
		if err != nil {
			slog.Error("error finding center", "error", err)
		}
		slog.Debug("section contains pages", "pages", pages, "center", center)

		sum = sum + int(center)
	}

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
	sections := []Section{}

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
			rules = append(rules, Rule{Earlier: PageNumber(first), Later: PageNumber(second)})
		} else if sectionLineRe.MatchString(line) {
			slog.Debug("line matches Section pattern", "line", line)

			strs := strings.Split(line, ",")
			ints := Map(strs, func(s string) int {
				i, err := strconv.Atoi(s)
				if err != nil {
					slog.Error(err.Error())
					return -1
				}
				return i
			})
			var pages []PageNumber
			for _, i := range ints {
				if i != -1 {
					pages = append(pages, PageNumber(i))
				}
			}
			sections = append(sections, Section{Pages: pages, Rules: []Rule{}})
		} else {
			slog.Debug("line matches neither pattern", "line", line)
		}
	}

	slog.Debug("parsed input", "rules", rules, "sections", sections)

	sum := puzzle1(rules, sections)
	slog.Info("found puzzle1 word matches", "sum", sum)

}
