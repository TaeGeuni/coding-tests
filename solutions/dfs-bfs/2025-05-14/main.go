package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X int
	Y int
}

type MoveAndPoint struct {
	point Point
	move  int
}

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func BFS(n, m int, maze []string) int {
	queue := []MoveAndPoint{{point: Point{X: 0, Y: 0}, move: 1}}
	visited := make(map[Point]bool)
	visited[Point{X: 0, Y: 0}] = true

	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.point.X == m-1 && node.point.Y == n-1 {
			return node.move
		}

		for i := 0; i < 4; i++ {
			nx := node.point.X + dx[i]
			ny := node.point.Y + dy[i]

			if nx >= 0 && nx < m && ny >= 0 && ny < n {
				next := Point{X: nx, Y: ny}
				if maze[ny][nx] == '1' && !visited[next] {
					visited[next] = true
					queue = append(queue, MoveAndPoint{point: next, move: node.move + 1})
				}
			}
		}
	}

	return -1
}

func main() {
	defer writer.Flush()

	var n, m int

	fmt.Fscan(reader, &n, &m)
	maze := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &maze[i])
	}

	result := BFS(n, m, maze)

	fmt.Fprintln(writer, result)

}
