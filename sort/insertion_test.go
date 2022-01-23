package sort

import (
	"constraints"
	"reflect"
	"testing"
)

// Deprecated
func TestInsertion(t *testing.T) {
	type test[T constraints.Ordered] struct {
		name string
		args []T
		want []T
	}
	var intTests = []test[int]{
		{name: "1-9", args: []int{6, 1, 7, 2, 8, 3, 9, 4, 5}, want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{name: "0-1", args: []int{0, 1, 1, 1, 0}, want: []int{0, 0, 1, 1, 1}},
	}

	for _, tt := range intTests {
		Insertion[int](tt.args)
		if !reflect.DeepEqual(tt.args, tt.want) {
			t.Errorf("name: %s, want: %v, got: %v", tt.name, tt.want, tt.args)
		}
	}
}
