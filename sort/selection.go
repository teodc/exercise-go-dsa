package sort

func Selection(items []int) []int {
	for i := 0; i < len(items); i++ {
		min := i
		for j := i + 1; j < len(items); j++ {
			if items[j] < items[min] {
				min = j
			}
		}
		items[i], items[min] = items[min], items[i]
	}

	return items
}
