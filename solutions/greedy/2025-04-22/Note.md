# 동전 뒤집기 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/1455)

## 문제 설명
세준이는 N×M 크기의 동전을 모두 앞면(0)으로 만들려고 합니다. 동전을 뒤집는 규칙은 다음과 같습니다:
- 특정 (a, b) 칸을 선택하면, (i, j) (1 ≤ i ≤ a, 1 ≤ j ≤ b)를 만족하는 영역 내의 동전이 모두 뒤집힙니다.
- 동전의 앞면은 0, 뒷면은 1로 표시됩니다.

목표는 모든 동전을 앞면으로 만들기 위해 필요한 최소 뒤집기 횟수를 계산하는 것입니다.

## 입력
1. 첫째 줄에 세로 크기 `N`과 가로 크기 `M`이 주어집니다. (`1 ≤ N, M ≤ 100`)
2. 둘째 줄부터 N개의 줄에 M개의 동전 상태가 주어집니다.
   - 동전 상태는 `0`(앞면)과 `1`(뒷면)로 이루어진 문자열입니다.

## 출력
- 첫째 줄에 동전을 모두 앞면으로 만들기 위해 필요한 최소 뒤집기 횟수를 출력합니다.

## 문제 풀이

### 알고리즘 설명
1. **뒤에서부터 탐색**:
   - 동전 배열의 오른쪽 아래에서부터 왼쪽 위로 탐색하며, 뒷면(`1`)이 발견되면 해당 칸을 기준으로 영역을 뒤집습니다.
2. **뒤집기 구현**:
   - (x, y) 칸에서 뒤집을 경우, (0, 0)부터 (x, y)까지의 모든 동전 상태를 반전시킵니다.
3. **최소 횟수 계산**:
   - 뒤집기를 수행할 때마다 횟수를 증가시키고, 모든 동전이 앞면이 될 때까지 반복합니다.

### 코드 구현
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

func TurnOver(n, m int, coin [][]byte) int {
	res := 0
	allZero := false

	for !allZero {
		x, y := 0, 0
		find := false
		for i := n - 1; i >= 0; i-- {
			for j := m - 1; j >= 0; j-- {
				if coin[i][j] == '1' {
					x, y = i, j
					find = true
				}
				if find {
					break
				}
			}
			if find {
				break
			}
		}

		if find {
			res++
			for i := 0; i <= x; i++ {
				for j := 0; j <= y; j++ {
					if coin[i][j] == '1' {
						coin[i][j] = '0'
					} else {
						coin[i][j] = '1'
					}
				}
			}
		} else {
			allZero = true
		}
	}

	return res
}

func main() {
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)
	coin := make([][]byte, n)

	for i := 0; i < n; i++ {
		input := ""
		fmt.Fscan(reader, &input)
		for j := 0; j < m; j++ {
			coin[i] = append(coin[i], input[j])
		}
	}
	res := TurnOver(n, m, coin)
	fmt.Fprintln(writer, res)
}
```

### 동작 설명
1. `TurnOver` 함수:
   - 동전 배열을 탐색하며 뒷면(`1`)이 발견되면 뒤집기 연산을 수행합니다.
   - 탐색은 오른쪽 아래에서 시작하여 모든 동전이 앞면(`0`)이 될 때까지 반복합니다.
2. `main` 함수:
   - 입력 데이터를 처리하고, 결과를 출력합니다.

### 입출력 예시
#### 입력
```
2 2
00
01
```
#### 출력
```
4
```

### 풀이의 핵심
- 문제의 핵심은 뒤집기 연산이 영향을 미치는 범위를 효율적으로 계산하는 것입니다.
- 오른쪽 아래부터 탐색하여 최소 횟수를 계산할 수 있습니다.

---

이 풀이 방법은 시간 복잡도가 O(N × M)으로, 입력 크기가 작아 효율적으로 동작합니다.
