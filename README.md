# softheap
Soft Heap Implementation in Go

Description
-----------
A soft heap is a priority queue data structure that allows for the create, meld (combine
two soft heaps) and delete operations to be performed in constant amortized time while the
insert operation operates in O(log{1/e}) time, with the caveat that at most e*n
elements in the queue have their key (the priority) corrupted (i.e., their value changed
from the original).

### References
-  The paper on which this implementation is based: [A simpler implementation and analysis of Chazelle’s Soft Heaps - Haim Kaplan, Uri Zwick](https://epubs.siam.org/doi/pdf/10.1137/1.9781611973068.53).
- The original paper: [The Soft Heap: An Approximate Priority Heap with Optimal Error Rate - Bernard Chazelle](https://www.cs.princeton.edu/~chazelle/pubs/sheap.pdf).

Usage
-----

```go
// A soft heap with integer keys and string values.
// The priority of a value is the inverse of the key.
heap := softheap.New[string]()
heap.Insert(10, "Rome")
heap.Insert(2, "Paris")
heap.Insert(5, "London")
k, v := heap.ExtractMin()
fmt.Printf("%d, %s\n", k, v) // 2, Paris
k, v = heap.ExtractMin()
fmt.Printf("%d, %s\n", k, v) // 5, London
k, v = heap.ExtractMin()
fmt.Printf("%d, %s\n", k, v) // 10, Rome
```
