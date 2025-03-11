# X보다 큰 수 중 가장 작은 수
👉🏻[문제 링크](https://www.acmicpc.net/problem/2992)

## 문제 설명
정수 X가 주어졌을 때, X와 구성이 같으면서 X보다 큰 수 중 가장 작은 수를 출력하는 프로그램을 작성한다.

수의 구성이 같다는 말은, 수를 이루고 있는 각 자리수가 같다는 뜻이다. 예를 들어, 123과 321은 수의 구성이 같지만, 123과 432는 구성이 같지 않다.

## 입력
첫째 줄에 X가 주어진다. (1 ≤ X ≤ 999999) X는 0으로 시작하지 않는다.

## 출력
첫째 줄에 결과를 출력한다. 만약 그러한 숫자가 없는 경우에는 0을 출력한다.

## 예제
### 예제 입력 1
```
156
```
### 예제 출력 1
```
165
```

### 예제 입력 2
```
330
```
### 예제 출력 2
```
0
```

### 예제 입력 3
```
27711
```
### 예제 출력 3
```
71127
```

## 풀이
이 문제는 주어진 수의 다음 순열(사전순으로 바로 다음에 오는 수)을 찾는 문제이다. 다음 순열을 찾는 방법은 다음과 같다:

1. 오른쪽에서부터 탐색하여, 처음으로 감소하는 위치 `m`을 찾는다.
2. `m`보다 오른쪽에서 `m`보다 큰 숫자 중 최솟값을 찾아 위치를 `n`으로 정한다.
3. `m`과 `n`을 교환한다.
4. `m` 이후의 숫자들을 오름차순으로 정렬한다.

이 과정을 거치면 X보다 크면서 가장 작은 수를 얻을 수 있다.

## 코드 구현 (Go)
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func BigAndSmallNumber(digits []byte) bool {
	l := len(digits)

	if l == 1 {
		return false
	}

	m := l - 2

	for (m >= 0) && (digits[m] >= digits[m+1]) {
		m--
	}
	if m < 0 {
		return false
	}

	n := l - 1

	for digits[m] >= digits[n] {
		n--
	}

	digits[m], digits[n] = digits[n], digits[m]

	sort.Slice(digits[m+1:], func(i, j int) bool {
		return digits[m+1+i] < digits[m+1+j]
	})

	return true
}

func main() {
	defer writer.Flush()
	var x string
	fmt.Fscan(reader, &x)

	digits := []byte(x)

	if BigAndSmallNumber(digits) {
		fmt.Fprintln(writer, string(digits))
	} else {
		fmt.Fprintln(writer, 0)
	}
}
```

## 설명
- **입력**: 문자열 형태로 숫자를 받아 바이트 배열로 변환한다.
- **BigAndSmallNumber 함수**:
  - 다음 순열을 찾는 과정 수행
  - 자리 바꿈 후 남은 숫자들을 정렬하여 가장 작은 수를 생성
  - 가능한 경우 true 반환, 불가능한 경우 false 반환
- **출력**: 가능한 경우 변경된 숫자를 출력하고, 불가능하면 `0`을 출력한다.

## 시간 복잡도
이 알고리즘의 시간 복잡도는 `O(n log n)`이다. 이유는:
- `m`을 찾는 과정: `O(n)`
- `n`을 찾는 과정: `O(n)`
- 정렬 과정: `O(n log n)`

따라서 전체적으로 `O(n log n)`의 시간 복잡도를 가진다.

