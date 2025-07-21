package smt

// SliceShift returns a pointer to the first element of a slice and
// removes it from the slice
func SliceShift[T interface{}](s []T) (*T, []T) {
	if len(s) == 0 {
		return nil, s
	}

	t := s[0]

	s = s[1:]

	return &t, s
}
