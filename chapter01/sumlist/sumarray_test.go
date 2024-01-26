package sumlist

import "testing"

func TestSumArray(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []int
		result int
	}{
		{
			desc:   "Sum array [1, 2, 3] should be 6",
			input:  []int{1, 2, 3},
			result: 6,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			sum := SumArray(tC.input)
			if tC.result != sum {
				t.Errorf("The sum of array should be %d, but got %d", tC.result, sum)
			}
		})
	}
}
