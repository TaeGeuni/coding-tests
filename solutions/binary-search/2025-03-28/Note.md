# 예산 배정 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/2343)

## 문제 설명

국가는 여러 지방의 예산 요청을 심사하여 국가의 예산을 분배합니다. 국가 예산의 총액은 미리 정해져 있어, 모든 요청을 배정할 수 없는 경우에는 특정한 상한액을 설정하여 예산을 배정합니다. 상한액 이하의 요청은 그대로 배정되고, 상한액을 초과하는 요청은 상한액만 배정됩니다. 목표는 정해진 총액 이하에서 가능한 한 최대의 총 예산을 배정하는 상한액을 찾는 것입니다.

## 입력 형식

1. 첫째 줄에 지방의 수를 나타내는 정수 `N` (3 ≤ N ≤ 10,000)이 주어집니다.
2. 둘째 줄에 각 지방의 예산 요청을 나타내는 `N`개의 정수 (1 ≤ 요청 금액 ≤ 100,000)가 공백으로 구분되어 주어집니다.
3. 셋째 줄에 총 예산을 나타내는 정수 `M` (N ≤ M ≤ 1,000,000,000)이 주어집니다.

## 출력 형식

1. 상한액을 출력합니다.

### 예제 입력 1

```
4
120 110 140 150
485
```

### 예제 출력 1

```
127
```

### 예제 입력 2

```
5
70 80 30 40 100
450
```

### 예제 출력 2

```
100
```

---

## 풀이 과정

이 문제는 **이진 탐색**을 사용하여 상한액을 효율적으로 찾을 수 있습니다.

1. **초기 설정**:
   - 요청 중 최대값을 `right`로 설정합니다.
   - 최소 상한액 `0`을 `left`로 설정합니다.
   - 상한액 후보를 저장할 `res`를 초기화합니다.

2. **이진 탐색**:
   - `mid` 값을 계산하고, 이 값을 상한액으로 했을 때 모든 요청을 처리한 후 예산 총액 `M`을 초과하는지 확인합니다.
   - 예산이 남으면 `left = mid + 1`로 상한액을 늘립니다.
   - 예산이 부족하면 `right = mid - 1`로 상한액을 줄입니다.

3. **최종 출력**:
   - 가능한 상한액 중 최대값을 출력합니다.

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

func BinarySearch(biggest, total int, requestBudget []int) int {
	res := 0

	left, right := 0, biggest

	for left <= right {
		mid := (left + right) / 2
		t := total

		for i := 0; i < len(requestBudget); i++ {
			if requestBudget[i] > mid {
				t -= mid
			} else {
				t -= requestBudget[i]
			}
		}

		if t >= 0 {
			left = mid + 1
			res = mid
		} else {
			right = mid - 1
		}
	}

	return res
}

func main() {
	defer writer.Flush()

	var n, biggest, total int

	fmt.Fscan(reader, &n)
	requestBudget := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &requestBudget[i])
		if biggest < requestBudget[i] {
			biggest = requestBudget[i]
		}
	}
	fmt.Fscan(reader, &total)

	fmt.Fprintln(writer, BinarySearch(biggest, total, requestBudget))
}
```

---

## 코드 설명

1. **입력 처리**:
   - 지방의 수 `n`과 각 지방의 요청 예산을 입력받습니다.
   - 요청 예산 중 최대값 `biggest`를 계산합니다.
   - 총 예산 `total`을 입력받습니다.

2. **이진 탐색 함수**:
   - `BinarySearch` 함수는 상한액을 계산합니다.
   - `mid`를 기준으로 모든 요청을 처리하고, 남은 예산 `t`를 확인합니다.
   - 남은 예산이 0 이상이면 상한액을 늘리고, 아니면 줄입니다.

3. **결과 출력**:
   - 상한액 중 최대값을 출력합니다.

---

## 시간 복잡도

1. 이진 탐색: O(log(max(requestBudget)))
2. 각 탐색에서 요청 배열을 순회: O(N)

총 시간 복잡도: O(N \* log(max(requestBudget)))

---

## 주요 포인트

- 이진 탐색을 활용하여 상한액을 효율적으로 찾을 수 있습니다.
- 입력 크기가 크더라도 시간 복잡도가 낮아 높은 성능을 보장합니다.

