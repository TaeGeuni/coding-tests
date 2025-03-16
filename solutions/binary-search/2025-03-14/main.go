package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func BinarySearch(m, max int, snack []int) int {
	left, right := 1, max
	res := 0

	for left <= right {
		mid := (left + right) / 2
		v := 0

		for i := 0; i < len(snack); i++ {
			v += snack[i] / mid
		}
		if v >= m {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return res
}

func main() {
	defer writer.Flush()

	var m, n, max, total int
	fmt.Fscan(reader, &m, &n)

	snack := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &snack[i])
		total += snack[i]
		if snack[i] > max {
			max = snack[i]
		}
	}

	if total < m {
		fmt.Fprintln(writer, 0)
	} else {
		fmt.Fprintln(writer, BinarySearch(m, max, snack))
	}
}
