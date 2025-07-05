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

func MakeNewString(n, r int, sli []string) string {
	res, underline := sli[0], ""
	for i := 0; i < (r / (n - 1)); i++ {
		underline += "_"
	}
	remain := r - (n-1)*(r/(n-1))

	for i := 1; i < n; i++ {
		res += underline

		if remain != 0 {
			if n-i == remain || sli[i][0] > '_' {
				res += "_"
				remain--
			}
		}

		res += sli[i]
	}

	return res
}

func main() {
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)
	r := m

	sli := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &sli[i])
		r -= len(sli[i])
	}
	fmt.Fprintln(writer, MakeNewString(n, r, sli))
}
