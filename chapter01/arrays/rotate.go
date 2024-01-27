package arrays

func RotateArray(data []int, k int) {
	n := len(data)
	// reverse array from 0 position until k-1 position
	reverseArray(data, 0, k-1)
	// reverse array from k position until the end of array
	reverseArray(data, k, n-1)
	// rever array from the beginning until the end
	reverseArray(data, 0, n-1)
}

func reverseArray(data []int, start int, end int) {
	i := start
	j := end

	for i < j {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}
