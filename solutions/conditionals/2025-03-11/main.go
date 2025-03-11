package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func BigAndSmallNumber(digits []byte) bool {
	l := len(digits)

	if l == 1 {
		return false
	}

	m := l - 2

	for (m >= 0) && (digits[m] >= digits[m+1]) {
		m--
	}
	if m < 0 {
		return false
	}

	n := l - 1

	for digits[m] >= digits[n] {
		n--
	}

	digits[m], digits[n] = digits[n], digits[m]

	sort.Slice(digits[m+1:], func(i, j int) bool {
		return digits[m+1+i] < digits[m+1+j]
	})

	return true

}

func main() {
	defer writer.Flush()
	var x string
	fmt.Fscan(reader, &x)

	digits := []byte(x)

	if BigAndSmallNumber(digits) {
		fmt.Fprintln(writer, string(digits))
	} else {
		fmt.Fprintln(writer, 0)
	}

}
