## 문제: 기타 레슨
👉🏻[문제 링크](https://www.acmicpc.net/problem/2343)

강토는 자신의 기타 강의 동영상을 N개의 강의로 나누어 녹화하였다. 이 강의들은 반드시 순서대로 녹화되어야 하며, M개의 블루레이에 나누어 담을 예정이다. 각 블루레이의 크기는 동일해야 하고, 그 크기를 최소로 하려 한다.

블루레이에 담기는 강의는 연속된 구간이어야 하며, 각 강의의 길이가 주어질 때 가능한 블루레이 크기 중 최소를 구하는 프로그램을 작성하라.

### 입력

* 첫째 줄에 강의 수 N (1 ≤ N ≤ 100,000)과 블루레이 수 M (1 ≤ M ≤ N)이 주어진다.
* 둘째 줄에 N개의 강의 길이가 순서대로 주어진다. (각 길이는 10,000을 넘지 않는 자연수)

### 출력

* 가능한 블루레이 크기 중 최소값을 출력한다.

### 예제 입력 1

```
9 3
1 2 3 4 5 6 7 8 9
```

### 예제 출력 1

```
17
```

---

## 해결 방법: 이분 탐색 (Binary Search)

이 문제는 **블루레이 크기를 이분 탐색**으로 찾아가는 전형적인 문제이다.

### 핵심 아이디어

* 블루레이의 최소 크기 = 가장 긴 강의 길이
* 블루레이의 최대 크기 = 모든 강의의 총합
* 이분 탐색으로 가능한 블루레이 크기(mid)를 기준으로, 해당 크기로 나눴을 때 M개 이하의 블루레이로 가능한지 확인한다.

### 구현 방식

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

func BinarySearch(n, m, total int, lesson []int) int {
	res := 0
	left, right := 1, total

	for left <= right {
		done := 0
		max := 0
		num := 0
		queue := lesson
		mid := (left + right) / 2

		for i := 0; i < n; i++ {
			if queue[0] > mid {
				break
			}

			if num+queue[0] <= mid {
				num += queue[0]
				queue = queue[1:]
				if max < num {
					max = num
				}
			} else {
				if max < num {
					max = num
				}
				num = 0
				num += queue[0]
				queue = queue[1:]
				done++
			}

			if i == n-1 {
				done++
			}
		}

		if len(queue) != 0 {
			left = mid + 1
		} else {
			if done > m {
				left = mid + 1
			} else {
				right = mid - 1
				res = max
			}
		}
	}
	return res
}

func main() {
	defer writer.Flush()

	var n, m, total int
	fmt.Fscan(reader, &n, &m)
	lesson := make([]int, n)

	for i := 0; i < n; i++ {
		min := 0
		fmt.Fscan(reader, &min)
		total += min
		lesson[i] = min
	}

	result := BinarySearch(n, m, total, lesson)
	fmt.Fprintln(writer, result)
}
```

### 시간복잡도

* 이분 탐색: O(log(sum of lesson lengths))
* 각 탐색마다 강의 리스트를 순회: O(N)
* 최종 시간복잡도: O(N log(sum))

### 핵심 포인트

* 이분 탐색으로 범위를 줄이며 최적값을 탐색
* mid값으로 가능한지 판단하는 서브루틴 구현 중요
* 가능한 최소 블루레이 크기를 구하는 문제이므로 가능한 경우는 결과로 저장하며 범위를 좁혀야 함

