package smt

// MapSlice applies a function to each element of a slice
func MapSlice[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))

	for i, t := range ts {
		result[i] = fn(t)
	}

	return result
}

// Map applies function to each value of map
func Map[T1 comparable, T2, V any](ts map[T1]T2, fn func(T1, T2) V) map[T1]V {
	result := make(map[T1]V)
	for k, v := range ts {
		result[k] = fn(k, v)
	}
	return result
}
