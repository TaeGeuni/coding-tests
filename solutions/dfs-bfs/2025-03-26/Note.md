# 배추흰지렁이 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/1012)

## 문제 설명

차세대 영농인 한나는 고랭지에서 유기농 배추를 재배하고 있습니다. 배추를 해충으로부터 보호하기 위해 배추흰지렁이를 구입하려고 합니다. 배추흰지렁이는 인접한 배추들로 이동할 수 있으므로, 인접한 배추들로 이루어진 "구역"마다 최소 한 마리의 지렁이가 필요합니다. 주어진 배추밭에서 필요한 최소 지렁이 마리 수를 구하는 프로그램을 작성하세요.

### 제약 조건
- 배추밭의 크기: 1 ≤ M, N ≤ 50
- 배추 위치의 개수: 1 ≤ K ≤ 2500
- 각 테스트 케이스는 독립적으로 처리됩니다.

---

## 입력 형식
1. 첫 줄에 테스트 케이스의 개수 `T`.
2. 각 테스트 케이스마다:
   - 첫 줄에 가로길이 `M`, 세로길이 `N`, 배추 위치 개수 `K`.
   - 이후 `K`개의 줄에 배추 위치 `X`, `Y`.

## 출력 형식
각 테스트 케이스에 대해 필요한 최소 배추흰지렁이 마리 수를 한 줄에 하나씩 출력합니다.

---

## 예제 입력

```
2
10 8 17
0 0
1 0
1 1
4 2
4 3
4 5
2 4
3 4
7 4
8 4
9 4
7 5
8 5
9 5
7 6
8 6
9 6
10 10 1
5 5
```

## 예제 출력

```
5
1
```

---

## 코드 구현

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
	Y int
	X int
}

func BFS(m, n int, field [][]int) int {
	res := 0

	queue := make([]Point, 0)
	visited := make(map[[2]int]bool)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if field[i][j] == 1 {
				if !visited[[2]int{i, j}] {
					queue = append(queue, Point{Y: i, X: j})
					visited[[2]int{i, j}] = true
					res++
				}
				for len(queue) > 0 {
					node := queue[0]
					if node.X == 0 {
						if field[node.Y][node.X+1] == 1 && !visited[[2]int{node.Y, node.X + 1}] {
							queue = append(queue, Point{Y: node.Y, X: node.X + 1})
							visited[[2]int{node.Y, node.X + 1}] = true
						}
					} else if node.X == m-1 {
						if field[node.Y][node.X-1] == 1 && !visited[[2]int{node.Y, node.X - 1}] {
							queue = append(queue, Point{Y: node.Y, X: node.X - 1})
							visited[[2]int{node.Y, node.X - 1}] = true
						}
					} else {
						if field[node.Y][node.X+1] == 1 && !visited[[2]int{node.Y, node.X + 1}] {
							queue = append(queue, Point{Y: node.Y, X: node.X + 1})
							visited[[2]int{node.Y, node.X + 1}] = true
						}
						if field[node.Y][node.X-1] == 1 && !visited[[2]int{node.Y, node.X - 1}] {
							queue = append(queue, Point{Y: node.Y, X: node.X - 1})
							visited[[2]int{node.Y, node.X - 1}] = true
						}
					}
					if node.Y == 0 {
						if field[node.Y+1][node.X] == 1 && !visited[[2]int{node.Y + 1, node.X}] {
							queue = append(queue, Point{Y: node.Y + 1, X: node.X})
							visited[[2]int{node.Y + 1, node.X}] = true
						}
					} else if node.Y == n-1 {
						if field[node.Y-1][node.X] == 1 && !visited[[2]int{node.Y - 1, node.X}] {
							queue = append(queue, Point{Y: node.Y - 1, X: node.X})
							visited[[2]int{node.Y - 1, node.X}] = true
						}
					} else {
						if field[node.Y+1][node.X] == 1 && !visited[[2]int{node.Y + 1, node.X}] {
							queue = append(queue, Point{Y: node.Y + 1, X: node.X})
							visited[[2]int{node.Y + 1, node.X}] = true
						}
						if field[node.Y-1][node.X] == 1 && !visited[[2]int{node.Y - 1, node.X}] {
							queue = append(queue, Point{Y: node.Y - 1, X: node.X})
							visited[[2]int{node.Y - 1, node.X}] = true
						}
					}
					queue = queue[1:]
				}
			}
		}
	}

	return res
}

func main() {
	defer writer.Flush()

	var t int
	var m, n, k int
	res := make([]int, 0)

	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		fmt.Fscan(reader, &m, &n, &k)
		field := make([][]int, n)
		for j := 0; j < n; j++ {
			field[j] = make([]int, m)
		}
		for j := 0; j < k; j++ {
			var x, y int
			fmt.Fscan(reader, &x, &y)
			field[y][x] = 1
		}
		res = append(res, BFS(m, n, field))
	}

	for i := 0; i < len(res); i++ {
		fmt.Fprintln(writer, res[i])
	}
}

```

---

## 풀이 설명

1. **자료구조 정의**:
   - BFS를 위한 `Point` 구조체.
   - 방문 여부를 확인하기 위한 `visited` 맵.

2. **BFS 탐색**:
   - 배추가 있는 위치에서 시작.
   - 인접한 배추를 큐에 추가하고 방문 처리.

3. **최소 지렁이 수 계산**:
   - BFS 탐색을 시작할 때마다 지렁이 수를 1 증가.

4. **결과 출력**:
   - 테스트 케이스별로 필요한 지렁이 수를 출력.

---

## 개선 포인트
- **코드 최적화**:
  - 방문 여부 확인을 위해 맵 대신 2차원 배열 사용 가능.
  - BFS 내 중복 로직을 함수로 분리하여 가독성 향상.

---

## 복잡도 분석

- **시간 복잡도**: O(M * N) (배추밭의 모든 위치를 한 번씩 방문)
- **공간 복잡도**: O(M * N) (배추밭 크기 및 방문 여부 저장)
