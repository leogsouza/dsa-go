package search

import "testing"

func TestSequentialSearch(t *testing.T) {
	testCases := []struct {
		desc   string
		data   []int
		value  int
		output bool
	}{
		{
			desc:   "Find value in an array of 5 elements",
			data:   []int{5, 7, 8, 13, 2},
			value:  2,
			output: true,
		},
		{
			desc:  "Dont find a value in an array of 7 elements",
			data:  []int{13, 27, 5, 9, 10, 8, 4},
			value: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := SequentialSearch(tC.data, tC.value)
			if tC.output != got {
				t.Errorf("Expected %t but got %t instead.", tC.output, got)
			}
		})
	}
}
