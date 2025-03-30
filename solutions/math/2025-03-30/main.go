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

func MinimumOperations(x, y int) int {
	res := 0
	count := 0
	num := 1
	check := false

	for (y - x) > res {
		if check {
			res += num
			num++
			check = false
		} else {
			res += num
			check = true
		}
		count++
	}

	return count
}

func main() {
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)
	result := make([]int, 0)

	for i := 0; i < t; i++ {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		result = append(result, MinimumOperations(x, y))
	}
	for i := 0; i < t; i++ {
		fmt.Fprintln(writer, result[i])
	}
}
