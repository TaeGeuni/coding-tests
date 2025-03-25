# RGB 거리 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/1149)

## 문제 설명
RGB거리에는 집이 N개 있다. 집은 빨강, 초록, 파랑 중 하나의 색으로 칠해야 하며, 아래 규칙을 만족해야 한다.

1. 1번 집의 색은 2번 집의 색과 같지 않아야 한다.
2. N번 집의 색은 N-1번 집의 색과 같지 않아야 한다.
3. i(2 ≤ i ≤ N-1)번 집의 색은 i-1번, i+1번 집의 색과 같지 않아야 한다.

각 집을 칠하는 비용이 주어졌을 때, 모든 집을 칠하는 비용의 최솟값을 구하는 문제이다.

## 입력
첫째 줄에 집의 수 N(2 ≤ N ≤ 1,000)이 주어진다.
둘째 줄부터 N개의 줄에는 각 집을 빨강, 초록, 파랑으로 칠하는 비용이 주어진다. 비용은 1,000 이하의 자연수이다.

## 출력
모든 집을 칠하는 비용의 최솟값을 출력한다.

## 예제 입력 및 출력
### 예제 입력 1
```
3
26 40 83
49 60 57
13 89 99
```
### 예제 출력 1
```
96
```

### 예제 입력 2
```
3
1 100 100
100 1 100
100 100 1
```
### 예제 출력 2
```
3
```

## 풀이 과정
이 문제는 동적 계획법(DP)으로 해결할 수 있다.

1. `cost[i][0]`, `cost[i][1]`, `cost[i][2]`는 각각 i번째 집을 빨강, 초록, 파랑으로 칠하는 비용을 의미한다.
2. `dp[i][0]`, `dp[i][1]`, `dp[i][2]`는 i번째 집을 해당 색으로 칠했을 때의 최소 비용을 저장한다.
3. 점화식은 다음과 같다:
   - `dp[i][0] = min(dp[i-1][1], dp[i-1][2]) + cost[i][0]`
   - `dp[i][1] = min(dp[i-1][0], dp[i-1][2]) + cost[i][1]`
   - `dp[i][2] = min(dp[i-1][0], dp[i-1][1]) + cost[i][2]`
4. 마지막 집에서 최소값을 출력하면 된다.

## 코드 구현 (Golang)
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func Minimum(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	defer writer.Flush()

	var num int
	fmt.Fscan(reader, &num)
	cost := make([][3]int, num)

	for i := 0; i < num; i++ {
		var r, g, b int
		fmt.Fscan(reader, &r, &g, &b)
		cost[i][0], cost[i][1], cost[i][2] = r, g, b
	}
	dp := make([][3]int, num)
	dp[0][0], dp[0][1], dp[0][2] = cost[0][0], cost[0][1], cost[0][2]

	for i := 1; i < num; i++ {
		dp[i][0] = Minimum(dp[i-1][1], dp[i-1][2]) + cost[i][0]
		dp[i][1] = Minimum(dp[i-1][0], dp[i-1][2]) + cost[i][1]
		dp[i][2] = Minimum(dp[i-1][0], dp[i-1][1]) + cost[i][2]
	}
	res := Minimum(dp[num-1][0], dp[num-1][1])
	res = Minimum(res, dp[num-1][2])

	fmt.Fprintln(writer, res)
}
```

## 시간 복잡도
이 코드는 O(N)으로 동작하며, 각 집에 대해 3가지 색상 선택에 대한 최소 비용을 계산하는 방식이다. `N ≤ 1,000`이므로 충분히 빠르게 실행된다.

