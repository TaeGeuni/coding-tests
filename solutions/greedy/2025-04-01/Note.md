# Log Jumping 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/11497)

## 문제 설명
N개의 통나무를 원형으로 배치하여 아이들이 뛰어다닐 수 있도록 한다. 이때, 인접한 두 통나무 간의 높이 차이를 최소화하는 배치를 찾는 것이 목표이다.

예를 들어, 통나무 높이가 `{2, 4, 5, 7, 9}`인 경우 `[2, 9, 7, 4, 5]`로 배치하면 최대 높이 차이가 `|2 - 9| = 7`이지만, `[2, 5, 9, 7, 4]`로 배치하면 최대 높이 차이가 `|5 - 9| = 4`로 줄어든다. 

## 입력
- 첫 줄에 테스트 케이스의 개수 T가 주어진다.
- 각 테스트 케이스의 첫 번째 줄에는 통나무의 개수 N(5 ≤ N ≤ 10,000)이 주어진다.
- 두 번째 줄에는 N개의 통나무 높이가 주어진다. (1 ≤ Li ≤ 100,000)

## 출력
각 테스트 케이스마다 최적의 통나무 배치에서 인접한 두 통나무 간의 최대 높이 차이를 출력한다.

## 예제 입력
```
3
7
13 10 12 11 10 11 12
5
2 4 5 7 9
8
6 6 6 6 6 6 6 6
```

## 예제 출력
```
1
4
0
```

## 풀이 방법
1. **정렬**: 통나무의 높이를 오름차순으로 정렬한다.
2. **번갈아 배치**: 정렬된 통나무를 번갈아 가며 배치하여 최대 높이 차이를 최소화한다.
   - 첫 번째 통나무에서 시작하여 한 칸씩 건너뛰며 배치
   - 반대 방향에서도 같은 방식으로 배치
3. **최대 높이 차이 계산**: 결과적으로 최적의 배치에서 인접한 통나무 간의 최대 높이 차이를 구한다.

## 시간 복잡도 분석
- **정렬 (O(N log N))**: 먼저 주어진 통나무 높이를 정렬한다.
- **배치 과정 (O(N))**: 번갈아 가며 통나무를 배치하는 과정은 한 번의 순회로 이루어진다.
- **최대 높이 차이 계산 (O(N))**: 인접한 통나무 간의 최대 높이 차이를 계산하는 과정도 한 번의 순회로 이루어진다.

따라서 전체 알고리즘의 시간 복잡도는 `O(N log N)`이다. 이는 정렬이 지배적인 연산이기 때문이다.

## 코드 구현
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

func LogJumping(logs []int) int {
	res := 0

	sort.Ints(logs)
	a := logs[0]

	if len(logs)%2 == 0 {
		i := 0
		for i < len(logs) {
			if res < logs[i]-a {
				res = logs[i] - a
			}
			a = logs[i]
			i += 2
		}
		i = len(logs) - 1
		for i > 0 {
			if res < a-logs[i] {
				res = a - logs[i]
			}
			a = logs[i]
			i -= 2
		}
	} else {
		i := 0
		for i < len(logs) {
			if res < logs[i]-a {
				res = logs[i] - a
			}
			a = logs[i]
			i += 2
		}
		i = len(logs) - 2
		for i > 0 {
			if res < a-logs[i] {
				res = a - logs[i]
			}
			a = logs[i]
			i -= 2
		}
	}

	return res
}

func main() {
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)
	result := make([]int, t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(reader, &n)
		logs := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &logs[j])
		}
		result[i] = LogJumping(logs)
	}
	for i := 0; i < t; i++ {
		fmt.Fprintln(writer, result[i])
	}
}
```

