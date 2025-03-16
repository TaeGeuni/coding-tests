package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func Ookami(input string) bool {
	var w, o, l, f int

	for i := 0; i < len(input); i++ {
		switch {
		case input[i] == 'w':
			if o == 0 && l == 0 && f == 0 {
				w++
			} else {
				return false
			}
		case input[i] == 'o':
			if w > o && l == 0 && f == 0 {
				o++
			} else {
				return false
			}
		case input[i] == 'l':
			if w == o && l < w && f == 0 {
				l++
			} else {
				return false
			}
		case input[i] == 'f':
			if w == o && w == l && f < w {
				f++
			} else {
				return false
			}
			if w == o && w == l && w == f {
				w, o, l, f = 0, 0, 0, 0
			}
		}
	}

	return w == 0 && o == 0 && l == 0 && f == 0
}

func main() {
	defer writer.Flush()

	var input string
	fmt.Fscan(reader, &input)

	if Ookami(input) {
		fmt.Fprintln(writer, 1)
	} else {
		fmt.Fprintln(writer, 0)
	}
}
