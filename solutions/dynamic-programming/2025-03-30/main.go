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

func Maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func DP(sequence []int) int {
	max := 0
	sum := 0

	max = sequence[0]
	if sequence[0] > 0 {
		sum = sequence[0]
	} else {
		sum = 0
	}

	for i := 1; i < len(sequence); i++ {
		max = Maximum(max, (sum + sequence[i]))
		sum += sequence[i]
		if sum < 0 {
			sum = 0
		}
	}

	return max
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	sequence := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &sequence[i])
	}
	fmt.Fprintln(writer, DP(sequence))
}
