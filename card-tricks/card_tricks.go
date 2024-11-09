package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

func isOutOfBand(slice []int, index int) bool {
	return index >= len(slice) || index < 0
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if isOutOfBand(slice, index) {
		return -1
	}
	return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if isOutOfBand(slice, index) {
		return append(slice, value)
	}
	return append(slice[:index], append([]int{value}, slice[index+1:]...)...)
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	return append(values, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if isOutOfBand(slice, index) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}
