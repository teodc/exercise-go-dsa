package sort

func Bubble(items []int) []int {
	switched := true
	for switched {
		switched = false
		for i := 0; i < len(items)-1; i++ {
			if items[i] > items[i+1] {
				items[i], items[i+1] = items[i+1], items[i]
				switched = true
			}
		}
	}

	return items
}
