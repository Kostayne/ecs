package utils

// Returns the index of the first encountered element with the given value.
func IndexOf[T comparable](arr []T, value T) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}

	return -1
}
