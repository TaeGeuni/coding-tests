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

type Point struct {
	Y int
	X int
}

func BFS(m, n int, field [][]int) int {
	res := 0

	queue := make([]Point, 0)
	visited := make(map[[2]int]bool)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if field[i][j] == 1 {
				if !visited[[2]int{i, j}] {
					queue = append(queue, Point{Y: i, X: j})
					visited[[2]int{i, j}] = true
					res++
				}
				for len(queue) > 0 {
					node := queue[0]
					if node.X == 0 {
						if field[node.Y][node.X+1] == 1 && !visited[[2]int{node.Y, node.X + 1}] {
							queue = append(queue, Point{Y: node.Y, X: node.X + 1})
							visited[[2]int{node.Y, node.X + 1}] = true
						}
					} else if node.X == m-1 {
						if field[node.Y][node.X-1] == 1 && !visited[[2]int{node.Y, node.X - 1}] {
							queue = append(queue, Point{Y: node.Y, X: node.X - 1})
							visited[[2]int{node.Y, node.X - 1}] = true
						}
					} else {
						if field[node.Y][node.X+1] == 1 && !visited[[2]int{node.Y, node.X + 1}] {
							queue = append(queue, Point{Y: node.Y, X: node.X + 1})
							visited[[2]int{node.Y, node.X + 1}] = true
						}
						if field[node.Y][node.X-1] == 1 && !visited[[2]int{node.Y, node.X - 1}] {
							queue = append(queue, Point{Y: node.Y, X: node.X - 1})
							visited[[2]int{node.Y, node.X - 1}] = true
						}
					}
					if node.Y == 0 {
						if field[node.Y+1][node.X] == 1 && !visited[[2]int{node.Y + 1, node.X}] {
							queue = append(queue, Point{Y: node.Y + 1, X: node.X})
							visited[[2]int{node.Y + 1, node.X}] = true
						}
					} else if node.Y == n-1 {
						if field[node.Y-1][node.X] == 1 && !visited[[2]int{node.Y - 1, node.X}] {
							queue = append(queue, Point{Y: node.Y - 1, X: node.X})
							visited[[2]int{node.Y - 1, node.X}] = true
						}
					} else {
						if field[node.Y+1][node.X] == 1 && !visited[[2]int{node.Y + 1, node.X}] {
							queue = append(queue, Point{Y: node.Y + 1, X: node.X})
							visited[[2]int{node.Y + 1, node.X}] = true
						}
						if field[node.Y-1][node.X] == 1 && !visited[[2]int{node.Y - 1, node.X}] {
							queue = append(queue, Point{Y: node.Y - 1, X: node.X})
							visited[[2]int{node.Y - 1, node.X}] = true
						}
					}
					queue = queue[1:]
				}
			}
		}
	}

	return res
}

func main() {
	defer writer.Flush()

	var t int
	var m, n, k int
	res := make([]int, 0)

	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		fmt.Fscan(reader, &m, &n, &k)
		field := make([][]int, n)
		for j := 0; j < n; j++ {
			field[j] = make([]int, m)
		}
		for j := 0; j < k; j++ {
			var x, y int
			fmt.Fscan(reader, &x, &y)
			field[y][x] = 1
		}
		res = append(res, BFS(m, n, field))
	}

	for i := 0; i < len(res); i++ {
		fmt.Fprintln(writer, res[i])
	}
}
