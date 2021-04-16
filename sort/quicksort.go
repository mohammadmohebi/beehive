package sort

func partition(A []int, lo int, hi int) int {
	i := lo
	if hi < len(A) {
		pivot := A[hi]
		for j := lo; j < hi; j++ {
			if A[j] < pivot {
				if i != j {
					A[i], A[j] = A[j], A[i]
				}
				i++
			}
		}
		A[i], A[hi] = A[hi], A[i]
	}
	return i
}

func quicksort(A []int, lo int, hi int) {
	if lo < hi {
		p := partition(A, lo, hi)
		quicksort(A, lo, p-1)
		quicksort(A, p+1, hi)
	}
}

func Quicksort(data []int) []int {
	d := make([]int, len(data))
	copy(d, data)
	if len(d) > 1 {
		lo := 0
		hi := len(d) - 1
		quicksort(d, lo, hi)
	}
	return d
}
