package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// 입출력 처리
var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func bfs(n, r int, graph map[int][]int) []int {
	// 방문 순서를 저장할 배열 (0으로 초기화)
	visitOrder := make([]int, n+1)

	// BFS 탐색을 위한 큐
	queue := []int{r}

	// 방문 여부 체크
	visited := make([]bool, n+1)
	visited[r] = true

	order := 1
	visitOrder[r] = order

	// BFS 실행
	for len(queue) > 0 {
		// 큐에서 정점 꺼내기 (FIFO)
		node := queue[0]
		queue = queue[1:]

		// 인접 노드 탐색 (오름차순 정렬)
		for _, next := range graph[node] {
			if !visited[next] {
				visited[next] = true
				order++
				visitOrder[next] = order
				queue = append(queue, next)
			}
		}
	}

	return visitOrder[1:] // 1번 인덱스부터 출력
}

func main() {
	defer writer.Flush()

	var n, m, r int
	fmt.Fscan(reader, &n, &m, &r)

	// 그래프 초기화
	graph := make(map[int][]int)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// 인접 리스트 정렬 (오름차순 방문)
	for key := range graph {
		sort.Ints(graph[key])
	}

	// BFS 실행
	result := bfs(n, r, graph)

	// 결과 출력
	for _, v := range result {
		fmt.Fprintln(writer, v)
	}
}
