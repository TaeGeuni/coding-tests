# 문제 풀이: 단지 번호 붙이기
👉🏻[문제 링크](https://www.acmicpc.net/problem/2667)

## 문제 설명
정사각형 형태의 지도가 주어질 때, '1'은 집이 있는 곳, '0'은 집이 없는 곳을 나타냅니다. 상하좌우로 연결된 집들을 하나의 단지로 정의하고, 각 단지에 번호를 붙입니다. 총 단지 수와 각 단지에 속하는 집의 수를 오름차순으로 출력하는 프로그램을 작성합니다.

### 입력
- 첫 번째 줄: 지도의 크기 `N` (5 <= N <= 25)
- 다음 `N`줄: 지도 정보 (0과 1로 이루어진 문자열)

### 출력
1. 총 단지 수
2. 각 단지의 집의 수 (오름차순 정렬)

---

## 문제 해결
이 문제는 그래프 탐색 알고리즘인 **BFS**를 사용하여 해결합니다.

### 구현 방법
1. 입력받은 지도를 기반으로 BFS를 수행하여 모든 단지를 탐색합니다.
2. BFS를 통해 연결된 집들을 탐색하며, 각 단지의 집 개수를 카운트합니다.
3. 탐색이 끝난 후, 총 단지 수와 각 단지의 집 개수를 오름차순으로 출력합니다.

---

## 코드 구현
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type Point struct {
	Y int
	X int
}

func BFS(n int, graph []string) (int, []int) {
	res := 0
	complex := make([]int, 0)
	visited := make(map[[2]int]bool)
	queue := make([]Point, 0)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if graph[i][j] == '1' && !visited[[2]int{i, j}] {
				number := 0

				visited[[2]int{i, j}] = true
				queue = append(queue, Point{Y: i, X: j})
				for len(queue) > 0 {
					node := queue[0]
					if node.X == 0 {
						if graph[node.Y][node.X+1] == '1' && !visited[[2]int{node.Y, node.X + 1}] {
							visited[[2]int{node.Y, node.X + 1}] = true
							queue = append(queue, Point{Y: node.Y, X: node.X + 1})
						}
					} else if node.X == n-1 {
						if graph[node.Y][node.X-1] == '1' && !visited[[2]int{node.Y, node.X - 1}] {
							visited[[2]int{node.Y, node.X - 1}] = true
							queue = append(queue, Point{Y: node.Y, X: node.X - 1})
						}
					} else {
						if graph[node.Y][node.X+1] == '1' && !visited[[2]int{node.Y, node.X + 1}] {
							visited[[2]int{node.Y, node.X + 1}] = true
							queue = append(queue, Point{Y: node.Y, X: node.X + 1})
						}
						if graph[node.Y][node.X-1] == '1' && !visited[[2]int{node.Y, node.X - 1}] {
							visited[[2]int{node.Y, node.X - 1}] = true
							queue = append(queue, Point{Y: node.Y, X: node.X - 1})
						}
					}
					if node.Y == 0 {
						if graph[node.Y+1][node.X] == '1' && !visited[[2]int{node.Y + 1, node.X}] {
							visited[[2]int{node.Y + 1, node.X}] = true
							queue = append(queue, Point{Y: node.Y + 1, X: node.X})
						}
					} else if node.Y == n-1 {
						if graph[node.Y-1][node.X] == '1' && !visited[[2]int{node.Y - 1, node.X}] {
							visited[[2]int{node.Y - 1, node.X}] = true
							queue = append(queue, Point{Y: node.Y - 1, X: node.X})
						}
					} else {
						if graph[node.Y+1][node.X] == '1' && !visited[[2]int{node.Y + 1, node.X}] {
							visited[[2]int{node.Y + 1, node.X}] = true
							queue = append(queue, Point{Y: node.Y + 1, X: node.X})
						}
						if graph[node.Y-1][node.X] == '1' && !visited[[2]int{node.Y - 1, node.X}] {
							visited[[2]int{node.Y - 1, node.X}] = true
							queue = append(queue, Point{Y: node.Y - 1, X: node.X})
						}
					}
					number++
					queue = queue[1:]
				}
				res++
				complex = append(complex, number)
			}
		}
	}

	sort.Ints(complex)

	return res, complex
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	graph := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &graph[i])
	}

	res, complex := BFS(n, graph)
	fmt.Println(res)
	for i := 0; i < len(complex); i++ {
		fmt.Println(complex[i])
	}

}

```

---

## 동작 방식
1. **입력 처리**: `graph` 배열에 지도 데이터를 저장합니다.
2. **BFS 탐색**:
   - 시작점에서 상하좌우로 이동하며 연결된 '1'을 방문 처리합니다.
   - 큐를 사용하여 탐색을 진행하며 단지에 속한 집의 수를 카운트합니다.
3. **결과 출력**:
   - 탐색이 종료되면 총 단지 수와 각 단지의 집 개수를 출력합니다.

---

## 예제 실행
### 입력
```
7
0110100
0110101
1110101
0000111
0100000
0111110
0111000
```

### 출력
```
3
7
8
9
```

---

## 주요 포인트
1. **방문 처리**: `visited` 맵을 사용하여 중복 방문 방지.
2. **단지 개수 계산**: 탐색이 완료되면 `res`를 증가시키고, 각 단지 크기를 `complex` 배열에 저장.
3. **결과 정렬**: `sort.Ints`로 집의 개수를 오름차순 정렬 후 출력.

