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
