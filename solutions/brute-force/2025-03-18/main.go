package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func CheapestCase(n int, garden [][]int) int {
	res := 0
	var x1, y1, x2, y2, x3, y3 int

	for i := 0; i < n*n; i++ {
		check1 := make(map[[2]int]struct{})
		cost1 := 0
		if x1 == n {
			x1 = 0
			y1++
		}
		if x1 == 0 || y1 == 0 || x1 == n-1 || y1 == n-1 {
			x1++
			continue
		}
		cost1 += (garden[y1][x1] + garden[y1-1][x1] + garden[y1][x1-1] + garden[y1][x1+1] + garden[y1+1][x1])
		check1[[2]int{y1, x1}] = struct{}{}
		check1[[2]int{y1 - 1, x1}] = struct{}{}
		check1[[2]int{y1, x1 - 1}] = struct{}{}
		check1[[2]int{y1, x1 + 1}] = struct{}{}
		check1[[2]int{y1 + 1, x1}] = struct{}{}
		x2, y2 = x1, y1
		for j := i; j < n*n; j++ {
			check2 := make(map[[2]int]struct{})
			cost2 := 0

			if x2 == n {
				x2 = 0
				y2++
			}
			if x2 == 0 || y2 == 0 || x2 == n-1 || y2 == n-1 {
				x2++
				continue
			}
			if _, ok := check1[[2]int{y2, x2}]; ok {
				x2++
				continue
			}
			if _, ok := check1[[2]int{y2 - 1, x2}]; ok {
				x2++
				continue
			}
			if _, ok := check1[[2]int{y2, x2 - 1}]; ok {
				x2++
				continue
			}
			if _, ok := check1[[2]int{y2, x2 + 1}]; ok {
				x2++
				continue
			}
			if _, ok := check1[[2]int{y2 + 1, x2}]; ok {
				x2++
				continue
			}
			cost2 += (garden[y2][x2] + garden[y2-1][x2] + garden[y2][x2-1] + garden[y2][x2+1] + garden[y2+1][x2])
			check2[[2]int{y2, x2}] = struct{}{}
			check2[[2]int{y2 - 1, x2}] = struct{}{}
			check2[[2]int{y2, x2 - 1}] = struct{}{}
			check2[[2]int{y2, x2 + 1}] = struct{}{}
			check2[[2]int{y2 + 1, x2}] = struct{}{}
			x3, y3 = x2, y2

			for k := j; k < n*n; k++ {
				cost3 := 0
				if x3 == n {
					x3 = 0
					y3++
				}
				if x3 == 0 || y3 == 0 || x3 == n-1 || y3 == n-1 {
					x3++
					continue
				}
				if _, ok := check1[[2]int{y3, x3}]; ok {
					x3++
					continue
				}
				if _, ok := check1[[2]int{y3 - 1, x3}]; ok {
					x3++
					continue
				}
				if _, ok := check1[[2]int{y3, x3 - 1}]; ok {
					x3++
					continue
				}
				if _, ok := check1[[2]int{y3, x3 + 1}]; ok {
					x3++
					continue
				}
				if _, ok := check1[[2]int{y3 + 1, x3}]; ok {
					x3++
					continue
				}
				if _, ok := check2[[2]int{y3, x3}]; ok {
					x3++
					continue
				}
				if _, ok := check2[[2]int{y3 - 1, x3}]; ok {
					x3++
					continue
				}
				if _, ok := check2[[2]int{y3, x3 - 1}]; ok {
					x3++
					continue
				}
				if _, ok := check2[[2]int{y3, x3 + 1}]; ok {
					x3++
					continue
				}
				if _, ok := check2[[2]int{y3 + 1, x3}]; ok {
					x3++
					continue
				}
				cost3 += (garden[y3][x3] + garden[y3-1][x3] + garden[y3][x3-1] + garden[y3][x3+1] + garden[y3+1][x3])
				if res == 0 {
					res = cost1 + cost2 + cost3
				} else {
					if res > (cost1 + cost2 + cost3) {
						res = cost1 + cost2 + cost3
					}
				}
				x3++
			}
			x2++
		}
		x1++
	}

	return res
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	garden := make([][]int, n)
	for i := 0; i < n; i++ {
		garden[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &garden[i][j])
		}
	}

	fmt.Fprintln(writer, CheapestCase(n, garden))

}
