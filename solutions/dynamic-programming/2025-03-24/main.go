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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	defer writer.Flush()

	var num int
	fmt.Fscan(reader, &num)

	stairs := make([]int, num+1)

	for i := 1; i < num+1; i++ {
		fmt.Fscan(reader, &stairs[i])
	}

	dp := make([]int, num+1)

	if num >= 1 {
		dp[1] = stairs[1]
	}
	if num >= 2 {
		dp[2] = stairs[1] + stairs[2]
	}

	for i := 3; i < num+1; i++ {
		dp[i] = max((dp[i-2] + stairs[i]), (dp[i-3] + stairs[i-1] + stairs[i]))
	}

	fmt.Fprintln(writer, dp[num])
}
