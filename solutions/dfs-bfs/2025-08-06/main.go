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
	X int
	Y int
}

type PointAndCount struct {
	p     Point
	count int
}

func treasureHunt(v, h int, treasureMap []string) int {
	res := 0
	dx := []int{0, 0, -1, 1}
	dy := []int{-1, 1, 0, 0}

	for i := 0; i < v; i++ {
		for j := 0; j < h; j++ {
			if treasureMap[i][j] == 'L' {
				visited := make([][]bool, v)
				for k := range visited {
					visited[k] = make([]bool, h)
				}
				queue := make([]PointAndCount, 0, v*h)
				queue = append(queue, PointAndCount{Point{X: j, Y: i}, 0})
				visited[i][j] = true
				head := 0
				maxDistInBfs := 0

				for head < len(queue) {
					node := queue[head]
					head++

					if node.count > maxDistInBfs {
						maxDistInBfs = node.count
					}

					for k := 0; k < 4; k++ {
						nx, ny := node.p.X+dx[k], node.p.Y+dy[k]

						if nx < 0 || nx >= h || ny < 0 || ny >= v {
							continue
						}
						if !visited[ny][nx] && treasureMap[ny][nx] == 'L' {
							visited[ny][nx] = true
							queue = append(queue, PointAndCount{Point{X: nx, Y: ny}, node.count + 1})
						}
					}
				}

				if maxDistInBfs > res {
					res = maxDistInBfs
				}
			}
		}
	}

	return res
}

func main() {
	defer writer.Flush()

	v, h := 0, 0

	fmt.Fscan(reader, &v, &h)

	treasureMap := make([]string, v)
	for i := 0; i < v; i++ {
		fmt.Fscan(reader, &treasureMap[i])
	}

	fmt.Fprintln(writer, treasureHunt(v, h, treasureMap))

}
