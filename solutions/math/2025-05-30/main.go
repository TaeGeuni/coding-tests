package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func Diet(g int) []int {
	res := make([]int, 0)

	for i := 1; i < g; i++ {
		num := g + (i * i)

		x := math.Sqrt(float64(num))

		if x == float64(int(x)) {
			res = append(res, int(x))
		}

	}

	if len(res) == 0 {
		res = append(res, -1)
	}

	return res
}

func main() {
	defer writer.Flush()

	var g int
	fmt.Fscan(reader, &g)

	weight := Diet(g)

	for i := 0; i < len(weight); i++ {
		fmt.Fprintln(writer, weight[i])
	}
}
