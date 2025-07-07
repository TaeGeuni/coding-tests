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

func sumOfSections(x1, y1, x2, y2 int, table [][]int) int {
	res := table[x2][y2] - table[x2][y1-1] - table[x1-1][y2] + table[x1-1][y1-1]

	return res
}

func main() {
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	table := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		table[i] = make([]int, n+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			fmt.Fscan(reader, &table[i][j])
			table[i][j] += table[i-1][j] + table[i][j-1] - table[i-1][j-1]
		}
	}

	res := make([]int, m)

	for i := 0; i < m; i++ {
		x1, y1, x2, y2 := 0, 0, 0, 0
		fmt.Fscan(reader, &x1, &y1, &x2, &y2)
		res[i] = sumOfSections(x1, y1, x2, y2, table)
	}

	for i := 0; i < m; i++ {
		fmt.Fprintln(writer, res[i])
	}
}
