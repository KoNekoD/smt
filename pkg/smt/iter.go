package smt

import "iter"

// IterToSlice converts an iterator into a slice
func IterToSlice[T any](iter iter.Seq[T]) []T {
	var result []T

	for v := range iter {
		result = append(result, v)
	}

	return result
}

// IterToMap converts an iterator into a map
func IterToMap[K comparable, V any](iter iter.Seq2[K, V]) map[K]V {
	result := make(map[K]V)

	for k, v := range iter {
		result[k] = v
	}

	return result
}
