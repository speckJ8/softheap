package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/speckJ8/softheap"
)

func priority_queue() {
	// the priority of an element corresponds to
	// the inverse of its key
	queue := softheap.New[int]()
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		// insert elements with random priorities
		key := rand.Intn(100)
		fmt.Printf("inserted: priority=1/%d, value=%d\n", key, i)
		queue.Insert(key, i)
	}
	for {
		// extract the current element in the queue with
		// the highest priority (i.e., the lowest key value )
		k, e := queue.ExtractMin()
		if e == nil {
			break
		}
		fmt.Printf("extrated: priority=1/%d, value=%d\n", k, *e)
	}
}

func approximate_sorting() {
	// the array to be sorted
	array := make([]int, 100)
	rand.Seed(time.Now().Unix())
	for i := range array {
		array[i] = rand.Intn(100)
	}
	fmt.Printf("original array {%v}: %v\n", softheap.ErrorParameter, array)
	heap := softheap.New[int]()
	for i := range array {
		heap.Insert(array[i], array[i])
	}
	for i := range array {
		// The array elements will be extracted from the heap
		// in order of lowest key to highest key.
		// So the i-th element extracted is the i-th smallest.
		// Note that at most epsilon*n elements in the heap have
		// their values corrupted so at most epsilon*n elements
		// in the final sorted array will be out of place (in unsorted
		// positions)
		e, _ := heap.ExtractMin()
		array[i] = e
	}
	fmt.Printf("(approx) sorted array: %v\n", array)
}

func selection() {
}

func main() {
	approximate_sorting()
}
