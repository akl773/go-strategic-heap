# Go-Strategic-Heap

Go-Strategic-Heap is a powerful and efficient Go library that implements heap data structures. It provides support for both min-heap and max-heap strategies, making it an ideal choice for projects that require priority queues or heap sorts. With its simple API and high-performance operations, Strategic-Heap-Go makes it easy to create and manipulate heaps.

## Features

- Easy-to-use API
- Support for both min-heap and max-heap strategies
- High-performance operations

## Installation

To install Strategic-Heap-Go, use the `go get` command:

```shell
go get github.com/akl773/Strategic-Heap-Go/heap
```

Usage
Here's a simple example that demonstrates how to use the library:

```shell
package main

import (
	"fmt"
	"github.com/akl773/Strategic-Heap-Go/heap"
)

func main() {
	h := heap.NewMinHeap()

	for i := 5; i > 0; i-- {
		h.Push(i)
	}

	for h.Len() != 0 {
		fmt.Println(h.Pop())
	}
}
```
This program creates a min heap, pushes some integers onto it, and then pops them off. When executed, it will print the integers 1 through 5 in ascending order.

## Testing

You can run the tests using the go test command:

```bash
go test
```

## Contributing
Pull requests are welcome. If you plan to make major changes, please open an issue first to discuss what you would like to change.

Please ensure that you update the tests as appropriate.
