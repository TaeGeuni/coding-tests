# BFS 문제 풀이 (Golang)
👉🏻[문제 링크](https://www.acmicpc.net/problem/24444)

## 문제 설명

오늘도 서준이는 너비 우선 탐색(BFS) 수업 조교를 하고 있다. 아빠가 수업한 내용을 학생들이 잘 이해했는지 문제를 통해서 확인해보자.

N개의 정점과 M개의 간선으로 구성된 무방향 그래프(undirected graph)가 주어진다. 정점 번호는 1번부터 N번이고 모든 간선의 가중치는 1이다. 정점 R에서 시작하여 너비 우선 탐색으로 노드를 방문할 경우 노드의 방문 순서를 출력하자.

### 입력
- 첫째 줄에 정점의 수 **N (5 ≤ N ≤ 100,000)**, 간선의 수 **M (1 ≤ M ≤ 200,000)**, 시작 정점 **R (1 ≤ R ≤ N)**이 주어진다.
- 다음 M개 줄에 간선 정보 **u v**가 주어지며, 이는 정점 u와 정점 v를 연결하는 양방향 간선을 나타낸다.

### 출력
- 첫째 줄부터 N개의 줄에 정수를 한 개씩 출력한다.
- i번째 줄에는 정점 i의 방문 순서를 출력한다.
- 시작 정점에서 방문할 수 없는 경우 0을 출력한다.

### 예제 입력
```
5 5 1
1 4
1 2
2 3
2 4
3 4
```

### 예제 출력
```
1
2
4
3
0
```

---

## 문제 풀이

### BFS(Breadth-First Search, 너비 우선 탐색)
BFS는 그래프 탐색 알고리즘 중 하나로, 가장 가까운 정점부터 차례대로 탐색하는 방식입니다. 이를 위해 **큐(Queue)** 자료구조를 사용하여 탐색할 정점을 관리합니다.

### BFS 의사 코드
```plaintext
bfs(V, E, R) {  # V: 정점 집합, E: 간선 집합, R: 시작 정점
    for each v ∈ V - {R}
        visited[v] <- NO;
    visited[R] <- YES;  # 시작 정점 R을 방문했다고 표시
    enqueue(Q, R);  # 큐 맨 뒤에 시작 정점 R 추가

    while (Q ≠ ∅) {
        u <- dequeue(Q);  # 큐 맨 앞쪽 요소 삭제
        for each v ∈ E(u)  # E(u): 정점 u의 인접 정점 집합
            if (visited[v] = NO) then {
                visited[v] <- YES;  # 정점 v 방문 표시
                enqueue(Q, v);  # 큐 맨 뒤에 정점 v 추가
            }
    }
}
```

---

### Golang 코드
```go
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
```

---

## 코드 설명

1. **그래프를 인접 리스트(map[int][]int)로 저장**
   - 각 노드의 연결 정보를 저장합니다.
   - 입력이 주어질 때 `map[int][]int` 구조로 저장하여 인접 리스트를 만듭니다.

2. **BFS 알고리즘 구현**
   - `queue`를 사용하여 탐색 순서를 관리합니다.
   - `visited` 배열로 방문 여부를 체크합니다.
   - 방문한 정점의 방문 순서를 `visitOrder`에 저장합니다.

3. **인접 정점 오름차순 정렬**
   - 문제 조건에 맞게 오름차순으로 방문하도록 정렬합니다.

4. **결과 출력**
   - `visitOrder` 배열을 순차적으로 출력합니다.

---

## 시간 복잡도

1. **입력 처리:** `O(M)`  
2. **그래프 정렬:** `O(N log N)` (각 정점의 연결 리스트 정렬)  
3. **BFS 탐색:** `O(N + M)`  
**→ 총 시간 복잡도: `O(N log N + M)`** (정렬을 고려한 최악의 경우)

---

## 공간 복잡도

1. **그래프 저장:** `O(N + M)`
2. **방문 배열:** `O(N)`
3. **큐:** `O(N)`
**→ 총 공간 복잡도: `O(N + M)`**

---

## 참고 사항
- **입출력 최적화:** `bufio.Reader`와 `bufio.Writer`를 사용하여 큰 입력/출력 처리 속도를 높임
- **인접 리스트 사용:** 메모리를 효율적으로 관리하면서 간선 정보 처리 가능
