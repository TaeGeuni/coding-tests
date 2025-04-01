package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	defer writer.Flush()

	var n int
	h := new(IntHeap)
	heap.Init(h)
	fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		input := 0
		fmt.Fscan(reader, &input)
		if input == 0 {
			if h.Len() == 0 {
				fmt.Fprintln(writer, 0)
			} else {
				fmt.Fprintln(writer, heap.Pop(h))
			}
		} else {
			heap.Push(h, input)
		}
	}
}
