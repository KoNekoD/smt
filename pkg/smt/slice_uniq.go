package smt

import "golang.org/x/exp/maps"

// SliceUniq - remove duplicates from slice
func SliceUniq[T comparable](slice []T) []T {
	keys := make(map[T]any)

	for _, entry := range slice {
		keys[entry] = nil
	}

	return maps.Keys(keys)
}
