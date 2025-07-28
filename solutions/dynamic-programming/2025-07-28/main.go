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

func lionCage(n int) int {
	prev := 1
	now := 3
	res := 3

	for i := 1; i < n; i++ {
		res = res*2 + prev
		prev = now
		now = res

		res = res % 9901
	}

	return res
}

func main() {
	defer writer.Flush()
	n := 0
	fmt.Fscan(reader, &n)

	fmt.Fprintln(writer, lionCage(n))
}
