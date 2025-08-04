# CCW (Counter ClockWise) 판별 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/11758)

## 문제 설명

2차원 좌표 평면 위에 있는 점 3개 `P1`, `P2`, `P3`가 주어진다.  
`P1`, `P2`, `P3`를 순서대로 이은 선분이 어떤 방향을 이루고 있는지 구하는 프로그램을 작성하시오.

- 반시계 방향이면 `1`
- 시계 방향이면 `-1`
- 일직선이면 `0`

을 출력한다.

---

## 입력

- 첫째 줄: P1의 좌표 `(x1, y1)`
- 둘째 줄: P2의 좌표 `(x2, y2)`
- 셋째 줄: P3의 좌표 `(x3, y3)`
- (-10,000 ≤ xi, yi ≤ 10,000)
- 모든 좌표는 정수이고 서로 다르다.

## 출력

- 방향에 따라 `1`, `-1`, `0` 중 하나를 출력한다.

---

## 예제 입력 1

```
1 1
5 5
7 3
```

### 예제 출력 1

```
-1
```

---

## 해결 아이디어

세 점이 이루는 방향은 **외적 (cross product)** 을 이용하여 판별할 수 있다.

점 `p1`, `p2`, `p3`가 있을 때 다음과 같은 벡터를 생각한다.

- 벡터 `A = p2 - p1`
- 벡터 `B = p3 - p1`

이 두 벡터의 외적의 부호를 보면 방향을 알 수 있다.

공식:

```
(p2.x - p1.x) * (p3.y - p1.y) - (p3.x - p1.x) * (p2.y - p1.y)
```

- 양수 → 반시계 방향 (1)
- 음수 → 시계 방향 (-1)
- 0 → 일직선 (0)

---

## Go 코드

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

type Point struct {
	x int
	y int
}

func ccw(p1, p2, p3 Point) int {
	res := (p2.x-p1.x)*(p3.y-p1.y) - (p3.x-p1.x)*(p2.y-p1.y)

	if res > 0 {
		return 1
	} else if res < 0 {
		return -1
	}
	return 0
}

func main() {
	defer writer.Flush()

	var p1, p2, p3 Point

	fmt.Fscan(reader, &p1.x, &p1.y)
	fmt.Fscan(reader, &p2.x, &p2.y)
	fmt.Fscan(reader, &p3.x, &p3.y)

	fmt.Fprintln(writer, ccw(p1, p2, p3))
}
```

---

## 시간복잡도

- O(1): 상수 시간 안에 외적 연산만 수행한다.

## 관련 알고리즘

- 기하학(Geometry)
- 벡터 외적(Cross Product)
