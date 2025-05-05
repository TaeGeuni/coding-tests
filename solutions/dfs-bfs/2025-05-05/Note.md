# 안전 영역 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/2468)

## 문제 설명

장마철을 대비해 지역의 높이 정보를 바탕으로 **비가 왔을 때 물에 잠기지 않는 안전한 영역의 최대 개수**를 구하는 문제입니다.

- 지역은 N×N 크기의 2차원 배열로 주어지고, 각 원소는 해당 지점의 높이를 나타냅니다.
- 어떤 비의 양에 대해 높이 이하의 모든 지역은 잠긴다고 가정합니다.
- **상하좌우로 인접한** 잠기지 않은 지역이 하나의 안전 영역이 됩니다.

## 입력 형식
- 첫 줄: 지역의 크기 N (2 ≤ N ≤ 100)
- 다음 N줄: 각 줄마다 N개의 정수 (1 ≤ 높이 ≤ 100)

## 출력 형식
- 비의 양을 0부터 100까지 변화시켜가며 구한 **안전한 영역의 최대 개수** 출력

---

## 예제 입력
### 입력 1
```
5
6 8 2 6 2
3 2 3 4 6
6 7 3 3 2
7 2 5 3 6
8 9 5 2 7
```
### 출력 1
```
5
```

### 입력 2
```
7
9 9 9 9 9 9 9
9 2 1 2 1 2 9
9 1 8 7 8 1 9
9 2 7 9 7 2 9
9 1 8 7 8 1 9
9 2 1 2 1 2 9
9 9 9 9 9 9 9
```
### 출력 2
```
6
```

---

## 풀이 접근 (BFS 활용)

- 각 비의 양(0~100)에 대해:
  - 해당 높이 이하인 지역은 잠긴 것으로 간주
  - BFS를 사용하여 잠기지 않은 구역을 탐색 → 안전 영역 1개로 카운트
  - BFS 탐색을 통해 안전한 영역의 개수를 구하고, 그 중 최대값을 갱신

---

## Go 코드
```go
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
	x int
	y int
}

func BFS(n int, coordinate [][]int) int {
	res := 0

	for i := 0; i < 101; i++ {
		safe := 0
		queue := make([]Point, 0)
		visited := make(map[[2]int]bool)

		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if i >= coordinate[j][k] || visited[[2]int{j, k}] {
					continue
				}

				safe++
				queue = append(queue, Point{x: j, y: k})
				visited[[2]int{j, k}] = true

				for len(queue) != 0 {
					node := queue[0]
					queue = queue[1:]

					dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
					for _, d := range dirs {
						nx, ny := node.x+d[0], node.y+d[1]
						if nx >= 0 && nx < n && ny >= 0 && ny < n {
							if coordinate[nx][ny] > i && !visited[[2]int{nx, ny}] {
								visited[[2]int{nx, ny}] = true
								queue = append(queue, Point{x: nx, y: ny})
							}
						}
					}
				}
			}
		}

		if res < safe {
			res = safe
		}
	}

	return res
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	coordinate := make([][]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			input := 0
			fmt.Fscan(reader, &input)
			coordinate[i] = append(coordinate[i], input)
		}
	}

	res := BFS(n, coordinate)
	fmt.Fprintln(writer, res)
}
```

---

## 정리
- 이 문제는 모든 비의 양에 대해 안전한 영역을 BFS로 탐색해야 하는 **완전탐색 + BFS** 유형입니다.
- 시간 복잡도는 O(101 × N × N) 정도이며, N ≤ 100이므로 충분히 처리할 수 있습니다.

---
