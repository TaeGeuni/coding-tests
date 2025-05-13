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

func Greedy(volunteer [][]int) int {
	pass := len(volunteer)

	sort.Slice(volunteer, func(i, j int) bool {
		return volunteer[i][0] < volunteer[j][0]
	})

	num := volunteer[0][1]

	for i := 1; i < len(volunteer); i++ {
		if num < volunteer[i][1] {
			pass--
		} else {
			num = volunteer[i][1]
		}
	}

	return pass
}

func main() {
	defer writer.Flush()

	var t, n int

	fmt.Fscan(reader, &t)

	res := make([]int, t)

	for i := 0; i < t; i++ {
		fmt.Fscan(reader, &n)
		volunteer := make([][]int, n)

		for j := 0; j < n; j++ {
			var document, interview int
			fmt.Fscan(reader, &document, &interview)
			volunteer[j] = append(volunteer[j], document)
			volunteer[j] = append(volunteer[j], interview)
		}

		res[i] = Greedy(volunteer)

	}

	for i := 0; i < t; i++ {
		fmt.Fprintln(writer, res[i])
	}
}
