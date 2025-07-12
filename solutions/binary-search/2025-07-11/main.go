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

func distributePizza(k int, heights, widths []int) int {
	res := 0

	for _, w := range widths {
		if w == 0 {
			continue
		}
		target := k / w

		num := sort.Search(len(heights), func(j int) bool {
			return heights[j] > target
		})
		res += num
	}

	return res
}

func main() {
	defer writer.Flush()

	var w, h, k int
	fmt.Fscan(reader, &w, &h, &k)

	var n, m int

	fmt.Fscan(reader, &n)
	ys := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &ys[i])
	}
	heights := make([]int, 0, n+1)
	prev := 0
	for i := 0; i < n; i++ {
		heights = append(heights, ys[i]-prev)
		prev = ys[i]
	}
	heights = append(heights, h-prev)

	sort.Ints(heights)

	fmt.Fscan(reader, &m)
	xs := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &xs[i])
	}
	widths := make([]int, 0, m+1)
	prev = 0
	for i := 0; i < m; i++ {
		widths = append(widths, xs[i]-prev)
		prev = xs[i]
	}
	widths = append(widths, w-prev)

	fmt.Fprintln(writer, distributePizza(k, heights, widths))
}
