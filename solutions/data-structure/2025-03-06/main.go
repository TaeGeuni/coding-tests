package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func CutPipe(input string) int {
	res := 0

	var pipe int
	var prev byte

	for i := 0; i < len(input); i++ {
		if input[i] == '(' {
			pipe++
			prev = '('
		} else {
			if prev == '(' {
				pipe--
				res += pipe
			} else {
				pipe--
				res++
			}
			prev = ')'
		}
	}

	return res
}

func main() {
	defer writer.Flush()

	input := ""
	fmt.Fscan(reader, &input)

	fmt.Fprintln(writer, CutPipe(input))
}
