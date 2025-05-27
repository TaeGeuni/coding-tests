package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func MinimumCost(n, k int, student []int) int {
	res := 0

	gap := make([]int, 0)

	for i := 0; i < n-1; i++ {
		gap = append(gap, student[i+1]-student[i])
	}
	sort.Ints(gap)

	for i := 0; i < len(gap)+1-k; i++ {
		res += gap[i]
	}

	return res
}

func main() {
	defer writer.Flush()
	var n, k int
	fmt.Fscan(reader, &n, &k)

	student := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &student[i])
	}

	fmt.Fprintln(writer, MinimumCost(n, k, student))

}
