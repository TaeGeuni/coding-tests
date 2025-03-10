# Padovan Sequence
👉🏻[문제 링크](https://www.acmicpc.net/problem/9461)

## 문제
그림과 같이 삼각형이 나선 모양으로 놓여져 있다. 첫 삼각형은 정삼각형으로 변의 길이는 1이다. 그 다음에는 다음과 같은 과정으로 정삼각형을 계속 추가한다. 나선에서 가장 긴 변의 길이를 k라 했을 때, 그 변에 길이가 k인 정삼각형을 추가한다.

![이미지 URL](https://www.acmicpc.net/upload/images/pandovan.png)

파도반 수열 P(N)은 나선에 있는 정삼각형의 변의 길이이다. P(1)부터 P(10)까지 첫 10개 숫자는 1, 1, 1, 2, 2, 3, 4, 5, 7, 9이다.

N이 주어졌을 때, P(N)을 구하는 프로그램을 작성하시오.

## 입력
첫째 줄에 테스트 케이스의 개수 T가 주어진다. 각 테스트 케이스는 한 줄로 이루어져 있고, N이 주어진다. (1 ≤ N ≤ 100)

## 출력
각 테스트 케이스마다 P(N)을 출력한다.

### 예제 입력 1
```
2
6
12
```

### 예제 출력 1
```
3
16
```

---

## 코드 구현 (Golang)
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func PadovanSequence(n int) int {
	res := 0

	if n == 1 || n == 2 || n == 3 {
		return 1
	}

	a, b, c := 1, 1, 1

	for i := 3; i < n; i++ {
		res = a + b
		a = b
		b = c
		c = res
	}

	return res
}

func main() {
	defer writer.Flush()

	var t, n int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		fmt.Fscan(reader, &n)
		fmt.Fprintln(writer, PadovanSequence(n))
	}
}
```

## 설명
- `PadovanSequence(n int) int`: 파도반 수열의 n번째 값을 반환하는 함수이다.
  - n이 1, 2, 3이면 1을 반환한다.
  - n이 4 이상이면 `P(n) = P(n-2) + P(n-3)`의 점화식을 이용하여 계산한다.
  - 동적 프로그래밍(DP, Dynamic Programming)을 활용하여 이전 결과를 저장하며 계산을 최적화한다.
- `main` 함수:
  - 입력을 받아 여러 개의 테스트 케이스를 처리한다.
  - `bufio`를 사용하여 빠른 입출력을 수행한다.

## 시간 복잡도
- `O(N)`, 즉 선형 시간 복잡도를 가진다. 최대 `N=100`이므로 충분히 빠르게 동작한다.

