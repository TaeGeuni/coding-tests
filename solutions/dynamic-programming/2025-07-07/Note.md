# 구간 합 구하기 - 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/11660)

## 🧩 문제 설명

N×N개의 수가 N×N 크기의 표에 채워져 있다. (x1, y1)부터 (x2, y2)까지 합을 구하는 프로그램을 작성하시오.

표에 채워져 있는 수와 합을 구하는 연산이 주어졌을 때, 이를 처리하는 프로그램을 작성하시오.

### 입력
- 첫째 줄에 표의 크기 N과 합을 구해야 하는 횟수 M이 주어진다. (1 ≤ N ≤ 1024, 1 ≤ M ≤ 100,000)
- 둘째 줄부터 N개의 줄에는 표에 채워져 있는 수가 1행부터 차례대로 주어진다.
- 다음 M개의 줄에는 네 개의 정수 x1, y1, x2, y2 가 주어진다.

### 출력
- 총 M줄에 걸쳐 (x1, y1)부터 (x2, y2)까지 합을 구해 출력한다.

---

## ✅ 예제 입력 1

```
4 3
1 2 3 4
2 3 4 5
3 4 5 6
4 5 6 7
2 2 3 4
3 4 3 4
1 1 4 4
```

## ✅ 예제 출력 1

```
27
6
64
```

---

## 💡 해결 아이디어

- 누적합을 활용하여 `(x1, y1)`부터 `(x2, y2)`까지의 합을 `O(1)` 시간에 구할 수 있도록 한다.
- 누적합 테이블 `S`를 다음과 같이 정의한다:

```
S[i][j] = (1,1) ~ (i,j)까지의 직사각형 영역의 총합
```

- 이를 이용해 구간 합은 다음과 같이 계산할 수 있다:

```
sum(x1, y1, x2, y2) = S[x2][y2] - S[x2][y1-1] - S[x1-1][y2] + S[x1-1][y1-1]
```

---

## ✅ Go 코드

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

func sumOfSections(x1, y1, x2, y2 int, table [][]int) int {
	res := table[x2][y2] - table[x2][y1-1] - table[x1-1][y2] + table[x1-1][y1-1]
	return res
}

func main() {
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	table := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		table[i] = make([]int, n+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			fmt.Fscan(reader, &table[i][j])
			table[i][j] += table[i-1][j] + table[i][j-1] - table[i-1][j-1]
		}
	}

	res := make([]int, m)
	for i := 0; i < m; i++ {
		x1, y1, x2, y2 := 0, 0, 0, 0
		fmt.Fscan(reader, &x1, &y1, &x2, &y2)
		res[i] = sumOfSections(x1, y1, x2, y2, table)
	}

	for i := 0; i < m; i++ {
		fmt.Fprintln(writer, res[i])
	}
}
```

---

## 🧠 시간 복잡도

- 입력 데이터 누적합 테이블 구성: `O(N^2)`
- 각 질의 처리 시간: `O(1)`
- 전체 시간 복잡도: `O(N^2 + M)` → 매우 빠른 풀이

---

## 🏁 핵심 요약

- 2차원 누적합을 활용하여 구간 합을 빠르게 구할 수 있다.
- 메모리를 넉넉히 사용해서 계산 속도를 향상시킨 전형적인 누적합 문제.

