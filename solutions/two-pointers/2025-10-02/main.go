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

func twoPointer(n, s int, sequence []int) int {
	res := 100_001
	down := 0
	sum := 0
	for up := 0; up < n; up++ {
		sum += sequence[up]

		for sum >= s {
			length := up - down + 1
			if res > length {
				res = length
			}
			sum -= sequence[down]
			down++
		}
	}

	if res == 100_001 {
		return 0
	}

	return res
}

func main() {
	defer writer.Flush()
	var n, s int

	fmt.Fscan(reader, &n, &s)

	sequence := make([]int, n)

	for i := range sequence {
		fmt.Fscan(reader, &sequence[i])
	}

	fmt.Fprintln(writer, twoPointer(n, s, sequence))

}
