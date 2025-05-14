# 미로 탐색 문제 풀이 (BFS)
👉🏻[문제 링크](https://www.acmicpc.net/problem/2178)

## 문제 설명

N×M 크기의 배열로 표현된 미로에서 '1'은 이동 가능한 칸, '0'은 이동 불가능한 칸입니다. 
시작점 (1,1)에서 도착점 (N,M)까지 이동할 수 있는 최단 경로의 칸 수를 구해야 합니다.
- 이동은 상하좌우 인접한 칸으로만 가능하며,
- 시작점과 도착점을 포함한 칸 수를 세야 합니다.

---

## 입력 형식
- 첫째 줄: N, M (2 ≤ N, M ≤ 100)
- 다음 N개의 줄: M개의 숫자로 구성된 문자열 (공백 없이 붙어 있음)

## 출력 형식
- 최소 칸 수 출력

---

## 예제 입력
```
4 6
101111
101010
101011
111011
```

## 예제 출력
```
15
```

---

## 알고리즘 설명 (BFS)

### 아이디어
- 최단 경로 문제는 BFS가 적합
- (0,0)에서 시작하여, 네 방향(상하좌우)으로 이동하며 탐색
- 이동할 수 있는 칸이고 아직 방문하지 않은 칸만 큐에 추가
- BFS를 통해 가장 먼저 (N-1, M-1)에 도달하는 경로가 최단 경로

### 시간 복잡도
- O(N×M): 모든 칸을 한 번씩만 방문

---

## Go 코드 구현
```go
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
```

---

## 결론
- BFS를 통해 인접 칸을 모두 탐색하며 최단 경로를 탐색
- 방문 여부는 `map[Point]bool`로 처리
- 실전에서 BFS를 익히기에 적절한 문제로, 다양한 미로 탐색 문제의 기본이 됨
