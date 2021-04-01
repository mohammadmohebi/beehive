package sort

func Bubble(data []int) []int {
	d := data
	c := 0
	done := false
	for !done && c < len(d) {
		done = true
		for i := 0; i < len(d); i++ {
			x, y := i, i+1
			if y < len(d) && d[x] > d[y] {
				d[x], d[y] = d[y], d[x]
				done = false
			}
		}
		c += 1
	}
	return d
}
