# 나무 자르기 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/2805)

## 문제 설명
상근이는 필요한 만큼의 나무를 얻기 위해 절단기를 사용하여 나무를 자릅니다. 절단기의 높이를 설정하면, 해당 높이 위의 부분만 잘려 나가고, 설정된 높이 이하의 나무는 그대로 유지됩니다. 상근이는 최소한 M 미터의 나무를 가져가야 하며, 이 조건을 만족하는 절단기의 최대 높이를 구하는 프로그램을 작성해야 합니다.

## 입력
- 첫째 줄: 나무의 수 `N` (1 ≤ N ≤ 1,000,000)과 필요한 나무의 길이 `M` (1 ≤ M ≤ 2,000,000,000)
- 둘째 줄: 각 나무의 높이 (각 높이는 1,000,000,000 이하의 양의 정수 또는 0)
- 나무의 높이 합은 항상 M보다 크거나 같음

## 출력
- 절단기에 설정할 수 있는 최대 높이

## 예제
### 입력 1
```
4 7
20 15 10 17
```
### 출력 1
```
15
```
### 입력 2
```
5 20
4 42 40 26 46
```
### 출력 2
```
36
```

---

## 해결 방법
이 문제는 **이진 탐색(Binary Search)** 을 활용하여 해결할 수 있습니다.

1. 절단기의 높이를 `left = 0`, `right = max(trees)` 범위 내에서 이진 탐색을 수행합니다.
2. `mid = (left + right) / 2`로 절단기의 높이를 설정하고, 나무를 잘라서 얻을 수 있는 나무 길이를 계산합니다.
3. 잘라낸 나무의 길이가 `M` 이상이면, 절단기의 높이를 더 높일 수 있으므로 `left = mid + 1`로 변경합니다.
4. 잘라낸 나무의 길이가 `M`보다 작으면, `right = mid - 1`로 변경하여 높이를 낮춥니다.
5. 이진 탐색이 종료될 때까지 위 과정을 반복하며, 최적의 절단기 높이를 찾습니다.

---

## Go 언어 풀이 코드
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n, m, left, mid, right, res int

	fmt.Fscan(reader, &n, &m)
	trees := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &trees[i])
		if trees[i] > right {
			right = trees[i]
		}
	}

	for left <= right {
		mid = (left + right) / 2
		cutTree := 0

		for i := 0; i < n; i++ {
			if trees[i] > mid {
				cutTree += trees[i] - mid
			}
		}

		if cutTree >= m {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	fmt.Fprintln(writer, res)
}
```

## 시간 복잡도 분석
- `left` ~ `right` 범위에서 이진 탐색을 수행하므로 **O(log max_height)**
- 나무의 개수 `N`만큼 순회하며 잘라낸 나무 길이를 계산하므로 **O(N)**
- 최종적으로 **O(N log max_height)**의 시간 복잡도를 가집니다.

## 정리
- **이진 탐색**을 이용하여 최적의 절단기 높이를 찾았습니다.
- **시간 복잡도 O(N log max_height)** 로 매우 효율적입니다.
- **큰 입력값도 충분히 처리 가능**합니다.

