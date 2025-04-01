package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func LogJumping(logs []int) int {
	res := 0

	sort.Ints(logs)

	a := logs[0]

	if len(logs)%2 == 0 {
		i := 0
		for i < len(logs) {
			if res < logs[i]-a {
				res = logs[i] - a
			}
			a = logs[i]
			i += 2
		}
		i = len(logs) - 1
		for i > 0 {
			if res < a-logs[i] {
				res = a - logs[i]
			}
			a = logs[i]
			i -= 2
		}
	} else {
		i := 0
		for i < len(logs) {
			if res < logs[i]-a {
				res = logs[i] - a
			}
			a = logs[i]
			i += 2
		}
		i = len(logs) - 2
		for i > 0 {
			if res < a-logs[i] {
				res = a - logs[i]
			}
			a = logs[i]
			i -= 2
		}
	}

	return res
}

func main() {
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)
	result := make([]int, t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(reader, &n)
		logs := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &logs[j])
		}
		result[i] = LogJumping(logs)
	}
	for i := 0; i < t; i++ {
		fmt.Fprintln(writer, result[i])
	}
}
