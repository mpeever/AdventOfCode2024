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
	eq0 := Equation{Expected: 123, Inputs: []string{"41", "3"}, Operators: []string{ADD, MULT}}
	permutations := eq0.Permutations()

	if len(permutations) != 2 {
		t.Errorf("expected 2 permutations, got %d", len(permutations))
	}
}

func TestEquation_Permutations_ThreeInputs(t *testing.T) {
	eq0 := Equation{Expected: 123, Inputs: []string{"1", "2", "3"}, Operators: []string{ADD, MULT}}
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
	got := Permutations(input, []string{ADD, MULT})
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Permutations failed: expected %v, got %v", expected, got)
	}
}

func TestPermutations_TwoElementInput(t *testing.T) {
	input := []string{"1", "2"}
	expected := [][]string{[]string{"1", "+", "2"}, []string{"1", "*", "2"}}
	got := Permutations(input, []string{ADD, MULT})
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
	got := Permutations(input, []string{ADD, MULT})
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
	got := Permutations(input, []string{ADD, MULT})
	if len(expected) != len(got) {
		t.Errorf("Permutations failed: expected %v, got %v", expected, got)
	}
}

func TestEquation_Eval_TwoElementMult(t *testing.T) {
	eq0 := Equation{Expected: 123, Inputs: []string{"41", "*", "3"}, Operators: []string{ADD, MULT}}
	value, err := eq0.Eval()
	if err != nil {
		t.Errorf("Equation eval failed: %v", err)
	}
	if value != eq0.Expected {
		t.Errorf("Equation eval failed: expected %v, got %v", eq0.Expected, value)
	}
}

func TestEquation_Eval_TwoElementAdd(t *testing.T) {
	eq0 := Equation{Expected: 44, Inputs: []string{"41", "+", "3"}, Operators: []string{}}
	value, err := eq0.Eval()
	if err != nil {
		t.Errorf("Equation eval failed: %v", err)
	}
	if value != eq0.Expected {
		t.Errorf("Equation eval failed: expected %v, got %v", eq0.Expected, value)
	}
}

func TestEval_ADD(t *testing.T) {
	v0, _ := Eval([]string{"12", ADD, "13"})
	if v0 != 25 {
		t.Errorf("Eval failed: expected %v, got %v", 25, v0)
	}
}

func TestEval_MULT(t *testing.T) {
	v0, _ := Eval([]string{"12", MULT, "13"})
	if v0 != 156 {
		t.Errorf("Eval failed: expected %v, got %v", 156, v0)
	}
}

func TestEval_COMPLEX(t *testing.T) {
	v0, _ := Eval([]string{"12", "*", "13", "+", "5"})
	if v0 != 161 {
		t.Errorf("Eval failed: expected %v, got %v", 161, v0)
	}
}

func TestEval_CAT(t *testing.T) {
	v0, _ := Eval([]string{"12", CAT, "13"})
	if v0 != 1213 {
		t.Errorf("Eval failed: expected %v, got %v", 1213, v0)
	}
}

func TestEval_COMPLEX_2(t *testing.T) {
	v0, _ := Eval([]string{"1", "+", "2", "*", "3", "||", "4"})
	if v0 != 94 {
		t.Errorf("Eval failed: expected %v, got %v", 94, v0)
	}
}

func TestEquation_Verify(t *testing.T) {
	eq0 := Equation{Expected: 44, Inputs: []string{"41", "3"}, Operators: []string{ADD, MULT}}
	if !eq0.Verify() {
		t.Errorf("equation verify failed: %v", eq0)
	}
	eq1 := Equation{Expected: 123, Inputs: []string{"41", "3"}, Operators: []string{ADD, MULT}}
	if !eq1.Verify() {
		t.Errorf("equation verify failed: %v", eq1)
	}
	eq2 := Equation{Expected: 41, Inputs: []string{"41"}, Operators: []string{ADD, MULT}}
	if !eq2.Verify() {
		t.Errorf("equation verify failed: %v", eq2)
	}
	eq3 := Equation{Expected: 41, Inputs: []string{"4", "1"}, Operators: []string{ADD, MULT, CAT}}
	if !eq3.Verify() {
		t.Errorf("equation verify failed: %v", eq3)
	}
	eq4 := Equation{Expected: 123, Inputs: []string{"4", "1", "3"}, Operators: []string{ADD, MULT, CAT}}
	if !eq4.Verify() {
		t.Errorf("equation verify failed: %v", eq4)
	}
}
