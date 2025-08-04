# 1로 만들기
👉🏻[문제 링크](https://www.acmicpc.net/problem/12852)

## 📌 문제 설명

정수 `N`이 주어졌을 때, 아래 세 가지 연산을 이용하여 `N`을 `1`로 만들려고 한다.  
연산을 사용하는 **최솟값**을 출력하고, 그 과정에서 거치는 수들을 역추적하여 출력하라.

### 사용 가능한 연산
1. `X가 3으로 나누어떨어지면` → `X / 3`  
2. `X가 2으로 나누어떨어지면` → `X / 2`  
3. `X - 1`

---

## ✅ 입력

- 자연수 `N` (1 ≤ N ≤ 10⁶)

---

## ✅ 출력

- 연산을 하는 **최솟값**
- `N`을 1로 만드는 방법에 포함된 수들을 역순이 아닌 **실제 경로 순서대로** 출력

---

## 🧠 풀이 전략 - BFS (너비 우선 탐색)

- 각 상태를 노드로 보고, 가능한 연산으로 다음 상태로 이동하며 탐색
- 최단 거리이므로 BFS가 적절
- BFS 탐색 시, `현재 수 → 연산 → 다음 수` 구조로 탐색하며, 가장 먼저 `1`에 도달하는 경로가 최적해

### 핵심 구현 포인트

- `value` 구조체를 정의하여, 현재 숫자, 연산 횟수, 경로를 저장
- 방문 여부를 체크하여 중복 탐색 방지
- 연산 순서를 `나누기 → 빼기` 순으로 하여 최단 거리 우선 탐색

---

## 🧾 Go 코드

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

type value struct {
    count   int
    number  int
    history []int
}

func make1(n int) (int, []int) {
    if n == 1 {
        return 0, []int{1}
    }

    queue := make([]value, 0)
    queue = append(queue, value{count: 0, number: n, history: []int{n}})

    visited := make(map[int]bool)
    visited[n] = true

    var res value

    for len(queue) > 0 {
        node := queue[0]
        var v value

        if node.number%3 == 0 && !visited[node.number/3] {
            v.count = node.count + 1
            v.number = node.number / 3
            newHistory := make([]int, len(node.history))
            copy(newHistory, node.history)
            newHistory = append(newHistory, v.number)
            v.history = newHistory

            visited[v.number] = true
            queue = append(queue, v)

            if v.number == 1 {
                res = v
                break
            }
        }
        if node.number%2 == 0 && !visited[node.number/2] {
            v.count = node.count + 1
            v.number = node.number / 2
            newHistory := make([]int, len(node.history))
            copy(newHistory, node.history)
            newHistory = append(newHistory, v.number)
            v.history = newHistory

            visited[v.number] = true
            queue = append(queue, v)

            if v.number == 1 {
                res = v
                break
            }
        }
        if node.number-1 >= 1 && !visited[node.number-1] {
            v.count = node.count + 1
            v.number = node.number - 1
            newHistory := make([]int, len(node.history))
            copy(newHistory, node.history)
            newHistory = append(newHistory, v.number)
            v.history = newHistory

            visited[v.number] = true
            queue = append(queue, v)

            if v.number == 1 {
                res = v
                break
            }
        }
        queue = queue[1:]
    }

    return res.count, res.history
}

func main() {
    defer writer.Flush()

    n := 0
    fmt.Fscan(reader, &n)

    count, history := make1(n)
    fmt.Fprintln(writer, count)
    for i := 0; i < len(history); i++ {
        fmt.Fprintf(writer, "%d ", history[i])
    }
}
```

---

## ✅ 예시

### 입력
```
10
```

### 출력
```
3
10 9 3 1
```

---

## ✅ 시간 복잡도

- 최악의 경우 BFS는 O(N)
- 각 수에 대해 최대 한 번만 방문하므로 효율적임

---

## ✅ 결론

- BFS를 통해 `최단 경로 문제`를 해결할 수 있음
- 경로 추적을 위해 history 배열을 복사해서 다음 노드로 전달
- 방문 처리로 불필요한 탐색 방지

