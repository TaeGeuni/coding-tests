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

func TwoPointers(n int, solution []int) (int, int) {
	res := 2_000_000_000
	var a, b int

	ls, rs := 0, n-1
	left, right := solution[ls], solution[rs]

	for i := 0; i < n-1; i++ {
		mid := left + right

		if mid < 0 {
			ls++
			mid = mid * (-1)
		} else if mid > 0 {
			rs--
		} else {
			a = left
			b = right
			break
		}
		if res >= mid {
			res = mid
			a = left
			b = right
		}

		left, right = solution[ls], solution[rs]
	}

	return a, b
}

func main() {
	defer writer.Flush()

	var n int

	fmt.Fscan(reader, &n)

	solution := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &solution[i])
	}

	a, b := TwoPointers(n, solution)

	fmt.Fprintln(writer, a, b)

}
