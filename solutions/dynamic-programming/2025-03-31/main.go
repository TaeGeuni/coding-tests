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

func Maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func DP(triangle [][]int) int {
	res := 0

	dpSli := make([]int, 0)
	dpSli = append(dpSli, triangle[0][0])

	for i := 1; i < len(triangle); i++ {
		data := make([]int, 0)
		for j := 0; j < i+1; j++ {
			if j == 0 {
				data = append(data, (dpSli[0] + triangle[i][0]))
			} else if j == i {
				data = append(data, (dpSli[len(dpSli)-1] + triangle[i][j]))
			} else {
				data = append(data, Maximum((dpSli[j-1]+triangle[i][j]), (dpSli[j]+triangle[i][j])))
			}
		}
		dpSli = data
	}

	for i := 0; i < len(dpSli); i++ {
		if res < dpSli[i] {
			res = dpSli[i]
		}
	}

	return res
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	triangle := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			num := 0
			fmt.Fscan(reader, &num)
			triangle[i] = append(triangle[i], num)
		}
	}

	fmt.Fprintln(writer, DP(triangle))
}
