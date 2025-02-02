package utils

// Fast removes the element at the given index, changes the order of the array
func FastRemoveI[T any](arr []T, i int) []T {
	if len(arr) == 0 {
		return []T{}
	}

	if i < 0 || i >= len(arr) {
		panic("Index out of range!")
	}

	arr[i] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}

// Fast removes the first encountered element with the given value, changes the order of the array
func FastRemove[T comparable](arr []T, item T) []T {
	index := IndexOf(arr, item)

	if index == -1 {
		return arr
	}

	return FastRemoveI(arr, IndexOf(arr, item))
}

// Slow removes the element at the given index, saves the order of the array
func ShiftRemoveI[T any](arr []T, index int) []T {
	if len(arr) == 0 {
		return []T{}
	}

	if index < 0 || index >= len(arr) {
		panic("Index out of range!")
	}

	return append(arr[:index], arr[index+1:]...)
}
