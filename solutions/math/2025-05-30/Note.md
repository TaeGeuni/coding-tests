# 다이어트
👉🏻[문제 링크](https://www.acmicpc.net/problem/1484)

## 📝 문제 설명

성원이는 다이어트를 시도 중이다. 성원이는 정말 정말 무겁기 때문에, 저울이 부서졌다.  
성원이의 힘겨운 다이어트 시도를 보고만 있던 엔토피아는 성원이에게 새로운 저울을 선물해 주었다.  
성원이는 엔토피아가 선물해준 저울 위에 올라갔다.

> “안돼!!!! G 킬로그램이나 더 쪘어ㅠㅠ”

여기서 말하는 G 킬로그램은:

```
현재 몸무게² - 기억하고 있던 몸무게² = G
```

이 수식을 만족하는 현재 몸무게 `x`들을 오름차순으로 출력하자.

---

## 📌 입력

- 첫째 줄에 G가 주어진다. (1 ≤ G ≤ 100,000)

---

## 📌 출력

- 가능한 현재 몸무게들을 한 줄에 하나씩 오름차순으로 출력한다.
- 가능한 몸무게가 없을 경우 `-1` 출력.
- 몸무게는 자연수로만 떨어지는 경우만 유효하다.

---

## ✅ 예제 입력

```
15
```

## ✅ 예제 출력

```
4
8
```

---

## 💡 해결 아이디어

이 문제는 다음 조건을 만족하는 두 자연수 `x`, `y`를 찾는 문제다:

```
x^2 - y^2 = G
=> (x - y)(x + y) = G
```

이 수식에서 `x`, `y`는 자연수이므로, 다음 조건을 만족해야 한다:

- `x = sqrt(G + y^2)`
- `x`는 정수여야 함

즉, `y`를 1부터 G-1까지 탐색하면서, `x = sqrt(G + y^2)`가 정수인지 확인한다.

시간 복잡도는 O(√G)로 매우 효율적이다.

---

## ✅ Go 언어 구현

```go
package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
)

var (
    reader = bufio.NewReader(os.Stdin)
    writer = bufio.NewWriter(os.Stdout)
)

func Diet(g int) []int {
    res := make([]int, 0)

    for i := 1; i < g; i++ {
        num := g + (i * i)

        x := math.Sqrt(float64(num))

        if x == float64(int(x)) {
            res = append(res, int(x))
        }
    }

    if len(res) == 0 {
        res = append(res, -1)
    }

    return res
}

func main() {
    defer writer.Flush()

    var g int
    fmt.Fscan(reader, &g)

    weight := Diet(g)

    for i := 0; i < len(weight); i++ {
        fmt.Fprintln(writer, weight[i])
    }
}
```

---

## 🧠 핵심 포인트

- `(현재 몸무게)^2 - (과거 몸무게)^2 = G` 형태를 역으로 계산
- `(x^2 - y^2 = G)` → `(x = sqrt(G + y^2))` 로 역산
- 이 때 `x`는 자연수여야만 출력 대상이 됨
- 모든 가능한 `y`에 대해 계산하고, 유효한 `x`만 수집

---

## ✅ 시간 복잡도

- O(√G) 이하
- 최악의 경우에도 충분히 빠름