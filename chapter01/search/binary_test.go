package search

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		desc  string
		data  []int
		value int
		res   bool
	}{
		{
			desc:  "Binary Search: [1, 27, 35, 42, 45, 74, 95, 98, 105], value =95 returns true",
			data:  []int{1, 27, 35, 42, 45, 74, 95, 98, 105},
			value: 95,
			res:   true,
		},
		{
			desc:  "Binary Search: [5, 8, 21, 47, 58], value 16 returns false",
			data:  []int{5, 5, 21, 47, 58},
			value: 16,
			res:   false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if got := BinarySearch(tC.data, tC.value); got != tC.res {
				t.Errorf("Expected %t, but got %t.", tC.res, got)
			}
		})
	}
}
