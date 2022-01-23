package sort

import (
	"constraints"
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	type test[T constraints.Ordered] struct {
		name string
		args []T
		want []T
	}

	var intTests = []test[int]{
		{name: "1-9", args: []int{6, 1, 7, 2, 8, 3, 9, 4, 5}, want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{name: "0-1", args: []int{0, 1, 1, 1, 0}, want: []int{0, 0, 1, 1, 1}},
		{name: "-10-10", args: []int{6, -4, -2, 8, -7}, want: []int{-7, -4, -2, 6, 8}},
		{name: "-100000-100000", args: []int{3273, -2155, 47452, 2342, 231, 0, 2134, -64352, 24325, -7546, -6754, -8931, 214},
			want: []int{-64352, -8931, -7546, -6754, -2155, 0, 214, 231, 2134, 2342, 3273, 24325, 47452}},
	}

	for _, tt := range intTests {
		Selection[int](tt.args)
		if !reflect.DeepEqual(tt.args, tt.want) {
			t.Errorf("algorithm:%s, name: %s, want: %v, got: %#v", "selection sort", tt.name, tt.want, tt.args)
		}
	}

	for _, tt := range intTests {
		Insertion[int](tt.args)
		if !reflect.DeepEqual(tt.args, tt.want) {
			t.Errorf("algorithm:%s, name: %s, want: %v, got: %#v", "insertion sort", tt.name, tt.want, tt.args)
		}
	}
}
