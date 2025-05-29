# 빌딩 옥상 정원 확인 문제
👉🏻[문제 링크](https://www.acmicpc.net/problem/6198)

## 🧩 문제 설명

도시에는 N개의 빌딩이 있고, 각각의 빌딩은 오른쪽에 있는 빌딩들의 옥상 정원을 **자신보다 낮은 건물에 한해서** 볼 수 있습니다.

즉, i번째 빌딩의 높이가 `h[i]`일 때, i보다 오른쪽에 있는 j번째 빌딩의 높이 `h[j]`가 `h[i]`보다 작다면 볼 수 있습니다. 단, 중간에 자신보다 **높거나 같은** 빌딩이 하나라도 나타나면 그 이후의 모든 건물은 볼 수 없습니다.

> 예시  
> 입력:  
> `H = [10, 3, 7, 4, 12, 2]`  
> 출력: `5`

## ✅ 입력 형식

- 첫 번째 줄: 빌딩의 수 `N` (1 ≤ N ≤ 80,000)
- 두 번째 줄부터 `N+1`번째 줄까지: 각 빌딩의 높이 `h_i` (1 ≤ h_i ≤ 1,000,000,000)

## ✅ 출력 형식

- 각 관리인이 볼 수 있는 옥상 정원의 총 개수를 출력한다.

## 💡 풀이 아이디어

이 문제는 **스택을 활용한 역방향 단조 감소 스택** 문제로 접근할 수 있습니다.

- 오른쪽에서부터 한 빌딩씩 보며, 스택에는 지금까지 오른쪽에서 본 "낮은 건물"만 유지되게 합니다.
- 각 빌딩이 볼 수 있는 빌딩 수는 스택에 남아있는 수의 개수입니다.

### ✅ 주요 전략

- 각 빌딩을 순서대로 보면서, 자신보다 **작은 빌딩**은 볼 수 있으므로 스택에 쌓습니다.
- 자신보다 **큰 빌딩이 나타나면** 스택을 pop 하면서 제거합니다.
- 이렇게 하면 각 빌딩이 볼 수 있는 다음 빌딩 수를 `O(N)`에 계산할 수 있습니다.

## ✅ Go 언어 코드

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

func Benchmark(n int, building []int) int {
	res := 0
	stack := make([]int, 0)
	stack = append(stack, building[0])

	for i := 1; i < n; i++ {
		if stack[len(stack)-1] > building[i] {
			res += len(stack)
			stack = append(stack, building[i])
		} else {
			for len(stack) > 0 {
				stack = stack[:len(stack)-1]
				if len(stack) > 0 && stack[len(stack)-1] > building[i] {
					break
				}
			}
			res += len(stack)
			stack = append(stack, building[i])
		}
	}

	return res
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	building := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &building[i])
	}

	fmt.Fprintln(writer, Benchmark(n, building))
}
```

## 🧪 예제

입력:
```
6
10
3
7
4
12
2
```

출력:
```
5
```

## 📌 시간 복잡도

- 시간 복잡도: O(N)
- 공간 복잡도: O(N)

스택을 이용한 효율적인 풀이 덕분에 최대 80,000개의 빌딩까지도 빠르게 계산할 수 있습니다.
