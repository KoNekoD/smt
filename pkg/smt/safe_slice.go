package smt

// SafeSlice - Returns a slice of a slice.
//
// s - The input slice.
//
// offset - The offset of the slice.
//
//	If offset is non-negative, the sequence will start at that offset in the slice.
//
//	If offset is negative, the sequence will start that far from the end of the slice.
//
//	    Note:
//
//	    The offset parameter denotes the position in the slice, not the key.
//
// length - The length of the slice.
//
//	If length is given and is positive, then the sequence will have up to that many elements in it.
//
//	If the slice is shorter than the length, then only the available slice elements will be present.
//
//	If length is given and is negative then the sequence will stop that many elements from the end of the slice.
//
//	If it is omitted(as 0), then the sequence will have everything from offset up until the end of the slice.
func SafeSlice[T any](s []T, offset, length int) []T {
	sliceLength := len(s)

	// Adjust offset for negative values
	if 0 > offset {
		offset = sliceLength + offset // Take the index as if he were counting from the end.
	}

	// If offset is out of bounds, return an empty slice
	if 0 > offset || offset >= sliceLength {
		return []T{}
	}

	// Determine the end position
	var endPosition int
	if length > 0 { // Positive length is the normal case
		endPosition = offset + length
	} else if 0 > length {
		endPosition = sliceLength + length
	} else {
		endPosition = sliceLength
	}

	// Clamp end position within bounds
	if endPosition > sliceLength {
		endPosition = sliceLength
	}
	if offset > endPosition {
		return []T{}
	}

	return s[offset:endPosition]
}
