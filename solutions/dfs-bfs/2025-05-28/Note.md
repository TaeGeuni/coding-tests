# 맥주 마시면서 걸어가기
👉🏻[문제 링크](https://www.acmicpc.net/problem/9205)

## 📝 문제 설명

- 상근이와 친구들이 락 페스티벌에 맥주를 마시면서 걸어가려 한다.
- 시작은 집이고, 50m마다 맥주를 한 병 마셔야 하며, 한 박스에는 20병이 있다.
- 편의점에서 맥주를 다시 채울 수 있지만 한 번에 최대 20병만.
- 출발지, 편의점들, 목적지(페스티벌) 좌표가 주어진다.
- 편의점들을 잘 이용하여 페스티벌까지 갈 수 있으면 `"happy"`, 그렇지 않으면 `"sad"`를 출력해야 한다.

## 📌 입력 조건

- 테스트 케이스 개수 t (1 ≤ t ≤ 50)
- 각 테스트 케이스는 다음 순서로 이루어짐:
  1. 편의점 개수 n (0 ≤ n ≤ 100)
  2. 상근이네 집 좌표
  3. n개의 편의점 좌표
  4. 페스티벌 좌표

## 🚀 접근 방식

- 한 번에 1000m 이하로만 이동할 수 있으므로, 각 지점을 노드로 보고, 거리가 1000m 이하이면 연결된 그래프로 본다.
- BFS를 사용하여 집에서 출발해 페스티벌까지 갈 수 있는지 확인한다.
- 노드: 집, 편의점들, 페스티벌
- 간선: 1000m 이하 거리인 경우

## ✅ 주요 함수

```go
// 두 좌표 간 맨해튼 거리 계산
func Distance(a, b Point) int {
    return abs(a.X - b.X) + abs(a.Y - b.Y)
}
```

```go
// BFS를 통해 도달 가능 여부 판별
func BFS(n int, home, festival Point, convini []Point) string {
    visited := make(map[Point]bool)
    queue := []Point{home}

    for len(queue) > 0 {
        node := queue[0]
        visited[node] = true

        // 페스티벌까지 갈 수 있는지 확인
        if Distance(festival, node) <= 1000 {
            return "happy"
        }

        // 편의점 탐색
        for i := 0; i < n; i++ {
            if Distance(convini[i], node) <= 1000 && !visited[convini[i]] {
                queue = append(queue, convini[i])
            }
        }

        queue = queue[1:] // pop front
    }

    return "sad"
}
```

## 💡 예제

### 입력

```
2
2
0 0
1000 0
1000 1000
2000 1000
2
0 0
1000 0
2000 1000
2000 2000
```

### 출력

```
happy
sad
```

## 🧠 복잡도

- 각 테스트 케이스마다 BFS 탐색을 수행하며, 최대 노드 수는 102개 (집 + 편의점 100 + 페스티벌).
- 시간 복잡도: O(N^2)로 충분히 빠르게 동작 가능.

## ✅ 결론

- 편의점이 많아도 맨해튼 거리 1000m 기준으로만 이동할 수 있는 그래프를 구성하고 BFS를 통해 도달 가능한지 확인하는 문제.
- 실전에서 BFS 연습에 적합하며, 거리 기반 탐색에 익숙해질 수 있음.
