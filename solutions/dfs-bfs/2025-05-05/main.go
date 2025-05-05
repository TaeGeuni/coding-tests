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
	x int
	y int
}

func BFS(n int, coordinate [][]int) int {
	res := 0

	for i := 0; i < 101; i++ {
		safe := 0
		queue := make([]Point, 0)
		visited := make(map[[2]int]bool)

		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if i >= coordinate[j][k] {
					continue
				} else {
					if !visited[[2]int{j, k}] {
						safe++
						queue = append(queue, Point{x: j, y: k})
						visited[[2]int{j, k}] = true
					}
					for len(queue) != 0 {
						node := queue[0]

						if node.x == 0 {
							if coordinate[node.x+1][node.y] > i && !visited[[2]int{node.x + 1, node.y}] {
								visited[[2]int{node.x + 1, node.y}] = true
								queue = append(queue, Point{x: node.x + 1, y: node.y})
							}
						} else if node.x == n-1 {
							if coordinate[node.x-1][node.y] > i && !visited[[2]int{node.x - 1, node.y}] {
								visited[[2]int{node.x - 1, node.y}] = true
								queue = append(queue, Point{x: node.x - 1, y: node.y})
							}
						} else {
							if coordinate[node.x+1][node.y] > i && !visited[[2]int{node.x + 1, node.y}] {
								visited[[2]int{node.x + 1, node.y}] = true
								queue = append(queue, Point{x: node.x + 1, y: node.y})
							}
							if coordinate[node.x-1][node.y] > i && !visited[[2]int{node.x - 1, node.y}] {
								visited[[2]int{node.x - 1, node.y}] = true
								queue = append(queue, Point{x: node.x - 1, y: node.y})
							}
						}
						if node.y == 0 {
							if coordinate[node.x][node.y+1] > i && !visited[[2]int{node.x, node.y + 1}] {
								visited[[2]int{node.x, node.y + 1}] = true
								queue = append(queue, Point{x: node.x, y: node.y + 1})
							}
						} else if node.y == n-1 {
							if coordinate[node.x][node.y-1] > i && !visited[[2]int{node.x, node.y - 1}] {
								visited[[2]int{node.x, node.y - 1}] = true
								queue = append(queue, Point{x: node.x, y: node.y - 1})
							}
						} else {
							if coordinate[node.x][node.y+1] > i && !visited[[2]int{node.x, node.y + 1}] {
								visited[[2]int{node.x, node.y + 1}] = true
								queue = append(queue, Point{x: node.x, y: node.y + 1})
							}
							if coordinate[node.x][node.y-1] > i && !visited[[2]int{node.x, node.y - 1}] {
								visited[[2]int{node.x, node.y - 1}] = true
								queue = append(queue, Point{x: node.x, y: node.y - 1})
							}
						}

						queue = queue[1:]

					}
				}
			}
		}

		if res < safe {
			res = safe
		}
	}

	return res
}

func main() {
	defer writer.Flush()

	var n int

	fmt.Fscan(reader, &n)
	coordinate := make([][]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			input := 0
			fmt.Fscan(reader, &input)
			coordinate[i] = append(coordinate[i], input)
		}
	}

	res := BFS(n, coordinate)
	fmt.Fprintln(writer, res)

}
