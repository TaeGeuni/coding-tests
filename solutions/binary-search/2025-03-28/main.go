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

func BinarySearch(biggest, total int, requestBudget []int) int {
	res := 0

	left, right := 0, biggest

	for left <= right {
		mid := (left + right) / 2
		t := total

		for i := 0; i < len(requestBudget); i++ {
			if requestBudget[i] > mid {
				t -= mid
			} else {
				t -= requestBudget[i]
			}
		}

		if t >= 0 {
			left = mid + 1
			res = mid
		} else {
			right = mid - 1
		}
	}

	return res
}

func main() {
	defer writer.Flush()

	var n, biggest, total int

	fmt.Fscan(reader, &n)
	requestBudget := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &requestBudget[i])
		if biggest < requestBudget[i] {
			biggest = requestBudget[i]
		}
	}
	fmt.Fscan(reader, &total)

	fmt.Fprintln(writer, BinarySearch(biggest, total, requestBudget))
}
