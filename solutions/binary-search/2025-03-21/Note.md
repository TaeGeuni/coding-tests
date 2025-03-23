# 📦 보석 나눠주기 - 질투심 최소화 문제
👉🏻[문제 링크](https://www.acmicpc.net/problem/2792)

## 📝 문제 설명
보석 공장에서 M가지 색상의 보석을 N명의 학생들에게 나누어주려 한다. 학생들은 반드시 같은 색상의 보석만 받을 수 있고, 보석을 받지 못하는 학생이 있어도 된다.

하지만 어떤 학생이 너무 많은 보석을 받게 되면, 다른 학생들이 질투를 하게 된다. 이때 질투심은 "가장 많은 보석을 받은 학생의 보석 개수"로 정의된다.

원장 선생님은 이 질투심을 최소화하여 보석을 나눠주고자 한다.

## 📌 입력
- 첫째 줄: 학생 수 N (1 ≤ N ≤ 1,000,000,000), 색상 수 M (1 ≤ M ≤ 300,000, M ≤ N)
- 다음 M줄: 각 줄마다 K번 색상 보석의 개수 (1 이상 1,000,000,000 이하의 정수)

## 🎯 출력
- 질투심의 최솟값을 출력한다.

## 🔍 예제 입력 1
```
5 2
7
4
```

## ✅ 예제 출력 1
```
3
```

## 🔍 예제 입력 2
```
7 5
7
1
7
4
4
```

## ✅ 예제 출력 2
```
4
```

---

## 💡 해결 아이디어
이 문제는 **이분 탐색(Binary Search)** 을 이용해 해결할 수 있다.

1. **질투심의 최소값**은 1, **최대값**은 가장 많은 보석 개수인 `maxGem`이다.
2. 질투심을 `mid`라고 했을 때, 각 색상의 보석을 최대 `mid`개씩 나눠준다고 가정한다.
3. 이 경우, 해당 색상의 보석을 나눠주기 위해 필요한 학생 수는 `ceil(gem / mid)`이다.
4. 모든 색상에 대해 필요한 학생 수의 총합이 N을 초과하면 불가능한 분배이므로, `mid`를 증가시킨다.
5. 가능하다면 더 낮은 질투심으로도 가능한지 확인해본다.

---

## 🔎 코드 구현 (Golang)
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

func BinarySearch(n int, maxGem int, gems []int) int {
	left, right := 1, maxGem
	result := maxGem

	for left <= right {
		mid := (left + right) / 2
		count := 0

		for _, gem := range gems {
			// 각 색깔의 구슬을 mid개 이하로 나눠주기 위해 필요한 아이 수
			count += (gem + mid - 1) / mid // == ceil(gem / mid)
		}

		if count > n {
			left = mid + 1 // 아이가 부족하므로 envy level을 늘려야 함
		} else {
			result = mid   // 조건 만족, envy level을 줄여볼 수 있음
			right = mid - 1
		}
	}

	return result
}

func main() {
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	gems := make([]int, m)
	maxGem := 0

	for i := range gems {
		fmt.Fscan(reader, &gems[i])
		if gems[i] > maxGem {
			maxGem = gems[i]
		}
	}

	fmt.Fprintln(writer, BinarySearch(n, maxGem, gems))
}
```

---

## ⏱ 시간 복잡도 분석
- 이분 탐색 구간은 `1 ~ maxGem`이므로 최악의 경우 약 `log₂(maxGem)`번 반복된다.
- 각 반복마다 모든 색상에 대해 학생 수 계산을 해야 하므로 반복당 `O(M)`의 시간이 걸린다.

따라서 전체 시간 복잡도는:
```
O(M × log(maxGem))
```
- 여기서 `M`은 색상의 수 (최대 300,000), `maxGem`은 가장 많은 보석 개수 (최대 1,000,000,000)
- 최악의 경우에도 충분히 빠르게 동작한다.

---

## 🧠 핵심 포인트 정리
- 이분 탐색으로 질투심의 최소값을 찾는다.
- 각 색상의 보석을 최대 `mid`개로 나누기 위해 필요한 학생 수를 계산한다.
- 총 학생 수가 `N`을 넘지 않는 최소의 `mid` 값을 찾는다.

---
