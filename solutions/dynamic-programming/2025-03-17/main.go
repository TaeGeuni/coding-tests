package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func Minimum(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	defer writer.Flush()

	var num int
	fmt.Fscan(reader, &num)
	cost := make([][3]int, num)

	for i := 0; i < num; i++ {
		var r, g, b int
		fmt.Fscan(reader, &r, &g, &b)

		cost[i][0], cost[i][1], cost[i][2] = r, g, b

	}
	dp := make([][3]int, num)

	dp[0][0], dp[0][1], dp[0][2] = cost[0][0], cost[0][1], cost[0][2]

	for i := 1; i < num; i++ {
		dp[i][0] = Minimum(dp[i-1][1], dp[i-1][2]) + cost[i][0]
		dp[i][1] = Minimum(dp[i-1][0], dp[i-1][2]) + cost[i][1]
		dp[i][2] = Minimum(dp[i-1][0], dp[i-1][1]) + cost[i][2]
	}
	res := Minimum(dp[num-1][0], dp[num-1][1])
	res = Minimum(res, dp[num-1][2])

	fmt.Fprintln(writer, res)

}
