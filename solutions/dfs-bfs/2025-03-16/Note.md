# DFS와 BFS 탐색 결과 출력
👉🏻[문제 링크](https://www.acmicpc.net/problem/1260)

## 문제 설명
그래프를 DFS(깊이 우선 탐색)와 BFS(너비 우선 탐색)로 탐색한 결과를 출력하는 프로그램을 작성한다.
- 방문할 수 있는 정점이 여러 개인 경우에는 **정점 번호가 작은 것부터 먼저 방문**한다.
- 더 이상 방문할 수 있는 정점이 없는 경우 탐색을 종료한다.
- 정점 번호는 1번부터 N번까지 주어진다.

## 입력 형식
1. 첫째 줄에 정점의 개수 **N(1 ≤ N ≤ 1,000)**, 간선의 개수 **M(1 ≤ M ≤ 10,000)**, 탐색을 시작할 정점의 번호 **V**가 주어진다.
2. 다음 M개의 줄에는 간선이 연결하는 두 정점의 번호가 주어진다.
3. 입력으로 주어지는 간선은 **양방향**이다.

## 출력 형식
1. 첫째 줄에 **DFS 탐색 결과**를 출력한다.
2. 둘째 줄에 **BFS 탐색 결과**를 출력한다.
3. 정점 방문 순서는 공백으로 구분한다.

## 예제 입력과 출력

### 예제 입력 1
```
4 5 1
1 2
1 3
1 4
2 4
3 4
```

### 예제 출력 1
```
1 2 4 3
1 2 3 4
```

### 예제 입력 2
```
5 5 3
5 4
5 2
1 2
3 4
3 1
```

### 예제 출력 2
```
3 1 2 5 4
3 1 4 2 5
```

### 예제 입력 3
```
1000 1 1000
999 1000
```

### 예제 출력 3
```
1000 999
1000 999
```

## 코드 구현
```go
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
```

## 해결 방법
1. **그래프 저장**:
   - 인접 리스트를 사용하여 그래프를 저장한다.
   - 양방향 간선이므로 `graph[a]`와 `graph[b]`에 각각 값을 추가한다.
   - 방문할 수 있는 노드를 **오름차순 정렬**하여 작은 번호부터 방문하도록 한다.

2. **DFS (깊이 우선 탐색)**
   - 스택을 활용하여 구현한다.
   - 방문한 정점을 체크하기 위해 `map[int]struct{}`을 사용한다.
   - 스택에서 마지막 노드를 꺼내 방문하고, 방문하지 않은 인접 노드를 다시 스택에 추가한다.
   - 인접 노드는 역순으로 추가하여 오름차순 방문을 유지한다.

3. **BFS (너비 우선 탐색)**
   - 큐를 활용하여 구현한다.
   - 방문한 정점을 체크하기 위해 `map[int]struct{}`을 사용한다.
   - 큐에서 첫 번째 노드를 꺼내 방문하고, 방문하지 않은 인접 노드를 다시 큐에 추가한다.

4. **결과 출력**
   - DFS와 BFS 탐색 결과를 공백으로 구분하여 출력한다.
   - `fmt.Sprint()`와 `strings.Trim()`을 활용하여 리스트 형식을 제거한다.

## 시간 복잡도 분석
- **DFS와 BFS의 시간 복잡도**: `O(N + M)` (정점과 간선의 개수에 비례)
- `sort.Ints()`의 정렬 비용: `O(N log N)`
- 따라서, 최종 시간 복잡도는 `O(N log N + M)`

## 결론
이 코드는 DFS와 BFS를 이용하여 정점을 방문하는 순서를 출력하는 문제를 해결하며, 방문할 수 있는 노드가 여러 개일 경우 **작은 번호부터 우선적으로 방문하도록 정렬**하여 구현하였다.

