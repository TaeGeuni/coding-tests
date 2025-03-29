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

func Overtake(input, output []string) int {
	res := 0

	check := make(map[string]bool)

	for i := 0; i < len(output); i++ {
		if check[input[0]] {
			input = input[1:]
			i--
			continue
		}
		if input[0] == output[i] {
			input = input[1:]
		} else {
			check[output[i]] = true
			res++
		}
	}

	return res
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	input := make([]string, n)
	output := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &input[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &output[i])
	}

	fmt.Fprintln(writer, Overtake(input, output))
}
