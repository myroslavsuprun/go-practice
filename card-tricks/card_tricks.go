package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if checkIfIndexInBoundary(slice, index) {
		return slice[index]
	}

	return -1
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if checkIfIndexInBoundary(slice, index) {
		slice[index] = value
		return slice
	}

	return append(slice, value)
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	return append(values, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if checkIfIndexInBoundary(slice, index) {
		return sliceRemove(slice, index)
	}

	return slice
}

// Removes a value from the provided slice with an index.
func sliceRemove(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

// Checks whether the index is within the slice.
func checkIfIndexInBoundary(slice []int, index int) bool {
	if index < len(slice) && index >= 0 {
		return true
	}

	return false
}
