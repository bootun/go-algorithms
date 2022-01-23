package sort

import (
	"constraints"
)

func Selection[T constraints.Ordered](a []T) {
	length := len(a)
	for i := 0; i < length; i++ {
		min := i
		for j := i; j < length; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}
