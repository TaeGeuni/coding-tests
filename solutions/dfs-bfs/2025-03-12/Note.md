# 웜 바이러스 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/2606)

## 문제 설명
신종 바이러스인 웜 바이러스는 네트워크를 통해 전파된다. 한 컴퓨터가 웜 바이러스에 걸리면 그 컴퓨터와 네트워크 상에서 연결되어 있는 모든 컴퓨터도 웜 바이러스에 감염된다.

예를 들어 7대의 컴퓨터가 네트워크 상에서 연결되어 있을 때, 1번 컴퓨터가 감염되면 2번과 5번 컴퓨터를 거쳐 3번과 6번 컴퓨터까지 전파된다. 하지만 4번과 7번 컴퓨터는 1번 컴퓨터와 직접 연결되지 않았으므로 감염되지 않는다.

1번 컴퓨터가 감염되었을 때, 감염되는 컴퓨터의 수를 구하는 프로그램을 작성하시오.

## 입력
- 첫째 줄: 컴퓨터의 수 (100 이하의 양의 정수)
- 둘째 줄: 네트워크 상에서 직접 연결되어 있는 컴퓨터 쌍의 수
- 다음 줄부터: 네트워크 상에서 연결된 컴퓨터 번호 쌍이 주어진다.

## 출력
- 1번 컴퓨터가 웜 바이러스에 감염되었을 때, 감염되는 컴퓨터의 수를 출력한다.

## 예제 입력
```
7
6
1 2
2 3
1 5
5 2
5 6
4 7
```

## 예제 출력
```
4
```

---

## 풀이 방법
이 문제는 그래프 탐색 알고리즘을 이용하여 해결할 수 있다. BFS(너비 우선 탐색)를 사용하여 1번 컴퓨터에서 연결된 모든 컴퓨터를 탐색하면서 감염된 컴퓨터 수를 구한다.

### BFS를 이용한 풀이
- 그래프를 인접 리스트로 표현한다.
- 1번 노드를 시작점으로 BFS 탐색을 수행한다.
- 방문한 노드를 `infected` 맵에 저장하여 중복 방문을 방지한다.
- 큐를 이용해 탐색하며 연결된 컴퓨터를 계속 방문한다.
- 탐색이 끝나면 감염된 컴퓨터의 수를 반환한다.

## 코드 구현 (Golang)
```go
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

```

## 코드 설명
1. **입력 처리**: `bufio`를 사용하여 빠르게 입력을 받는다.
2. **그래프 구성**: `graph`를 인접 리스트 형태로 만들어 각 컴퓨터의 연결 관계를 저장한다.
3. **BFS 탐색**:
   - `queue`를 사용하여 탐색할 노드를 관리한다.
   - `infected` 맵을 사용하여 방문한 노드를 저장하고 중복 방문을 방지한다.
   - 큐에서 노드를 하나씩 꺼내며 해당 노드와 연결된 노드를 탐색하고, 아직 감염되지 않은 노드를 감염시킨다.
4. **출력**: 감염된 컴퓨터의 수를 출력한다.

## 시간 복잡도
- 그래프 탐색은 `O(N + M)` (N: 노드 수, M: 간선 수)의 시간 복잡도를 가진다.
- `infected` 맵을 사용하여 중복 방문을 방지하여 효율적인 탐색이 가능하다.

## 결론
이 문제는 그래프 탐색 문제로 BFS를 활용하여 해결할 수 있다. BFS를 이용하면 연결된 컴퓨터를 빠르고 정확하게 탐색할 수 있으며, 중복 방문을 방지하기 위해 맵을 활용하는 것이 중요하다. 🚀
