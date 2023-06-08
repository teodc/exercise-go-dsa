package sort

func Merge(items []int) []int {
	if len(items) < 2 {
		return items
	}

	var middle = len(items) / 2
	var left = Merge(items[:middle])
	var right = Merge(items[middle:])

	var out = make([]int, len(left)+len(right))
	var i = 0 // left index
	var j = 0 // right index
	var k = 0 // out index

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			out[k] = left[i]
			i++
		} else {
			out[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		out[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		out[k] = right[j]
		j++
		k++
	}

	return out
}
