# 떡 먹는 호랑이
👉🏻[문제 링크](https://www.acmicpc.net/problem/2502)

## 문제 설명
하루에 한 번 산을 넘어가는 떡 장사 할머니는 호랑이에게 떡을 주어야 산을 넘어갈 수 있다. 욕심 많은 호랑이는 어제 받은 떡의 개수와 그저께 받은 떡의 개수를 더한 만큼의 떡을 받아야만 할머니를 무사히 보내 준다.

예를 들어:
- 첫째 날: 1개
- 둘째 날: 2개
- 셋째 날: 3개 (1+2)
- 넷째 날: 5개 (2+3)

...

이와 같은 방식으로 D번째 날에 K개의 떡을 주었다면, 첫째 날과 둘째 날에 각각 몇 개의 떡을 주었는지를 구하는 문제이다.

## 입력 형식
- 첫 줄에 정수를 입력한다 D (3 <= D <= 30) K (10 <= K <= 100,000)

## 출력 형식
- 첫째 줄에 첫 날에 준 떡의 개수 A 출력
- 둘째 줄에 둘째 날에 준 떡의 개수 B 출력

## 예제 입력/출력
**입력:**
```
6 41
```
**출력:**
```
2
7
```

**입력:**
```
7 218
```
**출력:**
```
10
21
```

---

## Golang 코드 구현
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

func DP(d, k int) (int, int) {
	sli := make([]int, d)
	sli[d-1] = k

	for i := k - 1; i >= 1; i-- {
		sli[d-2] = i
		ok := true
		for j := 0; j < d-2; j++ {
			sli[d-3-j] = sli[d-1-j] - sli[d-2-j]
			if sli[d-3-j] <= 0 || sli[d-3-j] > sli[d-2-j] {
				ok = false
				break
			}
		}
		if ok {
			break
		}
	}

	f, t := sli[0], sli[1]
	return f, t
}

func main() {
	defer writer.Flush()

	var d, k int
	fmt.Fscan(reader, &d, &k)

	f, t := DP(d, k)

	fmt.Fprintln(writer, f)
	fmt.Fprintln(writer, t)
}
```

## 풀이 설명
- 슬라이스 `sli`는 D일간의 떡 개수를 저장한다.
- 마지막 날에 받은 떡(K)은 고정된 값이다. 그 전 날(i)의 값은 가능한 후보로 두고 반복문을 통해 역으로 추론한다.
- `sli[d-3-j] = sli[d-1-j] - sli[d-2-j]`를 통해 피보나치처럼 거꾸로 계산한다.
- 모든 값이 양수이면서 오름차순 조건을 만족하면 그 값을 채택한다.
- 조건을 만족하는 첫 번째 값을 찾으면 그 값이 정답이다.

## 시간 복잡도 분석
- 최악의 경우, `i`는 1부터 `K-1`까지 반복할 수 있으므로 외부 루프는 최대 `K`번 반복할 수 있습니다.
- 내부 루프는 길이 `D`에 비례하여 최대 `D-2`번 반복합니다.
- 따라서 전체 시간복잡도는 **O(K × D)**입니다.
- 하지만 문제 조건상 `K <= 100,000`이고, D도 최대 30이기 때문에, 실질적인 수행 시간은 매우 짧고 성능 이슈는 없습니다.

## 핵심 아이디어
- 이 문제는 역 피보나치 수열의 문제로 볼 수 있다.
- 거꾸로 추론하여 첫째 날과 둘째 날의 떡 개수를 찾아야 한다.
- 주어진 조건(1 <= A <= B)을 만족하는 정답은 항상 존재한다.

---

**끝.**

