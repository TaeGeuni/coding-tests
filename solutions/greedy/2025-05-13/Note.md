# 신입 사원 선발 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/1946)

## 문제 설명

진영 주식회사는 최고의 인재를 뽑기 위해 1차 서류심사, 2차 면접시험 두 가지 전형을 봅니다. 
선발 기준은 다음과 같습니다:

- 어떤 지원자 A가 다른 지원자 B보다 **서류 성적과 면접 성적이 모두 낮다면** A는 선발되지 않음
- 즉, 성적이 둘 다 더 나은 지원자만 있으면 탈락

이 조건을 만족하면서 **최대로 선발 가능한 지원자의 수**를 구해야 합니다.

---

## 입력 형식
- 첫째 줄에 테스트 케이스 개수 T (1 ≤ T ≤ 20)
- 각 테스트 케이스마다 다음 정보를 입력:
  - 첫째 줄: 지원자의 수 N (1 ≤ N ≤ 100,000)
  - 이후 N개의 줄: 각 지원자의 서류 성적, 면접 성적 순위 (1위부터 N위까지, 중복 없음)

## 출력 형식
- 각 테스트 케이스마다 선발할 수 있는 최대 인원 수를 한 줄씩 출력

---

## 예제 입력
```
2
5
3 2
1 4
4 1
2 3
5 5
7
3 6
7 3
4 2
1 4
5 7
2 5
6 1
```

## 예제 출력
```
4
3
```

---

## 알고리즘 설명 (Greedy)

### 아이디어
- 서류 성적 기준으로 오름차순 정렬
- 이후 면접 성적을 비교하며 선발 기준을 만족하는 사람만 카운트

### 로직
1. 서류 성적 기준으로 지원자를 정렬
2. 첫 번째 지원자를 무조건 선발 (서류 1등이므로)
3. 그 이후에는 지금까지 선발된 사람 중 면접 성적 가장 높은 사람보다 더 높은 순위(낮은 숫자)를 가진 경우만 선발

### 시간 복잡도
- 정렬: O(N log N)
- 선발 판단: O(N)
- 최대 N = 100,000 이므로 충분히 빠름

---

## Go 코드 구현
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

func Greedy(volunteer [][]int) int {
	pass := len(volunteer)

	sort.Slice(volunteer, func(i, j int) bool {
		return volunteer[i][0] < volunteer[j][0]
	})

	num := volunteer[0][1]

	for i := 1; i < len(volunteer); i++ {
		if num < volunteer[i][1] {
			pass--
		} else {
			num = volunteer[i][1]
		}
	}

	return pass
}

func main() {
	defer writer.Flush()

	var t, n int
	fmt.Fscan(reader, &t)

	res := make([]int, t)

	for i := 0; i < t; i++ {
		fmt.Fscan(reader, &n)
		volunteer := make([][]int, n)

		for j := 0; j < n; j++ {
			var document, interview int
			fmt.Fscan(reader, &document, &interview)
			volunteer[j] = append(volunteer[j], document)
			volunteer[j] = append(volunteer[j], interview)
		}

		res[i] = Greedy(volunteer)
	}

	for i := 0; i < t; i++ {
		fmt.Fprintln(writer, res[i])
	}
}
```

---

## 결론
- 이 문제는 **정렬 + 그리디**로 해결하는 대표적인 문제입니다.
- 핵심은 하나의 기준으로 정렬한 뒤, 다른 기준을 기준으로 최솟값을 계속 유지하면서 선발 가능한지 판단하는 것입니다.
- 시간 효율성을 만족하면서도 직관적인 전략입니다.
