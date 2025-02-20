package main

import (
	"reflect"
	"testing"
)

func TestEquation_Clone(t *testing.T) {
	expected := 123
	inputs := []string{"a", "b", "c", "d", "e"}
	eq0 := Equation{Expected: expected, Inputs: inputs}
	eq1 := eq0.Clone()
	if !reflect.DeepEqual(eq0.Inputs, eq1.Inputs) {
		t.Errorf("Equation clone failed: expected %v, got %v", eq0.Inputs, eq1.Inputs)
	}
	eq1.Inputs[1] = "BLEAH"
	if reflect.DeepEqual(eq0.Inputs, eq1.Inputs) {
		t.Error("Equation clone failed")
	}
}

func TestEquation_Permutations_EmptyInputs(t *testing.T) {
	eq0 := Equation{Expected: 123, Inputs: []string{}}
	permutations := eq0.Permutations()
	if len(permutations) != 1 {
		t.Errorf("expected 1 permutation, got %d", len(permutations))
	}
}

func TestEquation_Permutations_TwoInputs(t *testing.T) {
	eq0 := Equation{Expected: 123, Inputs: []string{"41", "3"}}
	permutations := eq0.Permutations()

	if len(permutations) != 2 {
		t.Errorf("expected 2 permutations, got %d", len(permutations))
	}
}

func TestEquation_Permutations_ThreeInputs(t *testing.T) {
	eq0 := Equation{Expected: 123, Inputs: []string{"1", "2", "3"}}
	expected := []Equation{
		{Expected: 123, Inputs: []string{"1", "+", "2", "+", "3"}},
		{Expected: 123, Inputs: []string{"1", "*", "2", "+", "3"}},
		{Expected: 123, Inputs: []string{"1", "+", "2", "*", "3"}},
		{Expected: 123, Inputs: []string{"1", "*", "2", "*", "3"}},
	}
	permutations := eq0.Permutations()

	for i, p := range permutations {
		if !reflect.DeepEqual(p.Inputs, expected[i].Inputs) {
			t.Errorf("expected %v, got %v, %d", expected[i], p, i)
		}
	}
}

func TestPermutations_EmptyInput(t *testing.T) {
	input := []string{}
	expected := [][]string{input}
	got := Permutations(input)
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Permutations failed: expected %v, got %v", expected, got)
	}
}

func TestPermutations_TwoElementInput(t *testing.T) {
	input := []string{"1", "2"}
	expected := [][]string{[]string{"1", "+", "2"}, []string{"1", "*", "2"}}
	got := Permutations(input)
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Permutations failed: expected %v, got %v", expected, got)
	}
}

func TestPermutations_ThreeElementInput(t *testing.T) {
	input := []string{"1", "2", "3"}
	expected := [][]string{
		[]string{"1", "+", "2", "+", "3"},
		[]string{"1", "+", "2", "*", "3"},
		[]string{"1", "*", "2", "+", "3"},
		[]string{"1", "*", "2", "*", "3"},
	}
	got := Permutations(input)
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Permutations failed: expected %v, got %v", expected, got)
	}
}

func TestPermutations_FourElementInput(t *testing.T) {
	input := []string{"1", "2", "3", "4"}
	expected := [][]string{
		[]string{"1", "+", "2", "+", "3", "+", "4"},
		[]string{"1", "+", "2", "+", "3", "*", "4"},
		[]string{"1", "+", "2", "*", "3", "+", "4"},
		[]string{"1", "+", "2", "*", "3", "*", "4"},
		[]string{"1", "*", "2", "+", "3", "+", "4"},
		[]string{"1", "*", "2", "*", "3", "+", "4"},
		[]string{"1", "*", "2", "+", "3", "*", "4"},
		[]string{"1", "*", "2", "*", "3", "*", "4"},
	}
	got := Permutations(input)
	if len(expected) != len(got) {
		t.Errorf("Permutations failed: expected %v, got %v", expected, got)
	}
}

func TestEquation_Eval_TwpElementMult(t *testing.T) {
	eq0 := Equation{Expected: 123, Inputs: []string{"41", "*", "3"}}
	value, err := eq0.Eval()
	if err != nil {
		t.Errorf("Equation eval failed: %v", err)
	}
	if value != eq0.Expected {
		t.Errorf("Equation eval failed: expected %v, got %v", eq0.Expected, value)
	}
}

func TestEquation_Eval_TwpElementAdd(t *testing.T) {
	eq0 := Equation{Expected: 44, Inputs: []string{"41", "+", "3"}}
	value, err := eq0.Eval()
	if err != nil {
		t.Errorf("Equation eval failed: %v", err)
	}
	if value != eq0.Expected {
		t.Errorf("Equation eval failed: expected %v, got %v", eq0.Expected, value)
	}
}
