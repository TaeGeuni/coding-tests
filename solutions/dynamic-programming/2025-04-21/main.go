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

func DP(d, k int) (int, int) {
	sli := make([]int, d)
	sli[d-1] = k

	for i := k - 1; i >= 1; i-- {
		sli[d-2] = i
		ok := true
		for j := 0; j < d-2; j++ {
			sli[d-3-j] = sli[d-1-j] - sli[d-2-j]
			if sli[d-3-j] <= 0 || sli[d-3-j] > sli[d-2-j] {
				ok = false
				break
			}
		}
		if ok {
			break
		}
	}

	f, t := sli[0], sli[1]

	return f, t
}

func main() {
	defer writer.Flush()

	var d, k int

	fmt.Fscan(reader, &d, &k)

	f, t := DP(d, k)

	fmt.Fprintln(writer, f)
	fmt.Fprintln(writer, t)

}
