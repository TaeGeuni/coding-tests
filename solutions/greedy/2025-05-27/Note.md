# 행복 유치원 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/13164)

## 📝 문제 설명

태양이는 원생들을 키 순서대로 줄 세운 뒤, 이들을 **K개의 조**로 나누려 한다.  
각 조는 서로 인접한 원생들로만 구성되며, 조마다 단체 티셔츠를 맞추는 비용은 **가장 키가 큰 원생과 가장 키가 작은 원생의 차이**다.

**목표:** 티셔츠 제작 비용의 총합을 최소로 하여 조를 나누는 것.

---

## 📥 입력

- 첫 번째 줄: 원생 수 `N (1 ≤ N ≤ 300,000)`, 조의 수 `K (1 ≤ K ≤ N)`
- 두 번째 줄: 오름차순으로 정렬된 원생들의 키 `student[0..N-1]`

---

## 📤 출력

- 최소 티셔츠 제작 비용의 총합

---

## 📌 예제 입력

```
5 3
1 3 5 6 10
```

## 📌 예제 출력

```
3
```

---

## 💡 해결 아이디어

- 학생들은 키 순으로 정렬되어 있으므로, 인접한 학생들끼리의 키 차이를 미리 구한다.
- 이 차이들 중에서 **가장 큰 `K-1`개의 차이**를 기준으로 **조를 나누면 전체 비용이 최소**가 된다.
- 즉, 총 `N-1`개의 차이 중에서 **작은 값부터 더해서 `N-K`개만** 더하면 됨.

---

## ✅ 알고리즘 순서

1. 인접한 학생 간의 키 차이를 모두 계산해 리스트 `gap`에 저장.
2. `gap`을 오름차순 정렬.
3. 가장 작은 `N - K`개의 값을 더한 것이 최소 비용.

---

## 🧠 시간 복잡도

- `O(N log N)` — gap 계산 `O(N)`, 정렬 `O(N log N)`

---

## 💻 Go 코드

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func MinimumCost(n, k int, student []int) int {
	res := 0
	gap := make([]int, 0)

	for i := 0; i < n-1; i++ {
		gap = append(gap, student[i+1]-student[i])
	}
	sort.Ints(gap)

	for i := 0; i < n-k; i++ {
		res += gap[i]
	}

	return res
}

func main() {
	defer writer.Flush()
	var n, k int
	fmt.Fscan(reader, &n, &k)

	student := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &student[i])
	}

	fmt.Fprintln(writer, MinimumCost(n, k, student))
}
```

---

## 🧪 예제 설명

예제 입력: `5 3`  
학생 키: `1 3 5 6 10`  
인접 차이: `2 2 1 4` → 정렬: `1 2 2 4`  
`N-K = 2`, 가장 작은 두 개: `1 + 2 = 3`

따라서 최소 비용은 **3**
