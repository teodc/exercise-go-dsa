package main

import (
	"fmt"
	"math/rand"
	"time"

	"exercise-go-dsa/math"
	"exercise-go-dsa/sort"
)

func main() {
	fmt.Println(">> math.fibonacci:")
	fmt.Println(math.Fibonacci(10))

	fmt.Println(">> sort.quicksort:")
	rngSrc := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(rngSrc)
	items := rng.Perm(20)
	fmt.Println("input: ", items)
	sortedItems := sort.Quicksort(items)
	fmt.Println("output: ", sortedItems)
}
