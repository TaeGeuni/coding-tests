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

func BinarySearch(n, m, total int, lesson []int) int {
	res := 0

	left, right := 1, total

	for left <= right {
		done := 0
		max := 0
		num := 0
		queue := lesson
		mid := (left + right) / 2

		for i := 0; i < n; i++ {
			if queue[0] > mid {
				break
			}

			if num+queue[0] <= mid {
				num += queue[0]
				queue = queue[1:]
				if max < num {
					max = num
				}
			} else {
				if max < num {
					max = num
				}
				num = 0
				num += queue[0]
				queue = queue[1:]
				done++
			}

			if i == n-1 {
				done++
			}
		}

		if len(queue) != 0 {
			left = mid + 1
		} else {
			if done > m {
				left = mid + 1
			} else {
				right = mid - 1
				res = max
			}
		}

	}

	return res
}

func main() {
	defer writer.Flush()

	var n, m, total int

	fmt.Fscan(reader, &n, &m)

	lesson := make([]int, n)

	for i := 0; i < n; i++ {
		min := 0
		fmt.Fscan(reader, &min)

		total += min
		lesson[i] = min
	}

	result := BinarySearch(n, m, total, lesson)
	fmt.Fprintln(writer, result)

}
