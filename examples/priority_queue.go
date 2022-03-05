package main

import (
	"fmt"
	"math/rand"

	"github.com/speckJ8/softheap"
)

func main() {
	// the priority of an element corresponds to
	// the inverse of its key
	queue := softheap.New[int]()
	for i := 0; i < 10; i++ {
		// insert elements with random priorities
		queue.Insert(rand.Intn(100), i)
	}
	queue.Print()
	for {
		k, e := queue.ExtractMin()
		if e == nil {
			break
		}
		fmt.Printf("extrated priority=1/%d, value=%d\n", k, *e)
		queue.Print()
	}
}
