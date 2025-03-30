# 별 찍기 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/10997)

## 문제 설명
예제를 보고 규칙을 유추한 뒤, 별을 찍어 출력하는 문제입니다. 입력으로 주어진 정수 `N`에 따라 출력되는 별 모양은 정해진 규칙에 따라 반복됩니다.

### 입력
- 첫째 줄에 정수 `N`이 주어집니다. (1 ≤ N ≤ 100)

### 출력
- 첫째 줄부터 차례대로 별 모양을 출력합니다.

---

## 코드 풀이
다음은 이 문제를 해결하기 위한 Go 언어 코드입니다:

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

func Draw(n, x, y int, canvas [][]string) {
	if n == 1 {
		canvas[x][y] = "*"
		return
	}

	ySize := (n * 4) - 3
	xSize := ySize + 2

	maxX, maxY := (xSize + x), (ySize + y)

	for i := y; i < maxY; i++ {
		canvas[x][i] = "*"
		canvas[maxX-1][i] = "*"
	}
	for i := x; i < maxX; i++ {
		canvas[i][y] = "*"
		if i == x+1 {
			canvas[i+1][maxY-2] = "*"
		} else {
			canvas[i][maxY-1] = "*"
		}
	}

	if n == 2 {
		canvas[x+3][y+2] = "*"
		canvas[x+4][y+2] = "*"
	}

	Draw(n-1, x+2, y+2, canvas)
}

func main() {
	defer writer.Flush()

	var n, xSize, ySize int
	fmt.Fscan(reader, &n)
	ySize = (n * 4) - 3
	if n == 1 {
		xSize = 1
	} else {
		xSize = ySize + 2
	}
	canvas := make([][]string, xSize)
	for i := 0; i < xSize; i++ {
		canvas[i] = make([]string, ySize)
	}
	for i := 0; i < xSize; i++ {
		for j := 0; j < ySize; j++ {
			canvas[i][j] = " "
		}
	}
	Draw(n, 0, 0, canvas)
	for i := 0; i < xSize; i++ {
		for j := 0; j < ySize; j++ {
			fmt.Fprint(writer, canvas[i][j])
		}
		fmt.Fprintln(writer)
	}
}
```

---

## 문제 풀이 과정

### 1. 재귀적으로 별을 그리기
별을 그리는 핵심 함수는 `Draw` 함수입니다.
- 함수는 현재의 크기 `n`에 따라, 캔버스의 좌표 `(x, y)`부터 시작하여 별을 그립니다.
- 이후 크기를 줄여가며 재귀적으로 내부의 별을 그리게 됩니다.

### 2. 캔버스 초기화
캔버스를 생성한 뒤, 모든 칸을 공백 문자로 초기화합니다.

### 3. 외곽선을 그리기
각 단계에서 별의 외곽선을 그린 뒤, 내부에 별 모양을 추가합니다.

### 4. 재귀 호출
`n`이 1이 될 때까지 내부 사각형을 줄여가며 별을 그립니다.

---

## 입출력 예시

### 예제 입력 1
```
1
```

### 예제 출력 1
```
*
```

### 예제 입력 2
```
2
```

### 예제 출력 2
```
*****
*
* ***
* * *
* * *
*   *
*****
```

### 예제 입력 3
```
3
```

### 예제 출력 3
```
*********
*
* *******
* *     *
* * *** *
* * * * *
* * * * *
* *   * *
* ***** *
*       *
*********
```

### 예제 입력 4
```
4
```

### 예제 출력 4
```
*************
*
* ***********
* *         *
* * ******* *
* * *     * *
* * * *** * *
* * * * * * *
* * * * * * *
* * *   * * *
* * ***** * *
* *       * *
* ********* *
*           *
*************
```

---

## 핵심 아이디어
- 캔버스를 2차원 배열로 표현하여 별을 채웁니다.
- 문제의 규칙에 따라 별의 외곽선을 그리고, 재귀적으로 내부를 처리합니다.
- `n`이 커질수록 규칙적으로 별의 크기와 위치가 변화합니다.
