# Fly me to the Alpha Centauri
👉🏻[문제 링크](https://www.acmicpc.net/problem/1011)

## 문제 설명

우현이는 새로운 세계로 이동하기 위해 공간이동 장치를 사용해야 합니다. 이 장치는 특정 규칙에 따라 작동하며, 최소 작동 횟수로 목표 지점에 도달해야 합니다.

### 입력
1. 첫 줄에 테스트 케이스 개수 `T`가 주어집니다.
2. 각 테스트 케이스마다 현재 위치 `x`와 목표 위치 `y`가 주어집니다.

### 출력
각 테스트 케이스에 대해 최소 작동 횟수를 출력합니다.

---

## 접근 방식

### 핵심 아이디어
1. 이동 거리의 패턴은 `1, 1 2 1, 1 2 2 1, 1 2 3 2 1, ...`처럼 양 끝이 1로 끝나고, 가운데 숫자가 증가하고 다시 감소하는 구조입니다.
2. 이동 횟수가 증가할수록, 같은 숫자의 이동 거리(`2`, `3`, ...)가 두 번씩 등장합니다.
3. 목표 거리 `D = y - x`에 대해 위의 패턴을 기반으로 최소한의 이동 횟수를 계산할 수 있습니다.

### 알고리즘
1. 초기 변수 설정:
   - `res`: 현재까지 누적된 이동 거리
   - `num`: 현재 이동 거리의 최대값
   - `count`: 장치 작동 횟수
2. 누적 이동 거리 `res`가 목표 거리 `D`보다 작을 때까지 반복:
   - `num` 값을 한 번 혹은 두 번씩 더해서 `res`를 증가
   - `count`를 증가시키며 횟수를 기록
3. 반복이 종료되면 `count`가 최소 작동 횟수가 됩니다.

---

## 구현 코드 (Go)

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

func MinimumOperations(x, y int) int {
	res := 0
	count := 0
	num := 1
	check := false

	for (y - x) > res {
		if check {
			res += num
			num++
			check = false
		} else {
			res += num
			check = true
		}
		count++
	}

	return count
}

func main() {
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)
	result := make([]int, 0)

	for i := 0; i < t; i++ {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		result = append(result, MinimumOperations(x, y))
	}
	for i := 0; i < t; i++ {
		fmt.Fprintln(writer, result[i])
	}
}
```

---

## 예제 설명

### 입력 예시
```
3
0 3
1 5
45 50
```

### 출력 예시
```
3
3
4
```

### 동작 설명
#### 테스트 케이스 1: `x = 0, y = 3`
1. `D = 3`
2. 이동 순서: `1 -> 1 -> 1`
3. 총 이동 횟수: **3**

#### 테스트 케이스 2: `x = 1, y = 5`
1. `D = 4`
2. 이동 순서: `1 -> 2 -> 1`
3. 총 이동 횟수: **3**

#### 테스트 케이스 3: `x = 45, y = 50`
1. `D = 5`
2. 이동 순서: `1 -> 2 -> 1 -> 1`
3. 총 이동 횟수: **4**

---

## 시간 및 공간 복잡도

- **시간 복잡도**: `O(D)` (`D = y - x`)만큼 반복하며 최소 이동 거리 누적
- **공간 복잡도**: `O(1)` — 단일 변수로 상태 추적

---

## 결론
이 문제는 이동 거리의 규칙적인 패턴을 이해하고, 그에 따라 누적 이동 거리와 작동 횟수를 관리하는 방식으로 해결할 수 있습니다. 단순한 구현으로도 정확하게 문제를 풀 수 있는 좋은 수학적 문제입니다.

