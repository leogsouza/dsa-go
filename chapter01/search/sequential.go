package search

// Write a method, which will search a list for some given value

// SequentialSearch looks for a value in the data list returning true if find a
// value otherwise returns false
func SequentialSearch(data []int, value int) bool {
	size := len(data)
	for i := 0; i < size; i++ {
		if value == data[i] {
			return true
		}
	}
	return false
}
