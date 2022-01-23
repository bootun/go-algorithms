package sort

import (
	"constraints"
)

func Insertion[T constraints.Ordered](a []T) {
	N := len(a)
	for i := 1; i < N; i++ {
		//for j := i; a[j] < a[j-1] && j > 0) j-- { // panic: runtime error: index out of range [-1]
		for j := i; j > 0 && a[j] < a[j-1]; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}
