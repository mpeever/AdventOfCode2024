package lib

import (
	"reflect"
	"testing"
)

func TestRule_Matches_False(t *testing.T) {
	p0 := PageNumber(23)
	p1 := PageNumber(47)
	p2 := PageNumber(12)
	rule := Rule{
		Earlier: p0,
		Later:   p2,
	}

	if rule.MatchesPages([]PageNumber{p0, p1}) {
		t.Fail()
	}

	if rule.MatchesPages([]PageNumber{p1, p2}) {
		t.Fail()
	}

	if !rule.MatchesPages([]PageNumber{p2, p0}) {
		t.Fail()
	}
}

func TestSection_Sort(t *testing.T) {
	rules := []Rule{
		Rule{Earlier: PageNumber(23), Later: PageNumber(47)},
		Rule{Earlier: PageNumber(23), Later: PageNumber(12)},
		Rule{Earlier: PageNumber(47), Later: PageNumber(18)},
		Rule{Earlier: PageNumber(12), Later: PageNumber(47)},
	}
	pages := []PageNumber{
		PageNumber(23), PageNumber(47), PageNumber(12), PageNumber(18),
	}

	section := Section{Rules: rules, Pages: pages}

	expected := []PageNumber{
		PageNumber(23), PageNumber(12), PageNumber(47), PageNumber(18),
	}

	actual, _ := section.Sort()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %v\nGot: %v", expected, actual)
		t.Fail()
	}
}
func TestSection_IsCorrect_Correct(t *testing.T) {
	rules := []Rule{
		Rule{Earlier: PageNumber(23), Later: PageNumber(47)},
		Rule{Earlier: PageNumber(23), Later: PageNumber(12)},
		Rule{Earlier: PageNumber(47), Later: PageNumber(18)},
		Rule{Earlier: PageNumber(12), Later: PageNumber(47)},
	}
	pages := []PageNumber{
		PageNumber(23), PageNumber(12), PageNumber(47), PageNumber(18),
	}

	section := Section{Rules: rules, Pages: pages}

	if !section.IsCorrect() {
		t.Fail()
	}
}

func TestSection_IsCorrect_InCorrect(t *testing.T) {
	rules := []Rule{
		Rule{Earlier: PageNumber(23), Later: PageNumber(47)},
		Rule{Earlier: PageNumber(23), Later: PageNumber(12)},
		Rule{Earlier: PageNumber(47), Later: PageNumber(18)},
		Rule{Earlier: PageNumber(12), Later: PageNumber(47)},
	}
	pages := []PageNumber{
		PageNumber(23), PageNumber(18), PageNumber(47), PageNumber(12),
	}

	section := Section{Rules: rules, Pages: pages}

	if section.IsCorrect() {
		t.Fail()
	}
}
