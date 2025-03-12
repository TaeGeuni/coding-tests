package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func BFS(graph [][]int) int {
	res := 0

	queue := make([]int, 0)
	queue = append(queue, 1)

	infected := make(map[int]struct{})
	infected[1] = struct{}{}

	for len(queue) != 0 {
		node := queue[0]

		for i := 0; i < len(graph[node]); i++ {
			if _, c := infected[graph[node][i]]; !c {
				queue = append(queue, graph[node][i])
			}
		}
		if _, c := infected[node]; !c {
			infected[node] = struct{}{}
			res++
		}
		queue = queue[1:]

	}

	return res
}

func main() {
	defer writer.Flush()

	var a, b int

	fmt.Fscan(reader, &a)
	fmt.Fscan(reader, &b)

	graph := make([][]int, a+1)

	for i := 0; i < b; i++ {
		var m, n int
		fmt.Fscan(reader, &m, &n)

		graph[m] = append(graph[m], n)
		graph[n] = append(graph[n], m)
	}

	fmt.Fprintln(writer, BFS(graph))
}
