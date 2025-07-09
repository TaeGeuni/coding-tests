# 숨바꼭질 (Hide and Seek) 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/1697)

## 문제 설명

수빈이는 동생과 숨바꼭질을 하고 있습니다. 수빈이는 현재 위치 `N`에 있고, 동생은 위치 `K`에 있습니다.  
수빈이는 다음 세 가지 방법으로 이동할 수 있습니다:

- `X-1` (1초 소요)
- `X+1` (1초 소요)
- `2*X` (1초 소요)

수빈이가 동생을 찾을 수 있는 **가장 빠른 시간(초)**을 구하는 프로그램을 작성하세요.

---

## 입력 형식

- 첫째 줄: 두 정수 `N`과 `K` (0 ≤ N, K ≤ 100,000)

## 출력 형식

- 수빈이가 동생을 찾는 데 걸리는 최소 시간(초)을 출력합니다.

---

## 예제 입력

```
5 17
```

## 예제 출력

```
4
```

## 풀이 전략

이 문제는 **최단 거리 탐색** 문제로, `BFS(너비 우선 탐색)`을 통해 해결할 수 있습니다.  
각 위치에서 수빈이가 이동할 수 있는 경우의 수는 3가지이며, 각 경우에 대해 1초씩 증가하므로 BFS 탐색을 통해 가장 빠르게 도달할 수 있는 시간을 찾습니다.

- `visited[]` 배열을 통해 이미 방문한 위치는 다시 방문하지 않도록 처리
- 큐(queue)를 통해 위치와 소요 시간 정보를 저장
- 목표 위치에 도달하면 해당 시간 반환

---

## Go 코드

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

const max = 100_000

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func hideAndSeek(n, k int) int {
	visited := make([]bool, max+1)
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{n, 0})
	visited[n] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		pos, time := current[0], current[1]
		if pos == k {
			return time
		}

		next := [3]int{pos - 1, pos + 1, pos * 2}

		for i := 0; i < 3; i++ {
			if next[i] >= 0 && next[i] <= max && !visited[next[i]] {
				queue = append(queue, [2]int{next[i], time + 1})
				visited[next[i]] = true
			}
		}
	}

	return -1
}

func main() {
	defer writer.Flush()

	n, k := 0, 0
	fmt.Fscan(reader, &n, &k)
	fmt.Fprintln(writer, hideAndSeek(n, k))
}
```

---

## 시간 복잡도

- BFS는 각 위치를 최대 1번씩 방문하므로 **O(N)**의 시간 복잡도를 가집니다.

## 핵심 포인트

- BFS는 최단 경로 탐색에 매우 적합한 알고리즘입니다.
- 이동 횟수가 같고, 가중치가 동일할 경우 BFS로 해결하는 것이 가장 효율적입니다.
