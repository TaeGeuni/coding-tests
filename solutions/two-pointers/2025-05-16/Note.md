# 두 용액 문제 풀이

## 🧪 문제 설명

KOI 부설 과학연구소에서는 산성 용액과 알칼리성 용액을 섞어 특성값이 **0에 가장 가까운 용액**을 만들고자 한다.

- 산성 용액: 특성값이 **양수**
- 알칼리성 용액: 특성값이 **음수**
- 두 용액을 혼합하면 **특성값의 합**

### 🔢 입력
- 첫째 줄: 전체 용액 수 \( N \) (2 ≤ N ≤ 100,000)
- 둘째 줄: 오름차순으로 정렬된 용액의 특성값들

### ✅ 출력
- 특성값이 0에 가장 가까운 혼합 용액을 만드는 두 용액의 특성값을 오름차순으로 출력

---

## 💡 해결 아이디어

- **투 포인터 알고리즘 사용**
  - 리스트가 오름차순으로 주어지므로 양 끝에서 시작하는 두 포인터 `left`, `right`를 활용
  - `left + right` 값을 기준으로 0에 더 가까운 조합을 찾는다

### 알고리즘

1. `left = 0`, `right = N-1`에서 시작
2. 두 용액의 합 `sum = solution[left] + solution[right]`
3. `abs(sum)`이 최소라면 갱신
4. 합이 0보다 작으면 `left++`, 크면 `right--`
5. 합이 정확히 0이면 바로 종료 가능
6. 끝까지 탐색하여 가장 0에 가까운 조합을 선택

---

## ✅ 예제 입력 1

```
5
-99 -2 -1 4 98
```

## ✅ 예제 출력 1

```
-99 98
```

---

## 🧑‍💻 코드 구현 (Go)

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

func TwoPointers(n int, solution []int) (int, int) {
	res := 2_000_000_000
	var a, b int

	ls, rs := 0, n-1
	left, right := solution[ls], solution[rs]

	for i := 0; i < n-1; i++ {
		mid := left + right

		if mid < 0 {
			ls++
			mid = -mid
		} else if mid > 0 {
			rs--
		} else {
			a = left
			b = right
			break
		}
		if res >= mid {
			res = mid
			a = left
			b = right
		}

		left, right = solution[ls], solution[rs]
	}

	return a, b
}

func main() {
	defer writer.Flush()

	var n int

	fmt.Fscan(reader, &n)

	solution := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &solution[i])
	}

	a, b := TwoPointers(n, solution)

	fmt.Fprintln(writer, a, b)
}
```

---

## 📝 시간 복잡도

- O(N): 한 번의 순회로 해결 가능

## 🧠 핵심 포인트

- 오름차순 정렬된 배열에서 투 포인터는 매우 강력한 도구
- 절댓값 비교로 0에 가까운 조합을 찾는 것이 핵심
