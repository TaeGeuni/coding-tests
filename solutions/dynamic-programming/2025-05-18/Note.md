# 카드 구매하기
👉🏻[문제 링크](https://www.acmicpc.net/problem/11052)

## 문제 설명

요즘 민규네 동네에서는 스타트링크에서 만든 PS카드를 모으는 것이 유행이다.

PS카드는 PS(Problem Solving)분야에서 유명한 사람들의 아이디와 얼굴이 적혀있는 카드이다. 각각의 카드에는 등급을 나타내는 색이 칠해져 있고, 다음과 같이 8가지가 있다:

- 전설카드
- 레드카드
- 오렌지카드
- 퍼플카드
- 블루카드
- 청록카드
- 그린카드
- 그레이카드

카드는 카드팩의 형태로만 구매할 수 있고, 카드팩의 종류는 카드 1개가 포함된 카드팩, 카드 2개가 포함된 카드팩, ..., 카드 N개가 포함된 카드팩과 같이 총 N가지가 존재한다.

민규는 카드의 개수가 적은 팩이더라도 가격이 비싸면 높은 등급의 카드가 많이 들어있을 것이라는 미신을 믿고 있다. 따라서, 민규는 돈을 **최대한 많이 지불해서 카드 N개 구매**하려고 한다.

카드 팩의 가격이 주어졌을 때, N개의 카드를 구매하기 위해 민규가 지불해야 하는 금액의 **최댓값**을 구하는 프로그램을 작성하시오.

> N개보다 많은 개수의 카드를 산 다음, 나머지를 버리는 것은 불가능하다.

---

## 입력

- 첫째 줄에 민규가 구매하려고 하는 카드의 개수 N이 주어진다. (1 ≤ N ≤ 1,000)
- 둘째 줄에는 P₁부터 Pₙ까지 카드팩의 가격이 주어진다. (1 ≤ Pᵢ ≤ 10,000)

## 출력

- 민규가 카드 N개를 갖기 위해 지불해야 하는 금액의 **최댓값**을 출력한다.

---

## 예제 입력 / 출력

### 예제 입력 1
```
4
1 5 6 7
```
### 예제 출력 1
```
10
```

---

## 해결 아이디어

이 문제는 **DP(동적 계획법)** 을 사용하여 해결할 수 있다.

- `dp[i]` = 카드 i개를 구매하는데 지불할 수 있는 최대 금액
- 점화식:
    ```
    dp[i] = max(dp[j] + dp[i - j]) for all 1 <= j <= i
    ```

- 단, 카드팩 가격이 `p[i]`로 주어졌기 때문에, `dp[i] = max(dp[j] + dp[i - j], p[i])` 도 고려해줘야 함.

---

## Go 코드 구현

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

func maximum(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func DP(n int, p []int) int {
    dp := make([]int, 0)
    dp = append(dp, p[0])

    for i := 1; i < n; i++ {
        left := 0
        right := i - 1
        num := 0
        for left <= right {
            if num < (dp[left] + dp[right]) {
                num = (dp[left] + dp[right])
            }
            left++
            right--
        }
        dp = append(dp, maximum(num, p[i]))
    }

    return dp[n-1]
}

func main() {
    defer writer.Flush()

    var n int
    fmt.Fscan(reader, &n)
    p := make([]int, n)

    for i := 0; i < n; i++ {
        fmt.Fscan(reader, &p[i])
    }

    fmt.Fprintln(writer, DP(n, p))
}
```

---

## 시간복잡도

- 이중 반복문을 사용하므로 **O(N²)** 이다. (최대 N은 1,000이므로 허용됨)

---

## 핵심 포인트

- 완전 탐색처럼 보이지만, 중복 계산을 줄이기 위해 DP 사용
- `i`개의 카드를 만들기 위해 `j`개 + `(i-j)`개로 나눠 가능한 최대값 갱신
- 카드팩 그대로 사용하는 경우도 고려해야 하므로 `p[i]`와도 비교

