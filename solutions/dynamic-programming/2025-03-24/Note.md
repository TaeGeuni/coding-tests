# 계단 오르기 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/2579)

## 문제 설명
계단 오르기 게임은 계단 아래 시작점부터 계단 꼭대기에 위치한 도착점까지 가는 게임입니다. 각각의 계단에는 일정한 점수가 쓰여 있는데, 계단을 밟으면 해당 점수를 얻게 됩니다. 이때 다음 규칙을 만족해야 합니다:

1. 계단은 한 번에 한 계단씩 또는 두 계단씩 오를 수 있습니다.
2. **연속된 세 개의 계단을 모두 밟을 수 없습니다.** (단, 시작점은 계단에 포함되지 않습니다.)
3. 마지막 도착 계단은 반드시 밟아야 합니다.

각 계단에 쓰여 있는 점수가 주어질 때, 게임에서 얻을 수 있는 총 점수의 최댓값을 구하세요.

---

## 입력 형식
- 첫째 줄: 계단의 개수 `n` (1 ≤ n ≤ 300)
- 둘째 줄부터 `n`개의 줄에 각 계단에 쓰여 있는 점수 (1 ≤ 점수 ≤ 10,000)

## 출력 형식
- 계단 오르기 게임에서 얻을 수 있는 총 점수의 최댓값을 출력합니다.

---

## 풀이
### 핵심 아이디어
동적 계획법(Dynamic Programming, DP)을 사용하여 문제를 해결할 수 있습니다. 

`dp[i]`를 **i번째 계단까지 올라왔을 때 얻을 수 있는 최대 점수**라고 정의합니다. 이때, `dp[i]`를 계산하기 위한 점화식은 다음과 같습니다:

1. `dp[i-2] + stairs[i]`: i-2에서 두 계단 올라온 경우
2. `dp[i-3] + stairs[i-1] + stairs[i]`: i-3에서 올라와서 i-1과 i를 연속으로 밟은 경우

이 두 값 중 큰 값을 선택하면 됩니다.

### 초기값 설정
- `dp[1] = stairs[1]`: 첫 번째 계단만 밟은 경우
- `dp[2] = stairs[1] + stairs[2]`: 첫 번째 계단과 두 번째 계단을 밟은 경우

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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	defer writer.Flush()

	var num int
	fmt.Fscan(reader, &num)

	stairs := make([]int, num+1)

	for i := 1; i < num+1; i++ {
		fmt.Fscan(reader, &stairs[i])
	}

	dp := make([]int, num+1)

	if num >= 1 {
		dp[1] = stairs[1]
	}
	if num >= 2 {
		dp[2] = stairs[1] + stairs[2]
	}

	for i := 3; i < num+1; i++ {
		dp[i] = max((dp[i-2] + stairs[i]), (dp[i-3] + stairs[i-1] + stairs[i]))
	}

	fmt.Fprintln(writer, dp[num])
}
```

---

## 입출력 예
### 입력
```
6
10
20
15
25
10
20
```

### 출력
```
75
```

### 설명
1. 첫 번째 계단(10점)을 밟음
2. 두 번째 계단(20점)을 밟음
3. 네 번째 계단(25점)을 밟음
4. 여섯 번째 계단(20점)을 밟음

총합: `10 + 20 + 25 + 20 = 75`

---

## 시간 복잡도
- 계단의 개수 `n`에 대해 한 번의 반복문으로 계산하므로 **O(n)**입니다.

## 공간 복잡도
- 동적 배열 `dp`와 입력 배열 `stairs`를 사용하므로 **O(n)**입니다.
