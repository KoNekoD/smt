package smt

import (
	"slices"
	"testing"
)

func TestSafeSlice(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	slices.Equal([]int{2, 3, 4}, SafeSlice(arr, 1, 3))
	slices.Equal([]int{4, 5}, SafeSlice(arr, -2, 2))
	slices.Equal([]int{1, 2, 3}, SafeSlice(arr, 0, -2))
	slices.Equal([]int{3, 4}, SafeSlice(arr, -3, -1))
	slices.Equal([]int{}, SafeSlice(arr, 10, 3))

	sliceOfStrings := []string{"a", "b", "c", "d", "e"}
	slices.Equal([]string{"c", "d", "e"}, SafeSlice(sliceOfStrings, 2, 0))
	slices.Equal([]string{"d"}, SafeSlice(sliceOfStrings, -2, 1))
	slices.Equal([]string{"a", "b", "c"}, SafeSlice(sliceOfStrings, 0, 3))

	slices.Equal([]string{"a", "b", "c"}, SafeSlice(sliceOfStrings, 0, 3))
}
