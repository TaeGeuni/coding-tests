package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func PadovanSequence(n int) int {
	res := 0

	if n == 1 || n == 2 || n == 3 {
		return 1
	}

	a, b, c := 1, 1, 1

	for i := 3; i < n; i++ {
		res = a + b
		a = b
		b = c
		c = res
	}

	return res
}

func main() {
	defer writer.Flush()

	var t, n int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		fmt.Fscan(reader, &n)

		fmt.Fprintln(writer, PadovanSequence(n))
	}
}
