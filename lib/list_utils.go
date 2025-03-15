package lib

import "errors"

func All[T any](input []T, fn func(T) bool) bool {
	if len(input) == 0 {
		return true
	}
	if !fn(input[0]) {
		return false
	}

	return All[T](input[1:], fn)
}

func Any[T any](input []T, fn func(T) bool) bool {
	if len(input) == 0 {
		return false
	}
	if fn(input[0]) {
		return true
	}

	return Any[T](input[1:], fn)
}

func RemoveIf[T any](input []T, fn func(T) bool) (output []T) {
	for _, i := range input {
		if !fn(i) {
			output = append(output, i)
		}
	}
	return
}

func RemoveIfNot[T any](input []T, fn func(T) bool) (output []T) {
	for _, i := range input {
		if fn(i) {
			output = append(output, i)
		}
	}
	return
}

func Map[T any, U any](input []T, fn func(T) U) []U {
	output := []U{}
	for _, j := range input {
		output = append(output, fn(j))
	}
	return output
}

func Unique[T comparable](arr []T) []T {
	m := make(map[T]bool)
	for _, v := range arr {
		m[v] = true
	}
	var result []T
	for k := range m {
		result = append(result, k)
	}
	return result
}

func Center[T comparable](arr []T) (t T, err error) {
	length := len(arr)
	if length == 0 {
		err = errors.New("can't find center of an empty array")
		return
	}
	if length%2 == 0 {
		return arr[length/2-1], nil
	}
	return arr[length/2], nil
}

func Pairs[T comparable](arr []T) [][]T {
	pairs := [][]T{}
	for i, t0 := range arr {
		for j, t1 := range arr {
			if i == j {
				continue
			}
			pairs = append(pairs, []T{t0, t1})
		}
	}
	return pairs
}
