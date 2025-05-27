# 알약
👉🏻[문제 링크](https://www.acmicpc.net/problem/4811)

## 문제 설명

박종수 할아버지는 매일 반 조각의 약을 먹습니다. 병에 N개의 약이 있을 때, 총 2N일 동안 약을 먹게 되고, 약을 꺼낸 방식에 따라 손녀에게 문자 `W` 또는 `H`를 보냅니다.

- `W`: 하나의 약을 꺼내 반으로 나누고 절반만 먹음 (전체 조각 꺼냄)
- `H`: 반 조각을 꺼내서 모두 먹음

문자열의 길이는 항상 2N이고, 가능한 서로 다른 문자열의 개수를 구하는 문제입니다.

이 문제는 **카탈란 수**로 해결할 수 있습니다.

---

## 카탈란 수 (Catalan Number)

카탈란 수는 다음과 같이 정의됩니다:

\[
C_n = \frac{1}{n + 1} \binom{2n}{n}
\]

또는 다음과 같이 계산할 수 있습니다:

\[
C_n = \frac{(2n)!}{(n+1)! \cdot n!}
\]

이 문제에서 `n`개의 전체 조각을 사용하여 만들 수 있는 올바른 문자열의 개수는 \( C_n \)과 동일합니다.

---

## 해결 방법 (Go 코드)

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

func catalanBinomial(n int) int {
	res := 1

	for i := 0; i < n; i++ {
		res = (res * (2*n - i)) / (i + 1)
	}

	return res / (n + 1)
}

func main() {
	defer writer.Flush()

	var n int
	for {
		fmt.Fscan(reader, &n)
		if n == 0 {
			break
		}
		fmt.Fprintln(writer, catalanBinomial(n))
	}
}
```

---

## 예제 입력 및 출력

**입력**:
```
6
1
4
2
3
30
0
```

**출력**:
```
132
1
14
2
5
3814986502092304
```

---

## 요약

이 문제는 괄호 문자열의 개수 또는 스택을 사용하는 문제들과 유사하며, **카탈란 수를 계산하는 공식을 활용**하여 빠르게 해결할 수 있습니다. N의 범위가 작기 때문에 정수형 연산으로 해결이 가능합니다.