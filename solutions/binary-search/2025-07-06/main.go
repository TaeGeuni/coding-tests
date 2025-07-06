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

func isItPossible(levels []int, k, mid int) bool {
	num := 0
	for i := 0; i < len(levels); i++ {
		if levels[i] < mid {
			num += mid - levels[i]
		}
		if num > k {
			return false
		}
	}

	return true
}

func binarySearchMaxTarget(levels []int, k int) int {
	res := 0
	lo, hi := levels[0], levels[len(levels)-1]+k

	for lo <= hi {
		mid := (lo + hi) / 2
		if isItPossible(levels, k, mid) {
			res = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return res
}

func main() {
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	levels := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &levels[i])
	}

	sort.Ints(levels)

	result := binarySearchMaxTarget(levels, k)
	fmt.Fprintln(writer, result)
}
