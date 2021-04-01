package sort

func Selection(data []int) []int {
	d := data
	for i := 0; i < len(d); i++ {
		x, y := i, i
		for j := i + 1; j < len(d); j++ {
			if d[j] < d[y] {
				y = j
			}
		}

		if x != y {
			d[x], d[y] = d[y], d[x]
		}
	}
	return d
}
