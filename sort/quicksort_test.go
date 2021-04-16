package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	tests := []struct {
		given []int
		want  []int
	}{
		{
			[]int{1, 3, 2, 5, 6},
			[]int{1, 2, 3, 5, 6},
		},
		{
			[]int{1, 3, 2, 5, 6},
			[]int{1, 2, 3, 5, 6},
		},
		{
			[]int{1, 1022, 2, 0, -1, 5, 1},
			[]int{-1, 0, 1, 1, 2, 5, 1022},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
	}

	for idx, tt := range tests {
		testname := fmt.Sprintf("%d", idx)
		t.Run(testname, func(t *testing.T) {
			ans := Quicksort(tt.given)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
	print("Quicksort test ")
}
