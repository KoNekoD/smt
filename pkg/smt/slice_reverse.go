package smt

import "slices"

// SliceReverse - Reverse a slice
// Initialization: i and j are initialized with the start and end indices of s, respectively.
//
//	“i” starts at index 0 (the first index of the slice).
//	“j” starts at index len(s)-1 (the last index of the slice).
//
// Loop condition: The loop continues until i will be less than j (i.e., they move to meet each other and until
// they meet and i passes an index farther than the middle and j crosses an index less than the middle).
//
// Swap the characters: Inside the loop, the characters at indices i and j are swapped.
// s[i], s[j] = s[j], s[i] is a concurrent assignment that swaps the values of s[i] and s[j].
//
// Increment/Decrement: After the swap, i will be incremented by 1 and j is decremented by 1 (the movement to
// meet each other described above).
//
// i, j = i+1, j-1 updates the indices for the next iteration.
//
// Repeat: Steps 2-4 are repeated until i will be less than j (each side will not pass its halfway
// point - they will do it simultaneously as the slice is divided in half between them), after which the loop ends.
//
// This algorithm has a time complexity of O(n/2), where n is the length of the string,
// making it efficient for reversing strings of any size.
//
// Example:
//
//   - [1 2 3 4 5] -> [5 4 3 2 1]
//   - Indexes [0,1,2,3,4]
//   - Iterations: swap [0, 4], swap [1, 3]
//   - So index in the middle not will be affected - it center index
//
// Example:
//
//   - [1 2 3 4 5 6] -> [6 5 4 3 2 1]
//   - Indexes [0,1,2,3,4,5]
//   - Iterations: swap [0, 5], swap [1, 4], swap [2, 3]
//   - It will end here because the next “i” will cross the half - the loop will not process it and will
//     return a reversed slice.
func SliceReverse[T any](s []T) []T {
	s = slices.Clone(s)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}
