package smt

// SliceFlip returns a map from values to indices
func SliceFlip[V comparable](ts []V) map[V]int {
	result := make(map[V]int, len(ts))

	for i, v := range ts {
		result[v] = i
	}

	return result
}
