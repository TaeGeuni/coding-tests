# 문제: 별 찍기
👉🏻[문제 링크](https://www.acmicpc.net/problem/10994)

예제를 보고 규칙을 유추한 뒤에 별을 찍는 문제입니다.

## 문제 설명

### 입력
첫째 줄에 정수 `N` (1 ≤ N ≤ 100)이 주어집니다.

### 출력
N에 따라 다음과 같은 규칙에 따라 별을 출력합니다:
- 첫 번째 패턴은 한 개의 별 `*`로 이루어져 있습니다.
- 두 번째 패턴은 테두리가 별로 감싸진 정사각형으로, 내부에 첫 번째 패턴이 포함됩니다.
- N번째 패턴은 테두리가 별로 감싸진 정사각형으로, 내부에 (N-1)번째 패턴이 포함됩니다.

### 예제
#### 입력 1
```
1
```
#### 출력 1
```
*
```

#### 입력 2
```
2
```
#### 출력 2
```
*****
*   *
* * *
*   *
*****
```

#### 입력 3
```
3
```
#### 출력 3
```
*********
*       *
* ***** *
* *   * *
* * * * *
* *   * *
* ***** *
*       *
*********
```

#### 입력 4
```
4
```
#### 출력 4
```
*************
*           *
* ********* *
* *       * *
* * ***** * *
* * *   * * *
* * * * * * *
* * *   * * *
* * ***** * *
* *       * *
* ********* *
*           *
*************
```

---

## 풀이
다음은 문제를 해결하기 위한 Go 언어 코드입니다:

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

func DrawStarDot(num, x, y int, stars [][]string) {
	if num == 1 {
		stars[x][y] = "*"
		return
	}

	size := 1 + (4 * (num - 1))
	endx, endy := (x + size), (y + size)

	// 테두리 그리기
	for i := y; i < endy; i++ {
		stars[x][i] = "*"
		stars[endx-1][i] = "*"
	}
	for i := x; i < endx; i++ {
		stars[i][y] = "*"
		stars[i][endy-1] = "*"
	}

	// 내부 재귀 호출
	DrawStarDot(num-1, x+2, y+2, stars)
}

func main() {
	defer writer.Flush()

	var num int
	fmt.Fscan(reader, &num)

	size := 1 + (4 * (num - 1))
	stars := make([][]string, size)
	for i := 0; i < size; i++ {
		stars[i] = make([]string, size)
		for j := 0; j < size; j++ {
			stars[i][j] = " "
		}
	}

	DrawStarDot(num, 0, 0, stars)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Fprint(writer, stars[i][j])
		}
		fmt.Fprintln(writer)
	}
}
```

---

## 코드 설명
1. **입력 처리:**
   - `N`을 입력받습니다.

2. **별 그리기 함수 (`DrawStarDot`):**
   - 주어진 패턴의 크기를 계산하고 테두리에 별을 그립니다.
   - 내부 정사각형을 재귀적으로 호출하여 별을 그립니다.

3. **출력:**
   - 2차원 배열 `stars`에 저장된 패턴을 출력합니다.

---

## 핵심 아이디어
- 패턴을 재귀적으로 그려 나갑니다.
- 현재 패턴의 테두리를 그리고, 내부 작은 패턴을 재귀 호출로 완성합니다.
- 별이 없는 공간은 공백으로 초기화합니다.

---

## 시간 복잡도
- 외부 루프는 `N`번 호출되며, 각 호출에서 크기 `O((4N)^2)`에 대해 작업합니다. 결과적으로, 시간 복잡도는 대략 `O(N^2)`입니다.
