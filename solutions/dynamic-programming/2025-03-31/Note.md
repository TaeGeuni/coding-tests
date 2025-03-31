# 정수 삼각형 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/1932)

## 문제 설명

다음과 같은 정수 삼각형이 주어진다.

```
        7
      3   8
    8   1   0
  2   7   4   4
4   5   2   6   5
```

맨 위층 7부터 시작하여 아래층의 숫자 중 하나를 선택하여 내려올 때, 선택된 숫자의 합이 최대가 되는 경로를 구하는 프로그램을 작성하라. 아래층의 숫자는 현재 위치에서 대각선 왼쪽 또는 대각선 오른쪽만 선택할 수 있다.

### 입력
- 첫째 줄에 삼각형의 크기 `n`(1 ≤ n ≤ 500)이 주어진다.
- 둘째 줄부터 `n+1`번째 줄까지 정수 삼각형이 주어진다.
- 삼각형을 이루는 각 수는 0 이상 9999 이하의 정수이다.

### 출력
- 합이 최대가 되는 경로에 있는 수의 합을 출력한다.

### 예제 입력
```
5
7
3 8
8 1 0
2 7 4 4
4 5 2 6 5
```

### 예제 출력
```
30
```

---

## 해결 방법

이 문제는 **동적 프로그래밍(DP)** 을 활용하여 해결할 수 있다. 각 위치에서 선택 가능한 두 경로 중 더 큰 값을 선택하면서 삼각형을 아래층으로 내려가는 방식으로 최대 합을 계산한다.

### 접근 방법
1. `dp[i][j]`를 `(i, j)` 위치에서 최대 경로 합으로 정의한다.
2. `dp[i][j] = max(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]` 형태로 값을 갱신한다.
3. 마지막 행에서 최대값을 찾으면 정답이 된다.

---

## 코드 구현 (Go)

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

func Maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func DP(triangle [][]int) int {
	res := 0
	dpSli := make([]int, 0)
	dpSli = append(dpSli, triangle[0][0])

	for i := 1; i < len(triangle); i++ {
		data := make([]int, 0)
		for j := 0; j < i+1; j++ {
			if j == 0 {
				data = append(data, (dpSli[0] + triangle[i][0]))
			} else if j == i {
				data = append(data, (dpSli[len(dpSli)-1] + triangle[i][j]))
			} else {
				data = append(data, Maximum((dpSli[j-1]+triangle[i][j]), (dpSli[j]+triangle[i][j])))
			}
		}
		dpSli = data
	}

	for i := 0; i < len(dpSli); i++ {
		if res < dpSli[i] {
			res = dpSli[i]
		}
	}

	return res
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	triangle := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			num := 0
			fmt.Fscan(reader, &num)
			triangle[i] = append(triangle[i], num)
		}
	}

	fmt.Fprintln(writer, DP(triangle))
}
```

---

## 시간 복잡도 분석

- 이 알고리즘은 `n` 단계의 삼각형을 한 층씩 내려가면서 최적의 합을 계산한다.
- 각 단계에서 `O(n)` 연산이 이루어지므로 전체 시간 복잡도는 **O(n²)** 이다.
- `n`의 최대값이 500이므로, 최대 연산 횟수는 `500 × 500 = 250000` 정도로 충분히 계산이 가능하다.

## 결론

이 문제는 **DP(동적 프로그래밍)** 를 활용하여 최적의 경로를 찾아가는 문제이다. 위에서부터 내려오면서 최댓값을 갱신하며 저장하는 방식으로 효율적으로 해결할 수 있다.

