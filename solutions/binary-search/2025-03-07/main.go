package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n, m, left, mid, right, res int

	fmt.Fscan(reader, &n, &m)
	trees := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &trees[i])
		if trees[i] > right {
			right = trees[i]
		}
	}

	for left <= right {
		mid = (left + right) / 2
		cutTree := 0

		for i := 0; i < n; i++ {
			if trees[i] > mid {
				cutTree += trees[i] - mid
			}
		}

		if cutTree >= m {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}

	}

	fmt.Fprintln(writer, res)

}
