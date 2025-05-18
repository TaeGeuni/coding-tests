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

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func DP(n int, p []int) int {
	dp := make([]int, 0)
	dp = append(dp, p[0])

	for i := 1; i < n; i++ {
		left := 0
		right := i - 1
		num := 0
		for left <= right {
			if num < (dp[left] + dp[right]) {
				num = (dp[left] + dp[right])
			}
			left++
			right--
		}
		dp = append(dp, maximum(num, p[i]))
	}

	return dp[n-1]
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	p := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &p[i])
	}

	fmt.Fprintln(writer, DP(n, p))
}
