package sort

import (
	"math/rand"
)

func Quicksort(items []int) []int {
	if len(items) < 2 {
		return items
	}

	head, tail := 0, len(items)-1
	pivot := rand.Intn(len(items))

	items[pivot], items[tail] = items[tail], items[pivot]

	for i := range items {
		if items[i] < items[tail] {
			items[head], items[i] = items[i], items[head]
			head++
		}
	}

	items[head], items[tail] = items[tail], items[head]

	Quicksort(items[:head])
	Quicksort(items[head+1:])

	return items
}
