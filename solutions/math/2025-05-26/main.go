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

func catalanBinomial(n int) int {
	res := 1

	for i := 0; i < n; i++ {
		res = (res * (2*n - i)) / (i + 1)
	}

	return res / (n + 1)
}

func main() {
	defer writer.Flush()

	var n int
	for {
		fmt.Fscan(reader, &n)
		if n == 0 {
			break
		}
		fmt.Fprintln(writer, catalanBinomial(n))
	}
}
