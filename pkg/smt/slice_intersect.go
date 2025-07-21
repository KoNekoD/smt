package smt

// SliceIntersect returns intersection of arrays
func SliceIntersect[T comparable](array1 []T, arrayOthers ...[]T) []T {
	c := make(map[T]bool)

	for i := 0; i < len(array1); i++ {
		if _, hasKey := c[array1[i]]; hasKey {
			c[array1[i]] = true
		} else {
			c[array1[i]] = false
		}
	}

	for i := 0; i < len(arrayOthers); i++ {
		for j := 0; j < len(arrayOthers[i]); j++ {
			if _, hasKey := c[arrayOthers[i][j]]; hasKey {
				c[arrayOthers[i][j]] = true
			} else {
				c[arrayOthers[i][j]] = false
			}
		}
	}

	result := make([]T, 0)

	for k, v := range c {
		if v {
			result = append(result, k)
		}
	}

	return result
}
