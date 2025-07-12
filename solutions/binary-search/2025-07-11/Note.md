# 🍕 피자 나누기 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/28081)

## 📘 문제 설명

용모는 가로 길이 `W`, 세로 길이 `H`인 직사각형 피자를 주문했습니다. 피자는 가로 및 세로 방향으로 커팅 선이 주어지며, 각 커팅은 피자를 여러 조각으로 분할합니다.

각 조각의 크기는 `가로 길이 * 세로 길이`입니다. 이때, 조각의 크기가 `K` 이하일 경우에만 CTP 부원에게 나눠줄 수 있습니다.

> 피자의 정보를 바탕으로 부원에게 나눠줄 수 있는 피자 조각의 **개수**를 출력하세요.

---

## 🔢 입력 형식

```
W H K
N
y1 y2 ... yN
M
x1 x2 ... xM
```

- `W`, `H`: 피자의 가로, 세로 길이 (2 ≤ W, H ≤ 10^9)
- `K`: 조각당 최대 크기 (1 ≤ K ≤ W * H)
- `N`: 가로 커팅 수 (1 ≤ N ≤ min(H-1, 100000))
- `y1 ... yN`: 가로 커팅 위치 (오름차순, 1 ≤ yi ≤ H - 1)
- `M`: 세로 커팅 수 (1 ≤ M ≤ min(W-1, 100000))
- `x1 ... xM`: 세로 커팅 위치 (오름차순, 1 ≤ xi ≤ W - 1)

---

## 🧾 출력 형식

- 부원에게 나눠줄 수 있는 피자 조각의 **개수**를 출력한다.

---

## 🧠 예제 입력

```
7 5 6
1
2
2
1 5
```

## ✅ 예제 출력

```
4
```

---

## ✅ 알고리즘 풀이

1. 세로(x축), 가로(y축) 커팅 위치들을 기준으로 조각의 **가로/세로 길이 배열**을 만든다.
2. 모든 가능한 조각은 `가로 * 세로`의 형태로 계산될 수 있다.
3. 정렬된 세로/가로 길이 배열을 이용해 **이분 탐색**을 통해 `높이 * 너비 <= K` 조건을 만족하는 조각 개수를 센다.

---

## 💻 사용한 코드 (Go)

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
)

var (
    reader = bufio.NewReader(os.Stdin)
    writer = bufio.NewWriter(os.Stdout)
)

func distributePizza(k int, heights, widths []int) int {
    res := 0
    for _, w := range widths {
        if w == 0 {
            continue
        }
        target := k / w
        num := sort.Search(len(heights), func(j int) bool {
            return heights[j] > target
        })
        res += num
    }
    return res
}

func main() {
    defer writer.Flush()

    var w, h, k int
    fmt.Fscan(reader, &w, &h, &k)

    var n, m int
    fmt.Fscan(reader, &n)
    ys := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Fscan(reader, &ys[i])
    }
    heights := make([]int, 0, n+1)
    prev := 0
    for i := 0; i < n; i++ {
        heights = append(heights, ys[i]-prev)
        prev = ys[i]
    }
    heights = append(heights, h-prev)
    sort.Ints(heights)

    fmt.Fscan(reader, &m)
    xs := make([]int, m)
    for i := 0; i < m; i++ {
        fmt.Fscan(reader, &xs[i])
    }
    widths := make([]int, 0, m+1)
    prev = 0
    for i := 0; i < m; i++ {
        widths = append(widths, xs[i]-prev)
        prev = xs[i]
    }
    widths = append(widths, w-prev)

    fmt.Fprintln(writer, distributePizza(k, heights, widths))
}
```

---

## 📝 시간복잡도

- `O(N log N + M log M)` — 길이 배열 정렬 및 이분 탐색.
