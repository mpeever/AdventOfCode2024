package lib

func All[T any](input []T, fn func(T) bool) bool {
	if len(input) == 0 {
		return true
	}
	if !fn(input[0]) {
		return false
	}

	return All(input[1:], fn)
}

func Any[T any](input []T, fn func(T) bool) bool {
	if len(input) == 0 {
		return false
	}
	if fn(input[0]) {
		return true
	}

	return Any(input[1:], fn)
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
