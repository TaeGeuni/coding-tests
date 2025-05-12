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

func maximum(a, b, c int) int {
	var d int
	if a > b {
		d = a
	} else {
		d = b
	}
	if d > c {
		return d
	}
	return c
}

func DP(n int, wine []int) int {
	if n == 1 {
		return wine[0]
	} else if n == 2 {
		return wine[0] + wine[1]
	}
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 0, wine[0], wine[0]+wine[1]

	for i := 2; i < n; i++ {
		dp[i+1] = maximum(dp[i-2]+wine[i-1]+wine[i], dp[i-1]+wine[i], dp[i])
	}

	return dp[n]
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	wine := make([]int, n+1)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &wine[i])
	}

	res := DP(n, wine)

	fmt.Fprintln(writer, res)
}
