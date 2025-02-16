package lib

import (
	"errors"
	"log/slog"
	"reflect"
	"slices"
)

type PageNumber int

type Rule struct {
	Earlier, Later PageNumber
}

// MatchesPages Check whether BOTH pages in our Rule is in this list of page numbers
func (r *Rule) MatchesPages(pages []PageNumber) bool {
	pp := NewSet(pages)
	return pp.Contains(r.Earlier) && pp.Contains(r.Later)
}

func (r *Rule) MatchesSection(s Section) bool {
	return r.MatchesPages(s.Pages)
}

type Section struct {
	Rules []Rule
	Pages []PageNumber
}

func (s *Section) AddRule(r Rule) int {
	s.Rules = append(s.Rules, r)
	slog.Debug("added rule", "section", s, "rule", r)
	return len(s.Rules)
}

func (s *Section) Sort() ([]PageNumber, error) {
	fn := func(p1, p2 PageNumber) int {
		rules := RemoveIfNot(s.Rules, func(rule Rule) bool {
			return rule.MatchesPages([]PageNumber{p1, p2})
		})
		for _, rule := range rules {
			if rule.MatchesPages(s.Pages) {
				slog.Debug("rule matches pages", "rule", rule, "pages", s.Pages)
				if rule.Earlier == p1 && rule.Later == p2 {
					slog.Debug("p1 < p2", "rule", rule, "p1", p1, "p2", p2)
					return -1
				}
				slog.Debug("p1 > p2", "rule", rule, "p1", p1, "p2", p2)
				return 1
			}
			// skip this Rule: shouldn't ever happen
		}
		slog.Error("Can't match rule", "section", s, "p1", p1, "p2", p2)
		return 0
	}

	pages := make([]PageNumber, len(s.Pages))
	i := copy(pages, s.Pages)
	if i < len(s.Pages) {
		return pages, errors.New("pages copy failed")
	}
	slices.SortFunc(pages, fn)
	return pages, nil
}

func (s *Section) IsCorrect() bool {
	sorted, err := s.Sort()
	if err != nil {
		slog.Error("sort failed", "err", err)
	}
	slog.Debug("sorted", "sorted", sorted, "pages", s.Pages)

	if reflect.DeepEqual(sorted, s.Pages) {
		slog.Debug("section correct", "section", s)
		return true
	}
	slog.Debug("section incorrect", "section", s)
	return false
}
