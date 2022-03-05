package main

import (
	"fmt"
	"math/rand"

	"github.com/speckJ8/softheap"
)

func main() {
	// the priority of an element corresponds to
	// the inverse of its key
	heap := softheap.New[int]()
	for i := 0; i < 10; i++ {
		// insert elements with random priorities
		heap.Insert(rand.Intn(100), i)
	}
	heap.Print()
	for {
		k, e := heap.ExtractMin()
		if e == nil {
			break
		}
		fmt.Printf("extrated priority=1/%d, value=%d\n", k, *e)
		heap.Print()
	}
}
