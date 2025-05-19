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

func LIS(n int, pole [][]int) int {

	tails := make([]int, 0)
	tails = append(tails, 0)

	for i := 0; i < n; i++ {
		if tails[len(tails)-1] < pole[i][1] {
			tails = append(tails, pole[i][1])
		} else {
			left, right := 0, len(tails)-1
			res := 0
			for left <= right {
				mid := (left + right) / 2
				num := tails[mid]
				if num >= pole[i][1] {
					res = mid
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
			tails[res] = pole[i][1]
		}
	}

	return (n - (len(tails) - 1))
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	pole := make([][]int, n)

	for i := 0; i < n; i++ {
		a, b := 0, 0
		fmt.Fscan(reader, &a, &b)
		pole[i] = append(pole[i], a)
		pole[i] = append(pole[i], b)
	}

	sort.Slice(pole, func(i, j int) bool {
		return pole[i][0] < pole[j][0]
	})

	fmt.Fprintln(writer, LIS(n, pole))

}
