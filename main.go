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

	fmt.Println(">> sort.merge:")
	rngSrc := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(rngSrc)
	items := rng.Perm(20)
	fmt.Println("input: ", items)
	// sortedItems := sort.Bubble(items)
	// sortedItems := sort.Insertion(items)
	// sortedItems := sort.Selection(items)
	sortedItems := sort.Merge(items)
	fmt.Println("output: ", sortedItems)
}
