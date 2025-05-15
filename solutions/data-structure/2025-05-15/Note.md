# 트럭 - 다리 건너기 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/13335)

## 문제 설명

하나의 차선으로 된 다리를 n개의 트럭이 순서대로 건너가야 한다. 

다리는 다음과 같은 조건을 가진다:

- 다리의 길이: `w` (단위길이)
- 최대 하중: `L`
- 트럭은 1초에 1단위 길이씩 이동
- 다리에는 동시에 최대 `w`개의 트럭이 존재할 수 있음
- 동시에 올라간 트럭들의 무게 합은 `L`을 초과하면 안 됨

모든 트럭이 다리를 건너는 데 걸리는 최소 시간을 구하라.

## 입력

- 첫 줄: 트럭의 수 `n` (1 ≤ n ≤ 1,000), 다리의 길이 `w` (1 ≤ w ≤ 100), 최대 하중 `L` (10 ≤ L ≤ 1,000)
- 둘째 줄: 트럭의 무게 `a1, a2, ..., an` (1 ≤ ai ≤ 10)

## 출력

- 모든 트럭이 다리를 건너는 데 걸리는 최소 시간 출력

---

## 예제 입력 1
```
4 2 10
7 4 5 6
```

## 예제 출력 1
```
8
```

---

## 풀이 방법 (Go 코드)

BFS 혹은 시뮬레이션 방식으로, 트럭이 다리 위에 올라가는 시점부터 다리를 완전히 빠져나갈 때까지를 시뮬레이션 한다.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

type MoveAndTruck struct {
	truckNumber int
	move        int
}

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func CrossBridge(n, w, l int, truck []int) int {
	res := 0
	nowWeight := 0
	next := 0
	end := false

	bridge := make([]MoveAndTruck, 0)
	bridge = append(bridge, MoveAndTruck{truckNumber: next, move: 0})
	nowWeight += truck[next]
	if next == n-1 {
		end = true
	} else {
		next++
	}

	for len(bridge) > 0 {
		for i := 0; i < len(bridge); i++ {
			bridge[i].move++
		}
		if bridge[0].move > w {
			nowWeight -= truck[bridge[0].truckNumber]
			bridge = bridge[1:]
		}

		if !end {
			if nowWeight+truck[next] <= l {
				bridge = append(bridge, MoveAndTruck{truckNumber: next, move: 1})
				nowWeight += truck[next]
				if next == n-1 {
					end = true
				} else {
					next++
				}
			}
		}
		res++
	}

	return res
}

func main() {
	defer writer.Flush()

	var n, w, l int
	fmt.Fscan(reader, &n, &w, &l)

	truck := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &truck[i])
	}

	result := CrossBridge(n, w, l, truck)
	fmt.Fprintln(writer, result)
}
```

---

## 시간 복잡도 분석
- 최악의 경우 트럭 수만큼 반복하면서 각각의 트럭을 다리에 올리고, 최대 `w`까지 유지되므로 O(nw) 수준이지만 n과 w가 충분히 작아 문제 없음

---

## 핵심 아이디어
- 다리를 큐처럼 사용하여 현재 트럭의 위치와 무게 추적
- 트럭이 다리를 벗어나면 무게 차감
- 다음 트럭을 조건에 따라 다리에 추가

---

## 예제 실행 결과
```
입력: 4 2 10
트럭 무게: [7, 4, 5, 6]
출력: 8
```
