package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	dp     = make(map[int]int)
)

func Sequence(n, p, q int) int {
	if n == 0 {
		return 1
	}
	if val, exists := dp[n]; exists {
		return val
	}
	dp[n] = Sequence(n/p, p, q) + Sequence(n/q, p, q)
	return dp[n]
}

func main() {
	defer writer.Flush()

	var n, p, q int
	fmt.Fscan(reader, &n, &p, &q)

	fmt.Fprintln(writer, Sequence(n, p, q))
}
