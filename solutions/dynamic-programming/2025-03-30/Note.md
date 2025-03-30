# 연속합
👉🏻[문제 링크](https://www.acmicpc.net/problem/1912)

## 문제 설명

**n**개의 정수로 이루어진 임의의 수열이 주어진다. 우리는 이 중 연속된 몇 개의 수를 선택해서 구할 수 있는 합 중 가장 큰 합을 구하려고 한다. 단, 수는 한 개 이상 선택해야 한다.

예를 들어서 다음과 같은 수열이 주어진 경우를 살펴보자:

```
10, -4, 3, 1, 5, 6, -35, 12, 21, -1
```

여기서 정답은 **12 + 21**인 **33**이 정답이 된다.

---

## 입력

- 첫째 줄에 정수 **n(1 ≤ n ≤ 100,000)**이 주어진다.
- 둘째 줄에는 **n**개의 정수로 이루어진 수열이 주어진다. 수는 **-1,000 ≤ x ≤ 1,000**이다.

---

## 출력

- 첫째 줄에 답을 출력한다.

---

## 예제 입력과 출력

### 예제 1

**입력:**
```
10
10 -4 3 1 5 6 -35 12 21 -1
```

**출력:**
```
33
```

### 예제 2

**입력:**
```
10
2 1 -4 3 4 -4 6 5 -5 1
```

**출력:**
```
14
```

### 예제 3

**입력:**
```
5
-1 -2 -3 -4 -5
```

**출력:**
```
-1
```

---

## 풀이 과정

이 문제는 연속 부분 수열의 합 중 최대값을 구하는 문제로, **카데인 알고리즘(Kadane's Algorithm)**을 이용하여 효율적으로 해결할 수 있다.

---

### 구현 코드

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

func DP(sequence []int) int {
	max := 0
	sum := 0

	max = sequence[0]
	if sequence[0] > 0 {
		sum = sequence[0]
	} else {
		sum = 0
	}

	for i := 1; i < len(sequence); i++ {
		max = Maximum(max, (sum + sequence[i]))
		sum += sequence[i]
		if sum < 0 {
			sum = 0
		}
	}

	return max
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	sequence := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &sequence[i])
	}
	fmt.Fprintln(writer, DP(sequence))
}
```

---

## 풀이 설명

1. **입력 처리**
   - `n`개의 정수를 입력받고, 이를 슬라이스로 저장한다.

2. **카데인 알고리즘 구현**
   - 현재 연속된 부분 수열의 합(`sum`)과 최대값(`max`)을 관리한다.
   - 각 단계에서 새로운 값과 이전의 합을 더한 결과 중 더 큰 값을 `max`에 업데이트한다.
   - 합(`sum`)이 음수가 되면 0으로 초기화한다.

3. **결과 출력**
   - 최대값을 출력한다.

---

## 시간 복잡도

- **O(n)**: 한 번의 순회로 최대 합을 구하므로 효율적이다.

## 공간 복잡도

- **O(1)**: 입력 크기에 관계없이 일정한 공간만 사용한다.

---

## 참고 사항

- 이 풀이에서는 음수만 있는 경우에도 최대값이 제대로 계산되도록 설계되었다.
- 입력의 크기가 최대 100,000이므로, O(n) 시간 복잡도로 충분히 처리 가능하다.
