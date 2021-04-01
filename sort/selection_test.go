package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSelection(t *testing.T) {
	tests := []struct {
		given []int
		want  []int
	}{
		{
			[]int{12, 3, 2, 5, 6},
			[]int{2, 3, 5, 6, 12},
		},
		{
			[]int{1, 30, 2, 5, 6},
			[]int{1, 2, 5, 6, 30},
		},
		{
			[]int{1, 1022, 2, 0, -1, 5, 1},
			[]int{-1, 0, 1, 1, 2, 5, 1022},
		},
	}

	for idx, tt := range tests {
		testname := fmt.Sprintf("%d", idx)
		t.Run(testname, func(t *testing.T) {
			ans := Selection(tt.given)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
