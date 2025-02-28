# 랜선 자르기 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/1654)
## 📌 문제 설명
- 오영식은 K개의 랜선을 가지고 있으며, 이를 N개의 같은 길이 랜선으로 잘라야 합니다.
- 랜선을 자를 때 손실되는 길이는 없으며, 자른 랜선은 다시 붙일 수 없습니다.
- 만들 수 있는 **최대 랜선 길이**를 구하는 프로그램을 작성해야 합니다.

## 🔹 입력
1. 첫째 줄에 K(1 ≤ K ≤ 10,000), N(1 ≤ N ≤ 1,000,000)이 주어집니다. (K ≤ N)
2. 이후 K개의 줄에 각 랜선의 길이가 주어집니다. (최대 2^31 - 1)

## 🔹 출력
- **N개를 만들 수 있는 최대 랜선 길이**를 출력합니다.

## 🔹 예제 입력/출력
### ✅ 입력 예시
```
4 11
802
743
457
539
```
### ✅ 출력 예시
```
200
```

---

## 🚀 해결 방법 (이분 탐색 활용)
이 문제는 **이분 탐색(Binary Search)** 을 이용해 해결할 수 있습니다.

### 1️⃣ 탐색 범위 설정
- 만들 수 있는 랜선의 길이는 최소 `1cm`부터, 최대 `가장 긴 랜선 길이`까지 가능합니다.
- 따라서 **이분 탐색의 범위를 [1, max(랜선 길이)]로 설정**합니다.

### 2️⃣ 이분 탐색 수행
1. **mid 값을 기준으로 랜선을 자르고** 몇 개를 만들 수 있는지 확인합니다.
2. `N개 이상`을 만들 수 있다면, **더 긴 길이가 가능할 수도 있으므로** `left = mid + 1`로 이동합니다.
3. `N개보다 적게` 만들 수 있다면, **길이를 줄여야 하므로** `right = mid - 1`로 이동합니다.
4. 탐색이 끝날 때까지 위 과정을 반복합니다.

---

## 🔹 Go 언어 코드 구현
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var K, N int
	fmt.Fscanf(reader, "%d %d\n", &K, &N)

	lanCables := make([]int, K)
	var maxLen int

	for i := 0; i < K; i++ {
		fmt.Fscanf(reader, "%d\n", &lanCables[i])
		if lanCables[i] > maxLen {
			maxLen = lanCables[i]
		}
	}

	left, right := 1, maxLen
	var result int

	for left <= right {
		mid := (left + right) / 2
		if countCables(lanCables, mid) >= N {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	fmt.Fprintln(writer, result)
}

func countCables(cables []int, length int) int {
	total := 0
	for _, cable := range cables {
		total += cable / length
	}
	return total
}
```

---

## 🧐 왜 `mid + 1`, `mid - 1`을 사용하는가?
- **탐색 범위를 줄이지 않으면 무한 루프 발생 가능**
- `left = mid` 또는 `right = mid`를 사용하면 `mid`가 반복되어 무한 루프가 생길 수 있음
- 따라서, **탐색 범위를 확실히 줄이기 위해 `mid + 1`, `mid - 1`을 사용**

### ✅ 정리
- **`left = mid + 1`**: `N개 이상`을 만들 수 있으면 더 긴 길이 가능성이 있음 → 길이 증가
- **`right = mid - 1`**: `N개 미만`이면 더 짧은 길이에서만 가능 → 길이 감소

---

## 🔹 시간 복잡도 분석
- 이분 탐색을 사용하므로 **O(K log M)** (M: 가장 긴 랜선 길이)

---

## 🎯 최종 정리
✅ **이분 탐색**을 활용해 `left`, `right` 범위를 조정하며 최적의 길이를 찾는다.
✅ `N개 이상 만들 수 있는 최대 랜선 길이`를 출력하면 정답!
