# 포도주 시식 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/2156)

## 문제 설명

효주는 포도주 시식회에 참가했습니다. 테이블 위에 일렬로 놓인 `n`개의 포도주 잔이 있으며, 다음과 같은 규칙이 있습니다:

1. 포도주 잔을 선택하면 그 잔에 들어있는 포도주는 **모두 마셔야** 한다.
2. **연속으로 3잔은 마실 수 없다**.

효주가 가장 많은 양의 포도주를 마실 수 있도록 하는 프로그램을 작성하는 것이 목표입니다.

---

## 입력
- 첫째 줄에 포도주 잔의 개수 `n`이 주어진다. (1 ≤ n ≤ 10,000)
- 둘째 줄부터 `n`개의 줄에 각 포도주 잔에 들어있는 포도주의 양이 순서대로 주어진다. (0 ≤ 포도주의 양 ≤ 1,000)

## 출력
- 최대로 마실 수 있는 포도주의 양을 출력한다.

---

## 접근 방식

### 핵심 아이디어 (DP)
`dp[i]`를 `i번째 포도주 잔까지 고려했을 때 최대로 마실 수 있는 포도주 양`이라고 정의합니다.

다음 3가지 선택지 중 최대값을 고릅니다:

1. `i번째` 잔과 `i-1번째` 잔을 마시고, `i-2`는 쉬기 → `dp[i-3] + wine[i-1] + wine[i]`
2. `i번째` 잔만 마시고, `i-1`은 쉬기 → `dp[i-2] + wine[i]`
3. `i번째` 잔을 마시지 않음 → `dp[i-1]`

### 점화식
```
dp[i] = max(dp[i-1], dp[i-2] + wine[i], dp[i-3] + wine[i-1] + wine[i])
```

### 초기값 설정
```
dp[0] = 0
n == 1: dp[1] = wine[0]
n == 2: dp[2] = wine[0] + wine[1]
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

func maximum(a, b, c int) int {
	var d int
	if a > b {
		d = a
	} else {
		d = b
	}
	if d > c {
		return d
	}
	return c
}

func DP(n int, wine []int) int {
	if n == 1 {
		return wine[0]
	} else if n == 2 {
		return wine[0] + wine[1]
	}
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 0, wine[0], wine[0]+wine[1]

	for i := 2; i < n; i++ {
		dp[i+1] = maximum(dp[i-2]+wine[i-1]+wine[i], dp[i-1]+wine[i], dp[i])
	}

	return dp[n]
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	wine := make([]int, n+1)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &wine[i])
	}

	res := DP(n, wine)
	fmt.Fprintln(writer, res)
}
```

---

## 예제 입력
```
6
6
10
13
9
8
1
```

## 예제 출력
```
33
```

## 설명
- 1, 2, 4, 5번째 포도주 잔을 마신 경우: `6 + 10 + 9 + 8 = 33`
- 3잔 연속 금지 조건을 만족하면서 가능한 최대값입니다.

---

## 결론
이 문제는 대표적인 DP 문제로, 연속 선택에 대한 제약 조건이 있는 최적화 문제입니다. 점화식을 정확히 세우고, 상태를 나누어 생각하는 것이 핵심입니다.
