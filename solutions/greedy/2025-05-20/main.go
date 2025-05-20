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

func AB(s, t string) int {
	for len(s) != len(t) {
		if t[len(t)-1] == 'A' {
			t = t[:len(t)-1]
		} else {
			t = t[:len(t)-1]
			sli := make([]byte, len(t))
			for i := 0; i < len(t); i++ {
				sli[i] = t[len(t)-1-i]
			}
			t = string(sli)
		}
	}

	if s == t {
		return 1
	}

	return 0
}

func main() {
	defer writer.Flush()

	var s, t string
	fmt.Fscan(reader, &s)
	fmt.Fscan(reader, &t)

	fmt.Fprintln(writer, AB(s, t))
}
