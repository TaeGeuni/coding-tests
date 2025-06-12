# 무한 수열 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/6198)

## 🧮 문제 설명

무한 수열 A는 아래와 같이 정의된다.

- A₀ = 1  
- Aᵢ = A⌊i/P⌋ + A⌊i/Q⌋ (i ≥ 1)

정수 N, P, Q가 주어질 때, Aₙ의 값을 구하는 프로그램을 작성하시오.

- ⌊x⌋는 x를 넘지 않는 가장 큰 정수 (내림)이다.

---

## 📥 입력

- 첫째 줄: 세 개의 정수 N, P, Q (0 ≤ N ≤ 10¹², 2 ≤ P, Q ≤ 10⁹)

## 📤 출력

- Aₙ을 출력한다.

---

## ✅ 예제

### 입력 1
```
7 2 3
```

### 출력 1
```
7
```

### 입력 2
```
0 2 3
```

### 출력 2
```
1
```

### 입력 3
```
10000000 3 3
```

### 출력 3
```
32768
```

---

## 💡 해결 아이디어

- 문제의 핵심은 **중복되는 연산을 피하기 위한 메모이제이션**이다.
- A₀ = 1이고, Aₙ = Aₙ/P + Aₙ/Q 이므로, 동일한 `n`에 대해 중복 호출을 방지하기 위해 **map[int]int** 자료구조를 사용하여 값을 저장해두자.
- 재귀적으로 호출하며 계산된 값을 저장하고, 이후에 같은 인덱스를 다시 계산하지 않도록 한다.

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
    dp     = make(map[int]int)
)

func Sequence(n, p, q int) int {
    if n == 0 {
        return 1
    }
    if val, exists := dp[n]; exists {
        return val
    }
    dp[n] = Sequence(n/p, p, q) + Sequence(n/q, p, q)
    return dp[n]
}

func main() {
    defer writer.Flush()

    var n, p, q int
    fmt.Fscan(reader, &n, &p, &q)

    fmt.Fprintln(writer, Sequence(n, p, q))
}
```

---

## ⏱️ 시간 복잡도

- 최악의 경우에도 `dp`에 저장되는 서로 다른 인덱스는 O(log N) 개수만큼이며, 각 인덱스는 한 번만 계산됨.
- 따라서 전체 시간복잡도는 **O(log N)**.

---

## 🧠 핵심 포인트

- 문제에서 N의 범위가 최대 10¹²이므로, 단순 재귀로는 해결이 불가능하다.
- **동일한 하위 문제의 중복 계산을 방지하는 메모이제이션(Memoization)** 이 필수.
