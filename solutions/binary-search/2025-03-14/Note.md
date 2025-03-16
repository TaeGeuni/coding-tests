# 막대 과자 나누기 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/16401)

## 문제 설명
명절이 되면, 홍익이 집에는 조카들이 놀러 온다. 떼를 쓰는 조카들을 달래기 위해 홍익이는 막대 과자를 하나씩 나눠준다.

조카들이 과자를 먹는 동안은 떼를 쓰지 않기 때문에, 홍익이는 조카들에게 최대한 긴 과자를 나눠주려고 한다.

그런데 나눠준 과자의 길이가 하나라도 다르면 조카끼리 싸움이 일어난다. 따라서 반드시 모든 조카에게 같은 길이의 막대 과자를 나눠주어야 한다.

M명의 조카가 있고 N개의 과자가 있을 때, 조카 1명에게 줄 수 있는 막대 과자의 최대 길이를 구하라.

단, 막대 과자는 길이와 상관없이 여러 조각으로 나눠질 수 있지만, 과자를 하나로 합칠 수는 없다. 단, 막대 과자의 길이는 양의 정수여야 한다.

---

## 입력
- 첫째 줄에 조카의 수 M (1 ≤ M ≤ 1,000,000), 과자의 수 N (1 ≤ N ≤ 1,000,000)이 주어진다.
- 둘째 줄에 과자 N개의 길이 L1, L2, ..., LN이 공백으로 구분되어 주어진다.
- 과자의 길이는 (1 ≤ L1, L2, ..., LN ≤ 1,000,000,000) 를 만족한다.

## 출력
- 첫째 줄에 조카 1명에게 줄 수 있는 막대 과자의 최대 길이를 출력한다.
- 단, 모든 조카에게 같은 길이의 막대과자를 나눠줄 수 없다면, 0을 출력한다.

---

## 예제 입력 및 출력

### 예제 입력 1
```
3 10
1 2 3 4 5 6 7 8 9 10
```
### 예제 출력 1
```
8
```

### 예제 입력 2
```
4 3
10 10 15
```
### 예제 출력 2
```
7
```

---

## 해결 방법

이 문제는 **이분 탐색 (Binary Search)** 을 활용하여 최적의 막대 과자 길이를 찾는 방식으로 해결할 수 있다. 주어진 과자를 특정 길이로 잘랐을 때, M명의 조카에게 모두 나눠줄 수 있는지를 검사하는 방식으로 진행된다.

1. 가능한 막대 과자의 최소 길이는 1, 최대 길이는 주어진 과자 중 가장 긴 과자의 길이로 설정한다.
2. 이분 탐색을 통해 중간 길이(mid)를 설정하고, 이 길이로 과자를 나누었을 때 조카들에게 충분히 나눠줄 수 있는지를 계산한다.
3. 충분히 나눠줄 수 있다면, 더 긴 과자로 나눠줄 가능성이 있으므로 길이를 증가시킨다.
4. 충분히 나눠줄 수 없다면, 길이를 줄여야 한다.
5. 최적의 길이를 찾을 때까지 반복한다.

---

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

func BinarySearch(m, max int, snack []int) int {
	left, right := 1, max
	res := 0

	for left <= right {
		mid := (left + right) / 2
		v := 0

		for i := 0; i < len(snack); i++ {
			v += snack[i] / mid
		}
		if v >= m {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return res
}

func main() {
	defer writer.Flush()

	var m, n, max, total int
	fmt.Fscan(reader, &m, &n)

	snack := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &snack[i])
		total += snack[i]
		if snack[i] > max {
			max = snack[i]
		}
	}

	if total < m {
		fmt.Fprintln(writer, 0)
	} else {
		fmt.Fprintln(writer, BinarySearch(m, max, snack))
	}
}
```

---

## 시간 복잡도 분석

- **이분 탐색**: 최대 길이를 `L`이라 할 때, 탐색 범위는 `O(log L)`
- **각 탐색마다 N개의 과자를 확인**: `O(N)`
- **총 시간 복잡도**: `O(N log L)`

최대 `L = 1,000,000,000`, `N = 1,000,000` 이므로, `O(N log L)`의 복잡도는 충분히 효율적이다.

