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

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func Distance(a, b Point) int {
	res := 0

	if a.X > 0 && b.X > 0 {
		if a.X > b.X {
			res += a.X - b.X
		} else {
			res += b.X - a.X
		}
	} else if a.X <= 0 && b.X <= 0 {
		if a.X > b.X {
			res += a.X - b.X
		} else {
			res += b.X - a.X
		}
	} else {
		if a.X-b.X < 0 {
			res += -(a.X - b.X)
		} else {
			res += a.X - b.X
		}
	}

	if a.Y > 0 && b.Y > 0 {
		if a.Y > b.Y {
			res += a.Y - b.Y
		} else {
			res += b.Y - a.Y
		}
	} else if a.Y <= 0 && b.Y <= 0 {
		if a.Y > b.Y {
			res += a.Y - b.Y
		} else {
			res += b.Y - a.Y
		}
	} else {
		if a.Y-b.Y < 0 {
			res += -(a.Y - b.Y)
		} else {
			res += a.Y - b.Y
		}
	}

	return res
}

func BFS(n int, home, festival Point, convini []Point) string {
	visited := make(map[Point]bool)

	queue := make([]Point, 0)
	queue = append(queue, home)

	for len(queue) > 0 {
		node := queue[0]
		visited[node] = true

		totalMove := Distance(festival, node)

		if totalMove <= 1000 {
			return "happy"
		}

		for i := 0; i < n; i++ {
			totalMove := Distance(convini[i], node)

			if totalMove <= 1000 && !visited[convini[i]] {
				queue = append(queue, convini[i])
			}

		}

		queue = queue[1:]

	}

	return "sad"
}

func main() {
	defer writer.Flush()

	var t, n int
	res := make([]string, 0)
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var home, festival Point
		fmt.Fscan(reader, &n)
		convini := make([]Point, n)

		fmt.Fscan(reader, &home.X, &home.Y)

		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &convini[i].X, &convini[i].Y)
		}
		fmt.Fscan(reader, &festival.X, &festival.Y)
		res = append(res, BFS(n, home, festival, convini))
	}

	for i := 0; i < t; i++ {
		fmt.Fprintln(writer, res[i])
	}
}
