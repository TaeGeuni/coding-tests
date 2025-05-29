package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func Benchmark(n int, building []int) int {
	res := 0

	stack := make([]int, 0)
	stack = append(stack, building[0])

	for i := 1; i < n; i++ {
		if stack[len(stack)-1] > building[i] {
			res += len(stack)
			stack = append(stack, building[i])
		} else {
			for len(stack) > 0 {
				stack = stack[:len(stack)-1]
				if len(stack) > 0 && stack[len(stack)-1] > building[i] {
					break
				}
			}
			res += len(stack)
			stack = append(stack, building[i])
		}
	}

	return res
}

func main() {
	defer writer.Flush()

	var n int

	fmt.Fscan(reader, &n)

	building := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &building[i])
	}

	fmt.Fprintln(writer, Benchmark(n, building))
}
