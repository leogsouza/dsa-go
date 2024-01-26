package sumlist

/**
* Write a method that will return the sum of all the elements of the integer list,
* given list as input argument.
*
 */

// SumArray returns the sum of all elements from array
func SumArray(data []int) int {
	size := len(data)
	total := 0
	for index := 0; index < size; index++ {
		total = total + data[index]
	}
	return total
}
