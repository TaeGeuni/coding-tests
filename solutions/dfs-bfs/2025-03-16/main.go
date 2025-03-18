package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func DFS(v int, graph [][]int) []int {
	res := make([]int, 0)

	stack := make([]int, 0)
	stack = append(stack, v)

	check := make(map[int]struct{})

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if _, ok := check[node]; ok {
			continue
		}

		check[node] = struct{}{}

		for i := len(graph[node]) - 1; i >= 0; i-- {
			next := graph[node][i]
			if _, ok := check[next]; !ok {
				stack = append(stack, next)
			}
		}
		res = append(res, node)
	}

	return res
}

func BFS(v int, graph [][]int) []int {
	res := make([]int, 0)
	queue := make([]int, 0)
	queue = append(queue, v)

	check := make(map[int]struct{})
	check[v] = struct{}{}

	for len(queue) > 0 {
		node := queue[0]

		for i := 0; i < len(graph[node]); i++ {
			if _, ok := check[graph[node][i]]; !ok {
				check[graph[node][i]] = struct{}{}
				queue = append(queue, graph[node][i])
			}
		}
		res = append(res, node)
		queue = queue[1:]
	}
	return res
}

func main() {
	defer writer.Flush()
	var n, m, v int

	fmt.Fscan(reader, &n, &m, &v)
	graph := make([][]int, n+1)

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)

		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	for i := 0; i < len(graph); i++ {
		sort.Ints(graph[i])
	}
	fmt.Fprintln(writer, strings.Trim(fmt.Sprint(DFS(v, graph)), "[]"))
	fmt.Fprintln(writer, strings.Trim(fmt.Sprint(BFS(v, graph)), "[]"))
}
