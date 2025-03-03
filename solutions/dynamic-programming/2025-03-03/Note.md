# 01타일 문제 풀이 및 해설
👉🏻[문제 링크](https://www.acmicpc.net/problem/1904)

## 문제 설명
지원이는 1과 00으로만 이루어진 이진 수열을 만들고자 합니다. 길이가 `N`인 모든 가능한 이진 수열의 개수를 구하는 문제입니다. 단, 결과는 `15746`으로 나눈 나머지를 출력해야 합니다.

### 예제
#### 입력
```
4
```
#### 출력
```
5
```

## 해결 방법
이 문제는 피보나치 수열과 동일한 점화식을 가집니다.

- `dp[n] = dp[n-1] + dp[n-2]` (N개의 타일을 만들기 위해서는 `N-1`의 경우에 `1`을 추가하거나, `N-2`의 경우에 `00`을 추가하는 두 가지 방법이 존재)
- 단, `dp[n]`을 매 연산마다 `15746`으로 나누어 오버플로우를 방지합니다.

### 코드 1: 배열을 사용한 DP 방식
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	N := 0
	fmt.Fscan(reader, &N)
	tile := make([]int, N)

	for i := 0; i < N; i++ {
		if i == 0 {
			tile[i] = 1
		} else if i == 1 {
			tile[i] = 2
		} else {
			tile[i] = (tile[i-2] + tile[i-1]) % 15746
		}
	}

	fmt.Fprintln(writer, tile[N-1])
}
```

### 코드 2: 배열을 사용하지 않는 최적화 방식
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	N := 0
	fmt.Fscan(reader, &N)
	if N == 1 {
		fmt.Fprintln(writer, 1)
		return
	} else if N == 2 {
		fmt.Fprintln(writer, 2)
		return
	}

	a, b := 1, 2
	c := 0
	for i := 3; i < N+1; i++ {
		c = (a + b) % 15746
		a, b = b, c
	}

	fmt.Fprintln(writer, c)
}
```

## 코드 1 vs 코드 2 (차이점 분석)
| 비교 항목 | 코드 1 (배열 사용) | 코드 2 (배열 미사용) |
|-----------|-----------------|-----------------|
| 메모리 사용 | `O(N)` (`N` 크기의 배열 사용) | `O(1)` (세 개의 변수만 사용) |
| 실행 속도 | `O(N)` (배열 접근 포함) | `O(N)` (배열 접근 없음, 약간 더 빠름) |
| 공간 최적화 | 배열을 사용하여 `N`개의 값을 저장 | 이전 두 값만 저장하여 메모리 절약 |

### 결론
- **배열을 사용하는 방식(코드 1)**은 직관적이지만, `N`이 커질수록 메모리 사용량이 많아집니다.
- **배열을 사용하지 않는 방식(코드 2)**는 메모리를 적게 사용하고, 실행 속도도 약간 더 빠릅니다.
- 따라서 **두 번째 코드가 더 효율적인 해결 방법**입니다. 🚀

