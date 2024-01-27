package arrays

import "testing"

func TestRotateArray(t *testing.T) {
	testCases := []struct {
		desc string
		data []int
		k    int
		want []int
	}{
		{
			desc: "RotateArray [10,20,30,40,50,60] k =2 returns [30,40,50,60,10,20]",
			data: []int{10, 20, 30, 40, 50, 60},
			want: []int{30, 40, 50, 60, 10, 20},
			k:    2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			RotateArray(tC.data, tC.k)
			if !checkSlices(tC.data, tC.want) {
				t.Errorf("Expected %v but got %v instead.", tC.want, tC.data)
			}
		})
	}
}

func checkSlices(got []int, want []int) bool {
	if len(got) != len(want) {
		return false
	}

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}
