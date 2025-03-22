package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type Point struct {
	Y int
	X int
}

func BFS(n int, graph []string) (int, []int) {
	res := 0
	complex := make([]int, 0)
	visited := make(map[[2]int]bool)
	queue := make([]Point, 0)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if graph[i][j] == '1' && !visited[[2]int{i, j}] {
				number := 0

				visited[[2]int{i, j}] = true
				queue = append(queue, Point{Y: i, X: j})
				for len(queue) > 0 {
					node := queue[0]
					if node.X == 0 {
						if graph[node.Y][node.X+1] == '1' && !visited[[2]int{node.Y, node.X + 1}] {
							visited[[2]int{node.Y, node.X + 1}] = true
							queue = append(queue, Point{Y: node.Y, X: node.X + 1})
						}
					} else if node.X == n-1 {
						if graph[node.Y][node.X-1] == '1' && !visited[[2]int{node.Y, node.X - 1}] {
							visited[[2]int{node.Y, node.X - 1}] = true
							queue = append(queue, Point{Y: node.Y, X: node.X - 1})
						}
					} else {
						if graph[node.Y][node.X+1] == '1' && !visited[[2]int{node.Y, node.X + 1}] {
							visited[[2]int{node.Y, node.X + 1}] = true
							queue = append(queue, Point{Y: node.Y, X: node.X + 1})
						}
						if graph[node.Y][node.X-1] == '1' && !visited[[2]int{node.Y, node.X - 1}] {
							visited[[2]int{node.Y, node.X - 1}] = true
							queue = append(queue, Point{Y: node.Y, X: node.X - 1})
						}
					}
					if node.Y == 0 {
						if graph[node.Y+1][node.X] == '1' && !visited[[2]int{node.Y + 1, node.X}] {
							visited[[2]int{node.Y + 1, node.X}] = true
							queue = append(queue, Point{Y: node.Y + 1, X: node.X})
						}
					} else if node.Y == n-1 {
						if graph[node.Y-1][node.X] == '1' && !visited[[2]int{node.Y - 1, node.X}] {
							visited[[2]int{node.Y - 1, node.X}] = true
							queue = append(queue, Point{Y: node.Y - 1, X: node.X})
						}
					} else {
						if graph[node.Y+1][node.X] == '1' && !visited[[2]int{node.Y + 1, node.X}] {
							visited[[2]int{node.Y + 1, node.X}] = true
							queue = append(queue, Point{Y: node.Y + 1, X: node.X})
						}
						if graph[node.Y-1][node.X] == '1' && !visited[[2]int{node.Y - 1, node.X}] {
							visited[[2]int{node.Y - 1, node.X}] = true
							queue = append(queue, Point{Y: node.Y - 1, X: node.X})
						}
					}
					number++
					queue = queue[1:]
				}
				res++
				complex = append(complex, number)
			}
		}
	}

	sort.Ints(complex)

	return res, complex
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	graph := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &graph[i])
	}

	res, complex := BFS(n, graph)
	fmt.Println(res)
	for i := 0; i < len(complex); i++ {
		fmt.Println(complex[i])
	}

}
