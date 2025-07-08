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

func isThisStable(s string) int {
	res := 0
	count := 0

	for i := 0; i < len(s); i++ {
		if count == 0 && s[i] == '{' {
			count++
		} else if count == 0 && s[i] == '}' {
			res++
			count++
		} else if len(s)-i == count && s[i] == '{' {
			res++
			count--
		} else if len(s)-i == count && s[i] == '}' {
			count--
		} else if s[i] == '{' {
			count++
		} else if s[i] == '}' {
			count--
		} else if s[i] == '-' {
			return -1
		}
	}

	return res
}

func main() {
	defer writer.Flush()

	s := ""
	count := 0

	for {
		fmt.Fscan(reader, &s)

		n := isThisStable(s)

		if n < 0 {
			break
		}
		count++
		fmt.Fprintf(writer, "%d. %d\n", count, n)
	}

}
