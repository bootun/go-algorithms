package sort

import "constraints"

func ShellSort[T constraints.Ordered](a []T) {
	N := len(a)
	h := 1
	for h < N/3 {
		h = h*3 + 1
	}

	for h > 1 {
		// TODO: implement shell's sort, time is too late
	}
}
