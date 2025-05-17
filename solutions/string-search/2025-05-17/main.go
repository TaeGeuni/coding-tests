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

func IOIOI(n, m int, text string) int {
	res := 0

	count := 0
	ok := false
	next := false

	ioi := (2 * n) + 1

	for i := 0; i < m; i++ {

		if ok && next {
			if text[i] == 'O' {
				next = false
			} else {
				ok = false
				next = false
				if count == ioi {
					res++
				} else if count > ioi {
					count -= ioi
					res++
					res += count / 2
				}
				count = 0
			}
		} else if ok && !next {
			if text[i] == 'I' {
				next = true
				count += 2
			} else {
				ok = false
				next = false
				if count == ioi {
					res++
				} else if count > ioi {
					count -= ioi
					res++
					res += count / 2
				}
				count = 0
			}
		}
		if text[i] == 'I' && !ok {
			ok = true
			next = true
			count++
		}

		if i == m-1 {
			if count == ioi {
				res++
			} else if count > ioi {
				count -= ioi
				res++
				res += count / 2
			}
		}
	}

	return res
}

func main() {
	defer writer.Flush()

	var n, m int

	fmt.Fscan(reader, &n, &m)
	text := ""
	fmt.Fscan(reader, &text)

	res := IOIOI(n, m, text)

	fmt.Fprintln(writer, res)

}
