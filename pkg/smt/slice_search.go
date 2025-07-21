package smt

// SliceSearch returns index of needle in haystack or -1 if not found
func SliceSearch[T comparable](needle T, haystack []T) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}
	return -1
}
